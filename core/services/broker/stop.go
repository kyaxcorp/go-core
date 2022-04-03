package broker

import (
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/rs/zerolog"
)

func (b *Broker) Stop() error {
	info := func() *zerolog.Event {
		return b.LInfoF("Stop")
	}
	warn := func() *zerolog.Event {
		return b.LWarnF("Stop")
	}
	_error := func() *zerolog.Event {
		return b.LErrorF("Stop")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")

	// Check if it's not starting right now...
	if b.isStarting.Get() {
		warn().Msg("already starting...")
		return define.Err(0, "already starting...")
	}
	// Check if is started already ...
	if !b.isStarted.Get() {
		// It's not started, exiting...
		warn().Msg("already started...")
		return define.Err(0, "already started...")
	}

	// Check if is stopping right now... if not then stop!
	if b.isStopping.IfFalseSetTrue() {
		warn().Msg("already stopping...")
		return define.Err(0, "already stopping...")
	}

	// Stopping all other processes of the broker...
	b._stop()

	// Stop the the HTTP Websocket Server
	_err := b.Server.Stop() //
	if _err != nil {
		_error().Err(_err).Msg("failed to stop websocket server")
		// TODO: what do we do here?!!!
	}

	// Set that is not stopping anymore!
	b.isStopping.False()
	// Set that the server is not started anymore!
	b.isStarted.False()
	return nil
}

func (b *Broker) _stop() {
	b.shutdownHubMonitoring <- true // Sending the shutdown signal!
}
