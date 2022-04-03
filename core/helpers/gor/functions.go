package gor

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
)

func (g *GInstance) IsRunning() bool {
	// If any of them is running, it means that is still running!
	return g.isCallbackRunning.Get() || g.isMonitoringRunning.Get()
}

// IsTerminating -> usually it's necessary for the callback, to know
// if the goroutines it's being declared to terminate or the parents!
func (g *GInstance) IsTerminating() bool {
	// We simply create a temporary context to not be blocked by using directly
	tmpContext := _context.WithCancel(g.parentCtx)
	if tmpContext.IsDone() {
		return true
	}
	if g.ctx.IsDone() {
		return true
	}
	return false
}

func (g *GInstance) GetContext() context.Context {
	return g.ctx.Context()
}
