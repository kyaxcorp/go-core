package config

import (
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
)

/*
maybe we should also add other broker settings like:
- status http
- channels inside the broker
- credentials?!...

*/
// TODO: maybe later we can add pipes authentication and isolation!?
type Config struct {
	AuthToken             string   `yaml:"auth_token" mapstructure:"auth_token" default:"default_token"`
	ListeningAddresses    []string `yaml:"listening_address" mapstructure:"listening_address"`
	ListeningAddressesSSL []string `yaml:"listening_address_ssl" mapstructure:"listening_address_ssl"`
	IsListenPlain         string   `yaml:"is_listen_plain" mapstructure:"is_listen_plain" default:"no"` // Don't listen on 30000
	IsListenSSL           string   `yaml:"is_listen_ssl" mapstructure:"is_listen_ssl" default:"yes"`    // Listen by default Secured! -> 30001
	// SubscribeToClusterNodes -> the server subscribes to another nodes from the cluster, and they will receive messages
	// on this channel! When they subscribe, they send metadata about themself!
	SubscribeToClusterNodes string `yaml:"subscribe_to_cluster_nodes" mapstructure:"subscribe_to_cluster_nodes" default:""`
	Logger                  loggerConfig.Config
}
