package _context

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/function"
)

type CancelCtx struct {
	ctx    context.Context
	cancel context.CancelFunc
}

var rootCtx *CancelCtx
var backgroundCtx context.Context

func (c *CancelCtx) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *CancelCtx) IsDone() bool {
	select {
	case <-c.Done():
		return true
	default:
		return false
	}
}

// Cancel Signal!
func (c *CancelCtx) Cancel() {
	if function.IsCallable(c.cancel) {
		c.cancel()
	}
}

// Context returns the context itself
func (c *CancelCtx) Context() context.Context {
	return c.ctx
}

// GetDefaultContext Get the ROOT Context
func GetDefaultContext() context.Context {
	// TODO: add some hooks to overrride the default context!!!
	return GetRootContext()
}

func GetRootContext() context.Context {
	if rootCtx == nil {
		backgroundCtx = context.Background()
		rootCtx = WithCancel(backgroundCtx)
	}
	return rootCtx.Context()
}

// Cancel This will terminate entirely the application where context it's being used!
func Cancel() {
	rootCtx.Cancel()
}

// WithCancel Parent ctx can be nil... if so, then we will set the main default context!
// We also return a simple structure of context
func WithCancel(parentCtx context.Context) *CancelCtx {
	if parentCtx == nil {
		parentCtx = GetDefaultContext()
	}
	// Also we can add here different hooks or intersting things...
	newCtx, cancelFunc := context.WithCancel(parentCtx)
	return &CancelCtx{
		ctx:    newCtx,
		cancel: cancelFunc,
	}
}

// TODO: create different additional contexts, which can help us handle application from different standpoints...
// like Internet connectivity, background work, filesystem work etc...
