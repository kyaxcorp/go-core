package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/helpers/certs"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_bool"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_map_string_interface"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_time"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_uint64"
	"github.com/kyaxcorp/go-core/core/listeners/http/config"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/connection"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/request_timing"
	"github.com/kyaxcorp/go-core/core/listeners/http/routes/ping"
	"github.com/kyaxcorp/go-core/core/logger"
	"github.com/kyaxcorp/go-core/core/logger/appLog"
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
	loggerPaths "github.com/kyaxcorp/go-core/core/logger/paths"
	"github.com/rs/zerolog"
)

// You can use the default constructor!
func New(
	ctx context.Context,
	config config.Config,
) (*Server, error) {
	info := func() *zerolog.Event {
		return appLog.InfoF("New HTTP Server")
	}
	warn := func() *zerolog.Event {
		return appLog.WarnF("New HTTP Server")
	}
	_debug := func() *zerolog.Event {
		return appLog.DebugF("New HTTP Server")
	}
	_error := func() *zerolog.Event {
		return appLog.ErrorF("New HTTP Server")
	}

	info().Msg("entering...")
	defer info().Msg("leaving...")

	if !conv.ParseBool(config.IsEnabled) {
		_error().Str("instance_name", config.Name).Msg("http server is disabled, check your config")
		return nil, define.Err(0, "http server is disabled, check your config", config.Name)
	}

	info().Msg("configuring logger")

	var loggerDirPath string
	// Setting default values for logger
	if config.Logger.Name == "" {
		config.Logger.Name = config.Name
	}
	// If DirLogPath is not defined, it will set the default folder!
	if config.Logger.DirLogPath == "" {
		loggerDirPath = loggerPaths.GetLogsPathForServers("http/" + config.Logger.Name)
		config.Logger.DirLogPath = loggerDirPath + filesystem.DirSeparator() + "server" + filesystem.DirSeparator()
		_debug().Str("generated_dir_log_path", config.Logger.DirLogPath).Msg("logger DirLogPath empty, generating...")
	} else {
		loggerDirPath = config.Logger.DirLogPath
		// Correct the path!
		config.Logger.DirLogPath = file.FilterPath(loggerDirPath + filesystem.DirSeparator() + "server" + filesystem.DirSeparator())
		_debug().Str("generated_dir_log_path", config.Logger.DirLogPath).Msg("correcting dir log path")
	}

	// Set Module Name
	if config.Logger.ModuleName == "" {
		config.Logger.ModuleName = "HTTP Server=" + config.Name
	}

	if conv.ParseBool(config.EnableUnsecure) && len(config.ListeningAddresses) == 0 {
		_error().Msg(color.Style{color.LightRed}.Render("unsecure connections are enabled but no listening addresses"))
		return nil, define.Err(0, "no listening addresses are provided", config.Name)
	}

	if conv.ParseBool(config.EnableSSL) && len(config.ListeningAddressesSSL) == 0 {
		_error().Msg(color.Style{color.LightRed}.Render("secure connections are enabled but no listening addresses"))
		return nil, define.Err(0, "no listening ssl addresses are provided", config.Name)
	}

	if conv.ParseBool(config.EnableSSL) {
		info().Msg("checking certificates...")
		// if ssl enabled, check certificates
		sslKeyFilePathEmpty := false
		sslCertFilePathEmpty := false
		if config.SSLKeyFilePath == "" {
			sslKeyFilePathEmpty = true
			warn().Msg("param SSLKeyFilePath empty...")
		}
		if config.SSLCertFilePath == "" {
			sslCertFilePathEmpty = true
			warn().Msg("param SSLCertFilePath empty...")
		}

		// Missing paths...
		if sslKeyFilePathEmpty && sslCertFilePathEmpty {
			warn().Msg("params SSLCertFilePath & SSLKeyFilePath are empty, checking auto generation...")

			// Auto Generating certificates
			if conv.ParseBool(config.SSLAutoGenerateCerts) {
				info().Msg("auto generating ssl certificates")
				// Auto Generate
				certsConfig := &certs.CertGeneration{
					Host: "localhost",
				}

				certificatesInstanceName := "http_" + config.Name
				info().Str("certs_instance_name", certificatesInstanceName).Msg("generating instance name")
				// TODO: should we filter the naming... it's important filtration for files!
				_err := certs.GenerateCerts(certificatesInstanceName, certsConfig)
				if _err != nil {
					return nil, define.Err(0, "failed to generate http certificates", _err.Error(), config.Name)
				}
				info().Msg(color.Style{color.LightGreen}.Render("certificates generated successfully"))

				config.SSLKeyFilePath = certsConfig.KeyPath
				config.SSLCertFilePath = certsConfig.CertPath
			}
		} else {
			// Error?!
			return nil, define.Err(0, "http ssl key file or certificate is empty", config.Name)
		}

		_debug().Str("ssl_cert_file_path", config.SSLCertFilePath).Msg("ssl certificate")
		_debug().Str("ssl_key_file_path", config.SSLKeyFilePath).Msg("ssl key")

		if config.SSLCertFilePath == "" && config.SSLKeyFilePath == "" {
			// If still empty... then let's throw an error!
			_err := define.Err(0, "both certificate and key are empty, server will not start", config.Name)
			_error().Err(_err).Msg("")
			return nil, _err
		}
	}

	// Set the default values for the config... that's in case something is missed
	loggerDefaultConfig, _err := loggerConfig.DefaultConfig(&config.Logger)
	if _err != nil {
		_error().Msg(color.Style{color.LightRed}.Render("failed to set default config for logger"))
		return nil, define.Err(0, "failed to set default config for logger", _err.Error())
	}
	info().Msg("creating server instance")

	s := &Server{
		Name:        config.Name,
		Description: config.Description,

		connectionID: _uint64.New(),
		startTime:    _time.New(),
		stopTime:     _time.New(),

		isStopped:     _bool.New(),
		isStopCalled:  _bool.New(),
		isStartCalled: _bool.New(),
		isStarted:     _bool.New(),

		LoggerDirPath: loggerDirPath,
		Logger:        logger.New(loggerDefaultConfig),

		//
		enableSSL:   conv.ParseBool(config.EnableSSL),
		sslCertPath: config.SSLCertFilePath,
		sslKeyPath:  config.SSLKeyFilePath,
		//
		enableUnsecure: conv.ParseBool(config.EnableUnsecure),
		//
		ListeningAddresses:    config.ListeningAddresses,
		ListeningAddressesSSL: config.ListeningAddressesSSL,

		// Events/Callbacks - they are common for all routes!
		onRequest:  _map_string_interface.New(),
		onResponse: _map_string_interface.New(),

		// Stop
		//onStop       map[string]OnStop
		onStop:       _map_string_interface.New(),
		onBeforeStop: _map_string_interface.New(),
		onStopped:    _map_string_interface.New(),

		// Start
		onStart:       _map_string_interface.New(),
		onBeforeStart: _map_string_interface.New(),
		onStarted:     _map_string_interface.New(),

		HttpServer: nil,

		enableServerStatus: _bool.New(),
	}

	infoServer := func() *zerolog.Event {
		return s.LInfoF("New HTTP Server")
	}

	infoServer().Msg("setting context")
	s.SetContext(ctx)

	// By Default we set server as Stopped!
	s.isStopped.Set(true)

	infoServer().Msg("setting http server to release mode")
	// Set in Release mode, in this way the debugging messages are off!
	// We should not control here the Mode of the GIN,
	// We should not even use gin debugging, use a middleware for better logging!
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)

	infoServer().Msg("creating new gin server")

	// TODO: use HTTP/2?

	// Create the HTTP SERVER
	s.HttpServer = gin.New()

	infoServer().Msg("enabling auto recovery")
	// Set Auto Recovery!
	s.HttpServer.Use(gin.Recovery())

	infoServer().Msg("setting default middleware for connections")
	// We set as default middle related to connections
	// TODO: maybe we should also log into a folder the arriving connections!
	s.HttpServer.Use(connection.GetMiddleware(s.Logger))
	// Latency in processing
	s.HttpServer.Use(request_timing.GetMiddleware(s.Logger))

	// Set ping listener
	ping.Ping(s.HttpServer)

	if conv.ParseBool(config.EnableServerStatus) {
		infoServer().Msg("enabling server status")
		s.SetStatusCredentials(config.ServerStatusUsername, config.ServerStatusPassword)
		s.EnableServerStatus()
	}

	infoServer().Msg("leaving http constructor")
	return s, nil
}
