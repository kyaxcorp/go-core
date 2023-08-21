package autoloader

import (
	"github.com/kyaxcorp/go-core/core/helpers/slice"
	// cassandraConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/cassandra/config"

	cockroachConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/cockroach/config"
	mysqlConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/mysql/config"
	websocketClientConfig "github.com/kyaxcorp/go-core/core/clients/websocket/config"
	websocketClientConnection "github.com/kyaxcorp/go-core/core/clients/websocket/connection"
	cfgData "github.com/kyaxcorp/go-core/core/config/data"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	httpConfig "github.com/kyaxcorp/go-core/core/listeners/http/config"
	websocketServerConfig "github.com/kyaxcorp/go-core/core/listeners/websocket/config"
	loggingConfig "github.com/kyaxcorp/go-core/core/logger/config"
)

func createMaps() {
	if cfgData.MainConfig.Clients.MySQL.Instances == nil {
		// Creating  the map, where we will set afterwards the default object
		cfgData.MainConfig.Clients.MySQL.Instances = make(map[string]mysqlConfig.Config)
	}
	if cfgData.MainConfig.Clients.Cockroach.Instances == nil {
		// Creating  the map, where we will set afterwards the default object
		cfgData.MainConfig.Clients.Cockroach.Instances = make(map[string]cockroachConfig.Config)
	}
	if cfgData.MainConfig.Listeners.Http.Instances == nil {
		// Creating  the map, where we will set afterwards the default object
		cfgData.MainConfig.Listeners.Http.Instances = make(map[string]httpConfig.Config)
	}
	if cfgData.MainConfig.Listeners.WebSocket.Instances == nil {
		// Creating  the map, where we will set afterwards the default object
		cfgData.MainConfig.Listeners.WebSocket.Instances = make(map[string]websocketServerConfig.Config)
	}
	if cfgData.MainConfig.Clients.WebSocket.Instances == nil {
		// Creating  the map, where we will set afterwards the default object
		cfgData.MainConfig.Clients.WebSocket.Instances = make(map[string]websocketClientConfig.Config)
	}
	if cfgData.MainConfig.Logging.Channels == nil {
		// Creating  the map, where we will set afterwards the default object
		cfgData.MainConfig.Logging.Channels = make(map[string]loggingConfig.Config)
	}
}

