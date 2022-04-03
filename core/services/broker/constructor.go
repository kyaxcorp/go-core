package broker

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_bool"
	websocketConfig "github.com/kyaxcorp/go-core/core/listeners/websocket/config"
	websocketServer "github.com/kyaxcorp/go-core/core/listeners/websocket/server"
	"github.com/kyaxcorp/go-core/core/logger"
	"github.com/kyaxcorp/go-core/core/logger/appLog"
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
	loggerPaths "github.com/kyaxcorp/go-core/core/logger/paths"
	brokerConfig "github.com/kyaxcorp/go-core/core/services/broker/config"
	"github.com/rs/zerolog"
	"sync"
)

// TODO: use middlewares!
// authentication
// TODO: use also Salsa20 encryption?! even if it's decrypted the SSL!
// Even if there'll be a man in middle attack!

/*
	Now we should create different things through Broker...
	FOr events, we need pipes...
	Meaning that events will subscribe and unsubscribe to specific events...
	The broker should transit requests

	If creating pipes mechanism, each pipe will have an ID/Name
	Each process can connect to a pipe by using an ID/Name
	Everyone can be a subscriber or a publisher, the broker will broadcast to everyone connected to it!
	The pipe will be created automatically when connecting through URI:    /api_ws/xxxx/YOUR_PIPE_NAME
	If the pipe is not utilized for more than 10 minutes, it will be closed...

	TODO: maybe later we will need to create a buffer of messages for one's that have lost connection!?


*/

func New(
	ctx context.Context,
	instanceName string,
	config brokerConfig.Config,
) (*Broker, error) {
	info := func() *zerolog.Event {
		return appLog.InfoF("New Broker")
	}
	_error := func() *zerolog.Event {
		return appLog.ErrorF("New Broker")
	}
	_debug := func() *zerolog.Event {
		return appLog.DebugF("New Broker")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")

	// Creating the cancel context
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}

	info().Msg("configuring logger...")

	var loggerDirPath string
	// Setting default values for logger
	if config.Logger.Name == "" {
		config.Logger.Name = instanceName
	}
	// If DirLogPath is not defined, it will set the default folder!
	if config.Logger.DirLogPath == "" {
		loggerDirPath = loggerPaths.GetLogsPathForServers("broker/" + config.Logger.Name)
		config.Logger.DirLogPath = loggerDirPath + filesystem.DirSeparator() + "server" + filesystem.DirSeparator()
		_debug().Str("generated_dir_log_path", config.Logger.DirLogPath).Msg("logger DirLogPath empty, generating...")
	} else {
		// Let's set the one defined by the config/user, but add another sub folder
		loggerDirPath = config.Logger.DirLogPath
		// Correct the path!
		config.Logger.DirLogPath = filesystem.FilterPath(loggerDirPath + filesystem.DirSeparator() + "server" + filesystem.DirSeparator())
		_debug().Str("generated_dir_log_path", config.Logger.DirLogPath).Msg("correcting dir log path")
	}

	// Set Module Name
	if config.Logger.ModuleName == "" {
		config.Logger.ModuleName = "Broker Server=" + instanceName
	}

	info().Msg("creating broker instance...")

	// Set the default values for the config... that's in case something is missed
	loggerDefaultConfig, _err := loggerConfig.DefaultConfig(&config.Logger)
	if _err != nil {
		_error().Msg(color.Style{color.LightRed}.Render("failed to set default config for logger"))
		return nil, define.Err(0, "failed to set default config for logger", _err.Error())
	}

	// Creating broker logger instance
	brokerLoggerInstance := logger.New(loggerDefaultConfig)

	info().Msg("creating websocket config...")
	// Create the websocket config
	wsConfig, _err := websocketConfig.DefaultConfig(nil)
	if _err != nil {
		_error().Msg("failed to get websocket default config")
		return nil, define.Err(0, "failed to get websocket default config", _err.Error())
	}

	// TODO: check if it's right to set directly same configuration from broker to websocket!
	// TODO: i think something should be modified or adapted
	wsConfig.Logger = loggerDefaultConfig

	wsConfig.Logger.ModuleName = "Broker WebSocket Server=" + instanceName

	// Setting the writer to websocket instance config, so it will write the logs to the brokers logs
	wsConfig.Logger.ParentWriter = brokerLoggerInstance.MainWriter
	// Set writing to parent as mandatory!
	wsConfig.Logger.WriteToParent = "yes"
	// Disable file writing, we don't need websocket to save separately the logs
	wsConfig.Logger.FileIsEnabled = "no"
	// Set the name of websocket from the brokers instance name
	wsConfig.Name = instanceName

	if conv.ParseBool(config.IsListenPlain) {
		// Listen on HTTP without SSL
		wsConfig.EnableUnsecure = "yes"
		wsConfig.ListeningAddresses = config.ListeningAddresses
	} else {
		wsConfig.EnableUnsecure = "no"
	}
	// Check if SSL Listening is allowed
	if conv.ParseBool(config.IsListenSSL) {
		// Listen on SSL
		wsConfig.EnableSSL = "yes"
		wsConfig.ListeningAddressesSSL = config.ListeningAddressesSSL
	}

	info().Interface("broker_config", config).Msg("websocket generated config")

	info().Msg("creating websocket instance...")
	// We create here the server with nil context, but later, on start, we will add it again!
	s, _err := websocketServer.New(nil, wsConfig)
	if _err != nil {
		_error().Err(_err).Msg(color.Style{color.LightRed}.Render("failed to create websocket instance"))
		return nil, _err
	}
	info().Msg(color.Style{color.LightGreen}.Render("websocket instance created..."))

	// Create the broker object
	broker := &Broker{
		parentCtx:             ctx,
		Server:                s,
		pipesLock:             sync.RWMutex{},
		shutdownHubMonitoring: make(chan bool),
		config:                config,

		isStarted:  _bool.New(),
		isStarting: _bool.New(),
		isStopping: _bool.New(),

		Logger: brokerLoggerInstance,
	}
	info().Msg(color.Style{color.LightGreen}.Render("broker instance created..."))

	info().Msg("initializing broker config...")
	// Initialize broker configurations
	broker.init()

	return broker, nil
}

