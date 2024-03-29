package http

import (
	"context"
	mainConfig "github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/errors2"
	"github.com/kyaxcorp/go-core/core/listeners/http/config"
	"github.com/kyaxcorp/go-core/core/listeners/http/instances"
	"github.com/kyaxcorp/go-core/core/listeners/http/server"
)

// New -> creates a new server based on the provided configuration
func New(
	ctx context.Context,
	config config.Config,
) (*server.Server, error) {
	return server.New(ctx, config)
}

func GetDefaultServer() (*server.Server, error) {
	instanceName := mainConfig.GetConfig().Listeners.Http.DefaultInstanceName
	return GetServer(_context.GetDefaultContext(), instanceName)
}

// GetServer -> gets the existing created instance if it's already created... if it's not created, then it's being created
func GetServer(ctx context.Context, instanceName string) (*server.Server, error) {
	if instanceName == "" {
		return nil, errors2.New(0, "http server instance name is empty")
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
		return nil, errors2.New(0, "instance name is empty")
	}
	if cfg, _err := GetConfig(instanceName); _err == nil {
		return New(ctx, cfg)
	}
	return nil, errors2.New(0, "http server configuration missing")
}

func GetConfig(instanceName string) (config.Config, error) {
	if cfg, ok := mainConfig.GetConfig().Listeners.Http.Instances[instanceName]; ok {
		return cfg, nil
	}
	return config.Config{}, errors2.New(0, "http server config doesn't exist")
}
