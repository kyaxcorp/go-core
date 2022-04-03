package gor

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_bool"
	"github.com/kyaxcorp/go-core/core/logger/model"
)

type GInstance struct {
	// If this is the callback that's running
	isCallbackRunning *_bool.Bool
	// This is the notification which the OnRun will send when it died or not running anymore
	// so the monitoring will turn it on again!
	notRunning chan bool
	// This is the current status of monitoring
	isMonitoringRunning *_bool.Bool
	// here we store the received config with default values
	config Config
	// This is the parent context, given by the user or the root Context...
	parentCtx context.Context
	// This is the Client Context which cancels everything in it! (this is a private one used only in the current scope)
	ctx *_context.CancelCtx
	// this var defines if the process has being finished as planned
	isAllFinished *_bool.Bool

	// Logger
	Logger *model.Logger

	// Locks
	isRunFuncRunning  *_bool.Bool
	isStopFuncRunning *_bool.Bool
}
