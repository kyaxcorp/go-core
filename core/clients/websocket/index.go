package websocket

import (
	"context"
	"github.com/kyaxcorp/go-core/core/clients/websocket/client"
	"github.com/kyaxcorp/go-core/core/clients/websocket/config"
	"github.com/kyaxcorp/go-core/core/clients/websocket/instances"
	mainConfig "github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/errors2"
)

// New -> creates a new client based on the provided configuration
func New(
	ctx context.Context,
	config config.Config,
) (*client.Client, error) {
	return client.New(ctx, config)
}

// GetDefaultClient -> it will return a new websocket client instance by taking the default configuration
func GetDefaultClient() (*client.Client, error) {
	instanceName := mainConfig.GetConfig().Clients.WebSocket.DefaultInstanceName
	return GetClient(_context.GetDefaultContext(), instanceName)
}

// GetClient -> gets the existing created instance if it's already created... if it's not created, then it's being created
func GetClient(ctx context.Context, instanceName string) (*client.Client, error) {
	if instanceName == "" {
		return nil, errors2.New(0, "websocket client instance name is empty")
	}
	// check if there is an existing instance
	srv, _err := instances.GetInstance(instanceName)
	if _err == nil {
		return srv, nil
	}
	srv, _err = NewClient(ctx, instanceName)
	if _err != nil {
		return nil, _err
	}
	// Save the instance
	instances.SaveInstance(instanceName, srv)
	return srv, nil
}

// NewClient -> Generates a new instance based on the configuration!
// Multiple instances can be generated based on the same configuration
func NewClient(ctx context.Context, instanceName string) (*client.Client, error) {
	if instanceName == "" {
		return nil, errors2.New(0, "websocket client instance name is empty")
	}
	if cfg, _err := GetConfig(instanceName); _err == nil {
		return New(ctx, cfg)
	}
	return nil, errors2.New(0, "websocket client configuration missing")
}

func GetConfig(instanceName string) (config.Config, error) {
	if cfg, ok := mainConfig.GetConfig().Clients.WebSocket.Instances[instanceName]; ok {
		return cfg, nil
	}
	return config.Config{}, errors2.New(0, "websocket client config doesn't exist")
}
