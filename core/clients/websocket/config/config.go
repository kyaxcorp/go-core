package config

import (
	"github.com/kyaxcorp/go-core/core/clients/websocket/connection"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
)

type ReconnectOptions struct {
	// TimeoutSeconds -> 5 Seconds, after disconnect happened, the client will reconnect after
	// indicated time!
	TimeoutSeconds uint16 `yaml:"timeout_seconds" mapstructure:"timeout_seconds" default:"5"`
	// MaxRetries -> -1 -> infinite! Maximum nr of retries....
	MaxRetries int16 `yaml:"max_retries" mapstructure:"max_retries" default:"-1"`
}

type Config struct {
	// Name -> is the name of the instance... usually it's being used for logger or other things as naming...
	Name string
	// AutoReconnect -> Enable auto reconnection in case of disconnect!
	AutoReconnect string `yaml:"auto_reconnect" mapstructure:"auto_reconnect" default:"yes"`
	Reconnect     ReconnectOptions
	// UseMultipleConnections -> Enable utilization of existing listed connections, if the client cannot connect to previous, will take the next one!
	UseMultipleConnections string `yaml:"use_multiple_connections" mapstructure:"use_multiple_connections" default:"yes"`
	//Connections            map[string]*connection.Connection
	Connections []*connection.Connection
	// This is the logger configuration!
	Logger loggerConfig.Config
}

// TODO: we should add connections, Reconnect options
// DefaultConfig -> it will return the default config with default values
func DefaultConfig(configObj *Config) (Config, error) {
	if configObj == nil {
		configObj = &Config{}
	}
	var _err error
	// Set the default values for the object!
	_err = _struct.SetDefaultValues(configObj)
	if _err != nil {
		return *configObj, _err
	}
	// Setting logger defaults
	_err = _struct.SetDefaultValues(&configObj.Logger)
	if _err != nil {
		return *configObj, _err
	}

	return *configObj, _err
}
