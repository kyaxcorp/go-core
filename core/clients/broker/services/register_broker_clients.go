package services

import (
	"github.com/kyaxcorp/go-core/core/bootstrap/register_service"
	"github.com/kyaxcorp/go-core/core/clients/broker"
	"github.com/kyaxcorp/go-core/core/clients/broker/client"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/logger/appLog"
)

type BrokerClientsService struct {
	brokerClients map[string]*client.Client
}

func (s BrokerClientsService) Run() {
	// TODO: do good logging
	for instanceName, brokerClient := range s.brokerClients {
		appLog.Info().Msg("starting broker client " + instanceName)
		brokerClient.Connect()
	}
}

func (s BrokerClientsService) Stop() {

}

// TODO: maybe other methods should be added for services

func RegisterBrokerService() {
	// I should run here a Method which will be in the interface...
	// This method should be a standard...
	// The objects should be stored somewhere!..?!

	// Start the broker clients

	brokerService := BrokerClientsService{
		brokerClients: broker.GenerateAllClients(_context.GetDefaultContext()),
	}
	register_service.RegisterService("broker_clients", brokerService)
}
