package gor

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
)

// Run -> start the goroutine
func (g *GInstance) Run() error {
	// It should monitor the create goroutine...
	// we can do it by:
	// 1. creating some variables.
	// 2. using defer func
	// 3. add panic catcher!

	// If stop function is running, we should wait for it to finish
	if g.isStopFuncRunning.Get() {
		return ErrStopFunctionAlreadyRunning
	}

	// if not running, set that is running ands
	if g.isRunFuncRunning.IfFalseSetTrue() {
		return ErrRunFunctionAlreadyRunning
	}

	// If it's not running set to true!
	if g.isMonitoringRunning.Get() {
		return ErrAlreadyRunning
	}

	// do on leaving...
	defer func() {
		// Set that is not running anymore!
		g.isRunFuncRunning.False()
	}()

	// Create the cancel context which will stop the goroutine!
	g.ctx = _context.WithCancel(g.parentCtx)
	// Run the monitoring...
	g.runMonitoring()

	return nil
}
