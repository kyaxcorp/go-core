package config

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/broker/connection"
	loggerConfig "github.com/KyaXTeam/go-core/v2/core/logger/config"
)

type ReconnectOptions struct {
	// TimeoutSeconds -> 5 Seconds, after disconnect happened, the client will reconnect after
	// indicated time!
	TimeoutSeconds uint16 `yaml:"timeout_seconds" mapstructure:"timeout_seconds" default:"5"`
	// MaxRetries -> -1 -> infinite! Maximum nr of retries....
	MaxRetries int16 `yaml:"max_retries" mapstructure:"max_retries" default:"-1"`
}

type Config struct {
	// Enable the broker...
	IsEnabled string `yaml:"is_enabled" mapstructure:"is_enabled" default:"no"`

	//----------RELATED TO BROKER---------\\
	PipeName  string `yaml:"pipe_name" mapstructure:"pipe_name" default:"default"` // This is the pipe Name where it connects (or the URI)
	AuthToken string `yaml:"auth_token" mapstructure:"auth_token" default:"default_token"`
	//----------RELATED TO BROKER---------\\

	//

	//----------RELATED TO CONNECTIONS---------\\
	// These configurations are taken from websocket! so they should be the same!
	// Hosts -> They are comma separated if there are multiple!
	AutoReconnect string `yaml:"auto_reconnect" mapstructure:"auto_reconnect" default:"yes"`
	Reconnect     ReconnectOptions
	// UseMultipleConnections -> Enable utilization of existing listed connections, if the client cannot connect to previous, will take the next one!
	UseMultipleConnections string `yaml:"use_multiple_connections" mapstructure:"use_multiple_connections" default:"yes"`
	Connections            []*connection.Connection
	//----------RELATED TO CONNECTIONS---------\\

	// This is the logger configuration!
	Logger loggerConfig.Config
}
