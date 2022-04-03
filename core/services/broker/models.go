package broker

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/server"
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	brokerConfig "github.com/KyaXTeam/go-core/v2/core/services/broker/config"
	"sync"
)

type Broker struct {
	// Name   string
	Server *server.Server         // This is the websocket server
	Pipes  map[string]*server.Hub // These are the pipes/hubs that are being created

	// protected
	pipesLock             sync.RWMutex
	shutdownHubMonitoring chan bool
	isStarted             *_bool.Bool
	isStarting            *_bool.Bool
	isStopping            *_bool.Bool
	config                brokerConfig.Config

	// Context
	parentCtx context.Context
	ctx       *_context.CancelCtx

	// Logger
	Logger *model.Logger
}
