package broker

import (
	"context"
	"github.com/kyaxcorp/go-core/core/clients/broker/client"
	"github.com/kyaxcorp/go-core/core/clients/broker/config"
	"github.com/kyaxcorp/go-core/core/clients/broker/instances"
	mainConfig "github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/err"
)

// New -> creates a new client based on the provided configuration
func New(
	ctx context.Context,
	config config.Config,
) (*client.Client, error) {
	return client.New(ctx, config)
}

func GetDefaultClient() (*client.Client, error) {
	connName := mainConfig.GetConfig().Clients.Broker.DefaultInstanceName
	return GetClient(_context.GetDefaultContext(), connName)
}

func GetClient(ctx context.Context, instanceName string) (*client.Client, error) {
	if instanceName == "" {
		return nil, err.New(0, "broker client instance name is empty")
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
		return nil, err.New(0, "broker client instance name is empty")
	}
	if cfg, _err := GetConfig(instanceName); _err == nil {
		return New(ctx, cfg)
	}
	return nil, err.New(0, "broker client configuration missing")
}

func GetConfig(instanceName string) (config.Config, error) {
	if cfg, ok := mainConfig.GetConfig().Clients.Broker.Instances[instanceName]; ok {
		return cfg, nil
	}
	return config.Config{}, err.New(0, "broker client config doesn't exist")
}

// GenerateAllClients -> it's being used only for bootstrap!
func GenerateAllClients(ctx context.Context) map[string]*client.Client {
	_instances := mainConfig.GetConfig().Clients.Broker.Instances
	_clients := make(map[string]*client.Client)
	for instanceName, instanceCfg := range _instances {
		if !conv.ParseBool(instanceCfg.IsEnabled) {
			// Skipping which is not enabled
			continue
		}

		_client, _err := GetClient(ctx, instanceName)
		if _err != nil {
			continue
		}
		_clients[instanceName] = _client
	}
	return _clients
}
