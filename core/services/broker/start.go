package broker

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
	"github.com/rs/zerolog"
)

// TODO: see what we are returning here!
// Start -> it can be launched with goroutine if needed
func (b *Broker) Start() error {
	info := func() *zerolog.Event {
		return b.LInfoF("Start")
	}
	warn := func() *zerolog.Event {
		return b.LWarnF("Start")
	}
	_error := func() *zerolog.Event {
		return b.LErrorF("Start")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")

	if b.isStopping.Get() {
		warn().Msg("already stopping...")
		return define.Err(0, "already stopping")
	}
	// Checking if it's already running
	if b.isStarted.Get() {
		warn().Msg("already started...")
		return define.Err(0, "already started")
	}
	if b.isStarting.IfFalseSetTrue() {
		warn().Msg("already starting...")
		return define.Err(0, "already starting")
	}

	// Creating broker's context!
	b.ctx = _context.WithCancel(b.parentCtx)
	// Inheriting context to childs
	b.Server.SetContext(b.ctx.Context())

	// Clean before starting... maybe previously we had something left...
	b.Clean()

	info().Msg("starting websocket server...")
	// Starting the HTTP Websocket Server
	_err := b.Server.Start()
	if _err != nil {
		_error().Msg("failed to start websocket server")
	}

	// Start hubs monitoring and other things...
	b._start()

	// Set as Started!
	b.isStarted.True()
	// Set that not starting...
	b.isStarting.False()
	return nil
}

func (b *Broker) _start() {
	info := func() *zerolog.Event {
		return b.LInfoF("_start")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")
	//
	go b.hubsMonitoring()
}
