package client

import (
	"context"
	"github.com/kyaxcorp/go-core/core/clients/broker/config"
	wsClient "github.com/kyaxcorp/go-core/core/clients/websocket/client"
	wsConfig "github.com/kyaxcorp/go-core/core/clients/websocket/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_bool"
	"github.com/kyaxcorp/go-core/core/logger/model"
)

type Client struct {
	// Broker Configuration
	config config.Config
	// WebSocket Configuration
	wsConfig wsConfig.Config

	// WebSocket Client which will be created on connect
	WSClient *wsClient.Client

	// Logger
	Logger *model.Logger

	//-----------CONTEXT-----------\\
	// This is for stopping entirely the client... meaning to disconnect, stop the connection process etc...!
	parentCtx context.Context
	ctx       *_context.CancelCtx
	// THis is for canceling the connection process
	ctxConnect *_context.CancelCtx
	// Here we just set if it's done or not!
	ctxDone *_bool.Bool
	//-----------CONTEXT-----------\\
}
