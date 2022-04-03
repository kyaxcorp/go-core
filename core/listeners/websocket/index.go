package websocket

import (
	"context"
	mainConfig "github.com/KyaXTeam/go-core/v2/core/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err"
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/config"
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/instances"
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/server"
)

// New -> creates a new server based on the provided configuration
func New(
	ctx context.Context,
	config config.Config,
) (*server.Server, error) {
	return server.New(ctx, config)
}

func GetDefaultServer() (*server.Server, error) {
	instanceName := mainConfig.GetConfig().Listeners.WebSocket.DefaultInstanceName
	return GetServer(_context.GetDefaultContext(), instanceName)
}

// GetServer -> gets the existing created instance if it's already created... if it's not created, then it's being created
func GetServer(ctx context.Context, instanceName string) (*server.Server, error) {
	if instanceName == "" {
		return nil, err.New(0, "websocket server instance name is empty")
	}
	// check if there is an existing instance
	srv, _err := instances.GetInstance(instanceName)
	if _err == nil {
		return srv, nil
	}
	srv, _err = NewServer(ctx, instanceName)
	if _err != nil {
		return nil, _err
	}
	// Save the instance
	instances.SaveInstance(instanceName, srv)
	return srv, nil
}

// NewServer -> Generates a new instance based on the configuration!
// Multiple instances can be generated based on the same configuration
func NewServer(ctx context.Context, instanceName string) (*server.Server, error) {
	if instanceName == "" {
		return nil, err.New(0, "instance name is empty")
	}
	if cfg, _err := GetConfig(instanceName); _err == nil {
		return New(ctx, cfg)
	}
	return nil, err.New(0, "websocket server configuration missing")
}

func GetConfig(instanceName string) (config.Config, error) {
	if cfg, ok := mainConfig.GetConfig().Listeners.WebSocket.Instances[instanceName]; ok {
		return cfg, nil
	}
	return config.Config{}, err.New(0, "websocket server config doesn't exist")
}
