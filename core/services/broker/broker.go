package broker

import (
	"context"
	mainConfig "github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/err"
)

// What is Broker?!
// It's an intermediary between applications
// It works as a separate launched service
//
// Communication between different apps can be done through it!
// It works on the principle of: creating hubs, subscribing to hubs, and listening for messages!
// Each connected client can be a subscriber and a listener!
// Different mechanisms can be created as overlays on it!
// It works based on websocket!
// Multiple Brokers can be launched, they are all defined in the configuration file .yaml
// broker:start <BROKER NAME>
// if no name indicated, then the default one it's being taken

/*
Multiple brokers can be launched
Each broker has same functionality
The only difference it's the load balancing!
*/

// GetDefaultBroker it gets the default broker from the config!
func GetDefaultBroker() (*Broker, error) {
	instanceName := mainConfig.GetConfig().Services.Broker.DefaultInstanceName
	// TODO: get default context from the system!? should be declared somewhere?!
	return GetBroker(_context.GetDefaultContext(), instanceName)
}

// GetBroker it gets a specific broker by name!
func GetBroker(ctx context.Context, brokerName string) (*Broker, error) {
	if brokerName == "" {
		return nil, err.New(0, "broker name is empty")
	}
	if cfg, ok := mainConfig.GetConfig().Services.Broker.Instances[brokerName]; ok {
		//cfg.ListeningAddresses
		return New(ctx, brokerName, cfg)
	}
	return nil, err.New(0, "broker configuration missing")
}