func setDefaults(cfg Config) error {
	// Create the map!
	//MainConfig.ClientsStatus.MySQL.Connections = make(map[string]mysql.Config)
	//MainConfig.Listeners.Http.Instances = make(map[string]http.Config)

	// This is the default config!

	// Set default settings

	var _err error

	//---------------------------------------------------------------------------------\\
	//---------------------------\\    MYSQL CLIENT    //------------------------------\\
	//---------------------------------------------------------------------------------\\

	// Mysql Default Config Instance
	if _, ok := cfgData.MainConfig.Clients.MySQL.Instances["default"]; !ok {
		dbInstance, _err := mysqlConfig.SetDefaults(nil)
		if _err != nil {
			panic(_err)
		}

		// If the map is empty... we need to create it
		//if cfgData.MainConfig.Clients.MySQL.Instances == nil {
		//	// Creating  the map, where we will set afterwards the default object
		//	cfgData.MainConfig.Clients.MySQL.Instances = make(map[string]mysqlConfig.Config)
		//}
		cfgData.MainConfig.Clients.MySQL.Instances["default"] = *dbInstance
	}

	// Looping through instances and setting defaults if they are missing
	for connectionName, dbInstance := range cfgData.MainConfig.Clients.MySQL.Instances {
		_, _err = mysqlConfig.SetDefaults(&dbInstance)
		if _err != nil {
			panic(_err)
		}
		cfgData.MainConfig.Clients.MySQL.Instances[connectionName] = dbInstance
	}

	//---------------------------------------------------------------------------------\\
	//---------------------------\\    MYSQL CLIENT    //------------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//---------------------------------------------------------------------------------\\
	//-----------------------\\    COCKROACHDB CLIENT    //----------------------------\\
	//---------------------------------------------------------------------------------\\

	// Cockroach Default Config Instance
	if _, ok := cfgData.MainConfig.Clients.Cockroach.Instances["default"]; !ok {
		dbInstance, _err := cockroachConfig.SetDefaults(nil)
		if _err != nil {
			panic(_err)
		}

		// If the map is empty... we need to create it
		//if cfgData.MainConfig.Clients.Cockroach.Instances == nil {
		//	// Creating  the map, where we will set afterwards the default object
		//	cfgData.MainConfig.Clients.Cockroach.Instances = make(map[string]cockroachConfig.Config)
		//}
		cfgData.MainConfig.Clients.Cockroach.Instances["default"] = *dbInstance
	}

	// Looping through instances and setting defaults if they are missing
	for connectionName, dbInstance := range cfgData.MainConfig.Clients.Cockroach.Instances {
		_, _err = cockroachConfig.SetDefaults(&dbInstance)
		if _err != nil {
			panic(_err)
		}
		cfgData.MainConfig.Clients.Cockroach.Instances[connectionName] = dbInstance
	}
	//---------------------------------------------------------------------------------\\
	//-----------------------\\    COCKROACHDB CLIENT    //----------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//

	//---------------------------------------------------------------------------------\\
	//-----------------------\\    CASSANDRA CLIENT    //------------------------------\\
	//---------------------------------------------------------------------------------\\

	// Cassandra Default Config Instance
	/*if _, ok := cfgData.MainConfig.ClientsStatus.Cassandra.Instances["default"]; !ok {
		_cassandra := &cassandraConfig.Config{
			Hosts: []cassandraConfig.Host{
				cassandraConfig.Host{
					Destination: "",
					Port:        0,
				},
			},
		}
		if _err := _struct.SetDefaultValues(_cassandra); _err != nil {
			panic(_err)
		}
		// If the map is empty... we need to create it
		if cfgData.MainConfig.ClientsStatus.Cassandra.Instances == nil {
			// Creating  the map, where we will set afterwards the default object
			cfgData.MainConfig.ClientsStatus.Cassandra.Instances = make(map[string]cassandraConfig.Config)
		}
		cfgData.MainConfig.ClientsStatus.Cassandra.Instances["default"] = *_cassandra
	}*/

	//---------------------------------------------------------------------------------\\
	//-----------------------\\    CASSANDRA CLIENT    //------------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//---------------------------------------------------------------------------------\\
	//--------------------------\\    HTTP SERVER    //--------------------------------\\
	//---------------------------------------------------------------------------------\\

	// Http Default Config Instance
	if _, ok := cfgData.MainConfig.Listeners.Http.Instances["default"]; !ok {
		_http := &httpConfig.Config{}
		if _err := _struct.SetDefaultValues(_http); _err != nil {
			panic(_err)
		}
		// If the map is empty... we need to create it
		//if cfgData.MainConfig.Listeners.Http.Instances == nil {
		//	// Creating  the map, where we will set afterwards the default object
		//	cfgData.MainConfig.Listeners.Http.Instances = make(map[string]httpConfig.Config)
		//}
		cfgData.MainConfig.Listeners.Http.Instances["default"] = *_http
	}

	// Loop through all connections and set the default values
	for instanceName, httpInstance := range cfgData.MainConfig.Listeners.Http.Instances {

		// Logger config
		// Setting default values for logger
		if _err := _struct.SetDefaultValues(&httpInstance.Logger); _err != nil {
			panic(_err)
		}

		if _err := _struct.SetDefaultValues(&httpInstance); _err != nil {
			panic(_err)
		}

		if len(httpInstance.ListeningAddresses) == 0 {
			// No listening addresses, let's add one
			httpInstance.ListeningAddresses = []string{
				// the + symbol is the rule that checks if the port is busy, and if it is, it will
				// do +1 until finds a free port!
				"0.0.0.0:8080+",
			}
		}

		if len(httpInstance.ListeningAddressesSSL) == 0 {
			// No listening addresses, let's add one
			httpInstance.ListeningAddressesSSL = []string{
				// the + symbol is the rule that checks if the port is busy, and if it is, it will
				// do +1 until finds a free port!
				"0.0.0.0:8443+",
			}
		}

		// Set the logger to websocket config
		cfgData.MainConfig.Listeners.Http.Instances[instanceName] = httpInstance
	}

	//---------------------------------------------------------------------------------\\
	//--------------------------\\    HTTP SERVER    //--------------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//---------------------------------------------------------------------------------\\
	//-----------------------\\    WEBSOCKET SERVER    //------------------------------\\
	//---------------------------------------------------------------------------------\\

	// Http Default Config Instance
	if _, ok := cfgData.MainConfig.Listeners.WebSocket.Instances["default"]; !ok {
		_websocketServer := &websocketServerConfig.Config{}
		if _err := _struct.SetDefaultValues(_websocketServer); _err != nil {
			panic(_err)
		}
		// If the map is empty... we need to create it
		//if cfgData.MainConfig.Listeners.WebSocket.Instances == nil {
		//	// Creating  the map, where we will set afterwards the default object
		//	cfgData.MainConfig.Listeners.WebSocket.Instances = make(map[string]websocketServerConfig.Config)
		//}
		cfgData.MainConfig.Listeners.WebSocket.Instances["default"] = *_websocketServer
	}

	// Loop through all connections and set the default values
	for instanceName, wsInstance := range cfgData.MainConfig.Listeners.WebSocket.Instances {

		// Logger config
		// Setting default values for logger
		if _err := _struct.SetDefaultValues(&wsInstance.Logger); _err != nil {
			panic(_err)
		}

		if _err := _struct.SetDefaultValues(&wsInstance); _err != nil {
			panic(_err)
		}

		if len(wsInstance.ListeningAddresses) == 0 {
			// No listening addresses, let's add one
			wsInstance.ListeningAddresses = []string{
				// the + symbol is the rule that checks if the port is busy, and if it is, it will
				// do +1 until finds a free port!
				"0.0.0.0:8080+",
			}
		}

		if len(wsInstance.ListeningAddressesSSL) == 0 {
			// No listening addresses, let's add one
			wsInstance.ListeningAddressesSSL = []string{
				// the + symbol is the rule that checks if the port is busy, and if it is, it will
				// do +1 until finds a free port!
				"0.0.0.0:8443+",
			}
		}

		// Set the logger to websocket config
		cfgData.MainConfig.Listeners.WebSocket.Instances[instanceName] = wsInstance
	}

	//---------------------------------------------------------------------------------\\
	//-----------------------\\    WEBSOCKET SERVER    //------------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//---------------------------------------------------------------------------------\\
	//---------------------\\    WEBSOCKET CLIENT    //--------------------------------\\
	//---------------------------------------------------------------------------------\\

	// WebSocket Client Default Config
	if _, ok := cfgData.MainConfig.Clients.WebSocket.Instances["default"]; !ok {
		// Create the default config for websocket
		_webSocketClient := &websocketClientConfig.Config{}
		if _err := _struct.SetDefaultValues(_webSocketClient); _err != nil {
			panic(_err)
		}

		// If the map is empty... we need to create it
		//if cfgData.MainConfig.Clients.WebSocket.Instances == nil {
		//	// Creating  the map, where we will set afterwards the default object
		//	cfgData.MainConfig.Clients.WebSocket.Instances = make(map[string]websocketClientConfig.Config)
		//}

		// Set finally the default config
		cfgData.MainConfig.Clients.WebSocket.Instances["default"] = *_webSocketClient
	}

	// Loop through all connections and set the default values
	for instanceName, wsInstance := range cfgData.MainConfig.Clients.WebSocket.Instances {

		// If the default connection doesn't exist, create it!
		if exists, _ := slice.IndexExists(wsInstance.Connections, 0); !exists {
			// Create a default connection config for websocket
			wsInstance.Connections = append(wsInstance.Connections, &websocketClientConnection.Connection{})
		}

		// Loop through other connections and check if hey are ok!
		for connIndex, conn := range wsInstance.Connections {
			// Set the standard values for the object
			if _err := _struct.SetDefaultValues(conn); _err != nil {
				panic(_err)
			}
			// Set the authentication options
			if _err := _struct.SetDefaultValues(&conn.AuthOptions); _err != nil {
				panic(_err)
			}

			wsInstance.Connections[connIndex] = conn
		}

		// Logger config
		// Setting default values for logger
		if _err := _struct.SetDefaultValues(&wsInstance.Logger); _err != nil {
			panic(_err)
		}

		// Reconnect Options
		if _err := _struct.SetDefaultValues(&wsInstance.Reconnect); _err != nil {
			panic(_err)
		}

		if _err := _struct.SetDefaultValues(&wsInstance); _err != nil {
			panic(_err)
		}

		// Set the logger to websocket config
		cfgData.MainConfig.Clients.WebSocket.Instances[instanceName] = wsInstance
	}
	//---------------------------------------------------------------------------------\\
	//---------------------\\    WEBSOCKET CLIENT    //--------------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//---------------------------------------------------------------------------------\\
	//---------------------\\    LOGGING CHANNELS    //--------------------------------\\
	//---------------------------------------------------------------------------------\\

	// Logging Default config Channel
	if _, ok := cfgData.MainConfig.Logging.Channels["default"]; !ok {
		_logging := &loggingConfig.Config{}
		if _err := _struct.SetDefaultValues(_logging); _err != nil {
			panic(_err)
		}
		// If the map is empty... we need to create it
		//if cfgData.MainConfig.Logging.Channels == nil {
		//	// Creating  the map, where we will set afterwards the default object
		//	cfgData.MainConfig.Logging.Channels = make(map[string]loggingConfig.Config)
		//}
		cfgData.MainConfig.Logging.Channels["default"] = *_logging
	}

	// Check if there are additional channels defined from the main app
	if cfg.AdditionalLoggingChannels != nil {
		for channelName, channel := range cfg.AdditionalLoggingChannels {
			// Set the channel to existing map

			// Check if the channel already exists in the config
			if _, ok := cfgData.MainConfig.Logging.Channels[channelName]; !ok {
				// if it doesn't exist, get the default object
				if _err := _struct.SetDefaultValues(&channel); _err != nil {
					panic(_err)
				}
				// Set back
				cfgData.MainConfig.Logging.Channels[channelName] = channel
			} else {
				// if exists do nothing...
			}

		}
	}

	// Let's set the default values  for all channels
	for channelName, channel := range cfgData.MainConfig.Logging.Channels {
		if _err := _struct.SetDefaultValues(&channel); _err != nil {
			panic(_err)
		}
		// Set back
		cfgData.MainConfig.Logging.Channels[channelName] = channel
	}
	//---------------------------------------------------------------------------------\\
	//---------------------\\    LOGGING CHANNELS    //--------------------------------\\
	//---------------------------------------------------------------------------------\\

	return nil
}
