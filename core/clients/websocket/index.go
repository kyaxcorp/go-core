package websocket

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/client"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/config"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/instances"
	mainConfig "github.com/KyaXTeam/go-core/v2/core/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err"
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
		return nil, err.New(0, "websocket client instance name is empty")
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
		return nil, err.New(0, "websocket client instance name is empty")
	}
	if cfg, _err := GetConfig(instanceName); _err == nil {
		return New(ctx, cfg)
	}
	return nil, err.New(0, "websocket client configuration missing")
}

func GetConfig(instanceName string) (config.Config, error) {
	if cfg, ok := mainConfig.GetConfig().Clients.WebSocket.Instances[instanceName]; ok {
		return cfg, nil
	}
	return config.Config{}, err.New(0, "websocket client config doesn't exist")
}
