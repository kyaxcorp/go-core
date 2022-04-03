package client

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_int"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_map_string_interface"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_time"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_uint16"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_uint64"
	"github.com/KyaXTeam/go-core/v2/core/logger"
	loggerConfig "github.com/KyaXTeam/go-core/v2/core/logger/config"
	loggerPaths "github.com/KyaXTeam/go-core/v2/core/logger/paths"
	"net/http"
)

func New(
	ctx context.Context,
	config config.Config,
) (*Client, error) {

	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}

	// If Name has not being set for the logger in the config,
	// It will automatically take the name from the instance!
	if config.Logger.Name == "" {
		config.Logger.Name = config.Name
	}
	// If DirLogPath is not defined, it will set the default folder!
	if config.Logger.DirLogPath == "" {
		config.Logger.DirLogPath = loggerPaths.GetLogsPathForClients("websocket/" + config.Logger.Name)
	}

	// Set Module Name
	if config.Logger.ModuleName == "" {
		config.Logger.ModuleName = "WebSocket Client=" + config.Name
	}

	// Set the default values for the config... that's in case something is missed
	loggerDefaultConfig, _err := loggerConfig.DefaultConfig(&config.Logger)

	if _err != nil {
		return nil, _err
	}

	client := &Client{
		config: config,

		RequestHeader: make(http.Header),

		connectStartTime:    _time.New(),
		connectEndTime:      _time.New(),
		disconnectStartTime: _time.New(),
		disconnectEndTime:   _time.New(),

		// Sent
		nrOfSentMessages:        _uint64.New(),
		nrOfSentTextMessages:    _uint64.New(),
		nrOfSentBinaryMessages:  _uint64.New(),
		nrOfSentFailedMessages:  _uint64.New(),
		nrOfSentSuccessMessages: _uint64.New(),
		sentBytes:               _uint64.New(),
		sentBinaryBytes:         _uint64.New(),
		sentTextBytes:           _uint64.New(),

		// Received
		nrOfReceivedMessages:       _uint64.New(),
		nrOfReceivedTextMessages:   _uint64.New(),
		nrOfReceivedBinaryMessages: _uint64.New(),
		receivedBytes:              _uint64.New(),
		receivedBinaryBytes:        _uint64.New(),
		receivedTextBytes:          _uint64.New(),

		// Other Statistics
		nrOfDisconnections: _uint64.New(),

		connectedTo:  _int.NewVal(-1), // -1 is nothing defined...
		connectingTo: _int.NewVal(-1), // -1 is nothing defined...
		isConnected:  _bool.New(),
		isConnecting: _bool.New(),

		isDisconnecting: _bool.New(),
		isReconnecting:  _bool.New(),
		reconnectRound:  _uint16.New(),

		// Create the logger with default settings
		Logger: logger.New(loggerDefaultConfig),

		writeChannel: make(chan []byte),

		// Events
		onReceive:               _map_string_interface.New(),
		onText:                  _map_string_interface.New(),
		onBinary:                _map_string_interface.New(),
		onSend:                  _map_string_interface.New(),
		onSendError:             _map_string_interface.New(),
		onReadError:             _map_string_interface.New(),
		onError:                 _map_string_interface.New(),
		onBeforeDisconnect:      _map_string_interface.New(),
		onLinkDisconnect:        _map_string_interface.New(),
		onDisconnect:            _map_string_interface.New(),
		onConnect:               _map_string_interface.New(),
		onReconnecting:          _map_string_interface.New(),
		onReconnected:           _map_string_interface.New(),
		onReconnectFailed:       _map_string_interface.New(),
		onBeforeConnectToServer: _map_string_interface.New(),
		onConnectError:          _map_string_interface.New(),
		onConnectFailed:         _map_string_interface.New(),
		onConnectFailedAll:      _map_string_interface.New(),
		onConnectSuccess:        _map_string_interface.New(),
		onTerminate:             _map_string_interface.New(),
		onStopConnectingFinish:  _map_string_interface.New(),

		parentCtx: ctx,
		ctxDone:   _bool.New(),
		//connections: make(map[string]*connection.Connection),
	}
	return client, nil
}
