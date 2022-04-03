package client

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/clients/broker/config"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket"
	websocketConfig "github.com/KyaXTeam/go-core/v2/core/clients/websocket/config"
	wsConnection "github.com/KyaXTeam/go-core/v2/core/clients/websocket/connection"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/conv"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/logger"
	loggerConfig "github.com/KyaXTeam/go-core/v2/core/logger/config"
	loggerPaths "github.com/KyaXTeam/go-core/v2/core/logger/paths"
)

func New(
	ctx context.Context,
	config config.Config,
) (*Client, error) {

	if !conv.ParseBool(config.IsEnabled) {
		return nil, define.Err(0, "broker client is disabled", config.PipeName)
	}

	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}

	// Generate the logger config
	// If Name has not being set for the logger in the config,
	// It will automatically take the name from the instance!
	if config.Logger.Name == "" {
		config.Logger.Name = config.PipeName
	}
	// If DirLogPath is not defined, it will set the default folder!
	if config.Logger.DirLogPath == "" {
		config.Logger.DirLogPath = loggerPaths.GetLogsPathForClients("broker/" + config.Logger.Name)
	}

	// Set Module Name
	if config.Logger.ModuleName == "" {
		config.Logger.ModuleName = "Broker Client=" + config.PipeName
	}

	// Set the default values for the config... that's in case something is missed
	loggerDefaultConfig, _err := loggerConfig.DefaultConfig(&config.Logger)

	if _err != nil {
		return nil, _err
	}

	brokerConnections := config.Connections
	var wsConnections []*wsConnection.Connection
	for _, brokerConn := range brokerConnections {
		wsConn, _err := wsConnection.DefaultConfig(nil)
		if _err != nil {
			// TODO:
		}
		wsConn.IsSecure = brokerConn.IsSecure
		wsConn.Host = brokerConn.Host
		wsConn.Port = brokerConn.Port
		wsConn.UriPath = "/api_ws/broker/pipes/" + config.PipeName
		wsConn.AcceptCertificate = brokerConn.AcceptCertificate
		wsConn.MaxRetries = brokerConn.MaxRetries
		wsConn.RetryTimeout = brokerConn.RetryTimeout
		// Enable always authentication
		wsConn.EnableAuth = "yes"
		wsConn.AuthOptions = wsConnection.AuthOptions{
			// Set authentication method to 1
			AuthType: wsConnection.AuthByToken,
			Token:    config.AuthToken,
		}
		// Add the connection
		wsConnections = append(wsConnections, &wsConn)
	}

	// Generating the final config
	wsConfig := websocketConfig.Config{
		Name: config.PipeName,
		// ---------RELATED TO CONNECTION----------\\
		AutoReconnect: config.AutoReconnect,
		Reconnect: websocketConfig.ReconnectOptions{
			TimeoutSeconds: config.Reconnect.TimeoutSeconds,
			MaxRetries:     config.Reconnect.MaxRetries,
		},
		UseMultipleConnections: config.UseMultipleConnections,
		// Set the connections
		Connections: wsConnections,
		// ---------RELATED TO CONNECTION----------\\
	}

	// Create the logger instance for broker
	brokerLoggerInstance := logger.New(loggerDefaultConfig)
	// Set same logger config for websocket, but make some minor changes
	wsConfig.Logger = loggerDefaultConfig
	// Setting the writer to websocket instance config, so it will write the logs to the brokers logs
	wsConfig.Logger.ParentWriter = brokerLoggerInstance.MainWriter
	// Set writing to parent as mandatory!
	wsConfig.Logger.WriteToParent = "yes"
	// Disable file writing, we don't need websocket to save separately the logs
	wsConfig.Logger.FileIsEnabled = "no"

	// We create the websocket client here... if additional functionality will be added to broker, we can move it to
	// Start/Connect function
	wsClient, _err := websocket.New(ctx, wsConfig)
	if _err != nil {

		return nil, define.Err(0, "failed to create websocket client for broker", _err.Error())
	}

	c := &Client{
		// Broker config
		config: config,
		// Websocket config
		wsConfig: wsConfig,
		// The client will be created on connect!
		WSClient: wsClient,
		// Set the parent Context
		parentCtx: ctx,
		// Create the bool
		ctxDone: _bool.New(),
		// Logger
		Logger: brokerLoggerInstance,
	}
	return c, nil
}