func (b *Broker) init() {
	info := func() *zerolog.Event {
		return b.LInfoF("New Broker - init")
	}
	_debug := func() *zerolog.Event {
		return b.LInfoF("New Broker - init")
	}
	_error := func() *zerolog.Event {
		return b.LInfoF("New Broker - init")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")

	info().Msg("defining routes & middlewares...")
	// Server status is already enabled!
	b.Server.WSServer.GET("/", func(c *gin.Context) {
		_debug().Msg("are you lost?")
		c.JSON(200, gin.H{
			"message": "Are you Lost?)",
		})
	})

	// Launch hubs monitoring and handler!
	// Setup Group Routing
	apiWSGroup := b.Server.WSServer.Group("/api_ws")
	brokerGroup := apiWSGroup.Group("/broker")
	pipesGroup := brokerGroup.Group("/pipes")
	// Enable authentication
	pipesGroup.Use(b.CheckIsAuthenticated()) // Set Authentication on Pipes
	// Setup main pipe route and handler
	pipesGroup.GET("/:pipeName", func(c *gin.Context) {

		// On a new request
		// Get pipe name
		pipeName := c.Param("pipeName")
		_debug().Str("pipe_name", pipeName).Msg("accessed pipe")

		// Check if there is a pipe name
		if pipeName == "" {
			_error().Msg("pipe name empty...")
			c.JSON(400, gin.H{
				"message": "Pipe name empty...",
			})
			return
		}

		// Use locks when reading from stack
		// Create a new hub for this pipe name if it doesn't exist!
		// If exists, then simply upgrade client to websocket
		b.pipesLock.RLock()
		if _, ok := b.Pipes[pipeName]; !ok {
			info().Str("pipe_name", pipeName).Msg("creating new pipe/hub")
			// If not exists
			b.pipesLock.RUnlock()
			b.pipesLock.Lock()
			// Create a new hub for this pipe
			b.Pipes[pipeName] = b.Server.NewHub(nil)
			// Start the hub
			b.Pipes[pipeName].Start()
			b.pipesLock.Unlock()
		} else {
			// Do nothing if exists...
			b.pipesLock.RUnlock()
		}

		// Upgrade to websocket
		b.Server.UpgradeToWS(
			c,
			func(message *websocketServer.ReceivedMessage, c *websocketServer.Client, s *websocketServer.Server) {
				// If we receive a message from someone... we simply broadcast it!
				// But don't broadcast it to the sender which generated the message!! We should add an exception list

				b.Pipes[pipeName].BroadcastByReceivedMessageTypeTo(
					message,
					websocketServer.FindClientsFilter{
						// To everyone
						All: true,
						// Send except this connection!
						ExceptConnections: []uint64{c.GetConnectionID()},
					},
				)

				// TODO: we should also broadcast to the other nodes from the cluster?!...
				// The idea is to send also to the other nodes from the cluster, and they therefore check
				// if there are any clients on their side... and if there are, they will also broadcast this message
				// the mechanism is simple, but sometimes it can be overwhelming...
				// so, we can make some optimizations for the pipes, we should make them know who's where...
				// And for that, they should sync about their clients
				// Each node on each connect/disconnect should broadcast information about:
				// - NODE ID, PIPE NAME, other information is simply unnecessary because thy simply need to know
				// if there are any connections on a node that need a specific info...
				// Each node from the cluster is a client and a server
				// Each of them should create a client connection to the other nodes!
				// In this case each node will have 2 connections with the other node! meaning that first node has connected
				// to the second node, and vice versa!
				// Broadcasting should be async if possible!...because there are different clients/server with different
				// latencies... and for some of them can take some time to write a message, for some not...
				// Each node connecting as a Client, will become as a listener for the server... meaning, it will
				// wait for messages from the server
				// On a node receives a message from a client, it will check in the stack what are nodes have
				// registered pipes with this name, and broadcast to them this message
				// the nodes that will receive broadcast messages, will simply send these messages further to their clients
				// It's kind of a P2P Connection Mechanism, or even 1 TO 1 or MANY TO MANY Connection Mechanism!
				// Each node can configure to which other nodes it connects to
				// Also, for better optimization, each node can have connections with not all of the nodes, but with some
				// of them, and therefore the other nodes will forward the messages to the other nodes
				// It's important that the message will not go in a loop through the cluster
				// For that we need to create a payload_id for the message
				// And also somehow attach to the payload_id to whom the message has being sent already, possibly the node id
				// Each node from the cluster will become some kind of a router...
				// These pipes are some kind of interconnected branches, which forward messages to one another
				// When the nodes are connecting to each other, they should also take care of sending metadata about
				// their connections
				// A problem it might be with connections... if some nodes disappear when messages are being sent..
				// The messages that are being send, they include marks of the nodes to which are being sent the message...
				// SO in case of disappearance, we will need to see what nodes have being disconnected, and possibly repeat it...
				// Or even better, we can simply attach to the Payload ID, and each of the nodes will save in memory if this payload id
				// was being before, and if it has being sent, in this case, each of them will transmit and will transmit to each other
				// the message...

				// Another interesting mechanism is to know where each of the nodes is localized, and what's the latency
				// between each of them...
				// each of them will have to measure the latency between them...
				// a pipe name, it's kind of a address, or a CONNECTION between 2 or more clients are communicating through
				// Through a pipe we can send anything or everything! They kind of replacing the need of IP Addresses and Ports
				// They are kind of a multi related neurons which are passing through the signals

				// For best performance and functionality of this kind of network, we need a good protocol of intercommunication
				// and ROUTE FORMATION!
				// For a good route formation, we need to know Latency and ThroughPut to that destination!
				// Also, for good consistency, we need confirmations from the other side (destination) that the message
				// has being received...
				// This kind of network it's kind of untraceable..., well but it can contain all the necessary data about the client

				// But still, the scope of this mechanism, is not to create a network, but to broadcast data to listeners,
				// that's it! A network can be simply created with VPN's
			}, func(c *websocketServer.Client, s *websocketServer.Server) {
				b.Pipes[pipeName].RegisterClient(c)
			})
	})

}
