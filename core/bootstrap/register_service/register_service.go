package register_service

import (
	"github.com/KyaXTeam/go-core/v2/core/logger/appLog"
	"github.com/rs/zerolog"
	"sync"
)

// Services -> Here we will be storing the registered services
var Services = make(map[string]RegisteredService)
var servicesLock sync.Mutex

type RegisteredService interface {
	Run()
	Stop()
	// TODO: maybe other methods should be added for services
}

func RegisterService(
	name string,
	service RegisteredService,
) {
	servicesLock.Lock()
	defer servicesLock.Unlock()
	Services[name] = service
}

// RunRegisteredServices -> this will run the registered services like bootstrap client or others...
func RunRegisteredServices() {

	// I should run here a Method which will be in the interface...
	// This method should be a standard...
	// The objects should be stored somewhere!..?!

	// Start the broker clients
	/*b.brokerClients = broker.GenerateAllClients(_context.GetRootContext())

	// TODO: do good logging
	for instanceName, brokerClient := range b.brokerClients {
		appLog.Info().Msg("starting broker client " + instanceName)
		brokerClient.Connect()
	}*/
	info := func() *zerolog.Event {
		return appLog.InfoF("RunRegisteredServices")
	}
	/*info().Msg("entering...")
	defer info().Msg("leaving...")*/

	for serviceName, _service := range Services {
		info().Str("service_name", serviceName).Msg("running service...")
		// Calling the Run Method of that service
		_service.Run()
	}
}
