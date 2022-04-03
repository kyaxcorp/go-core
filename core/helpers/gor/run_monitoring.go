package gor

import (
	"github.com/rs/zerolog"
	"log"
	"time"
)

// runMonitoring -> it turns on the monitoring routine
func (g *GInstance) runMonitoring() {
	// We still need here a monitorer! if it still dies, the monitorer it will
	// Revive the dead thread!
	// Also check if the monitoring is running!
	// Can the run monitoring panic?! somewhere in logs?

	// Check if it's terminating...(by context)
	// if yes, then we cannot run the monitoring
	if g.IsTerminating() {
		return
	}

	// When exiting this function, do some checks...
	defer func() {
		// Let's set anyway that monitoring is not running...
		// If it's running, then don't touch!
		// g.isMonitoringRunning.IfTrueSetFalse()

		// Check if no panic went out, and if there is one,
		// let's recover from it
		if r := recover(); r != nil {
			// Yes, it had a panic, let's run again the monitoring
			// We will not make any logs here... but we can write on screen!?
			// We will run the monitoring again in a goroutine! but first let's sleep a bit! Half a second
			// TODO: we should use a global function which will gather this logs in 1 location

			log.Println("gor monitoring panic -> recovered", r)

			// Check if the app it's not terminating
			if !g.IsTerminating() {
				// Sleep a bit... and then rerun the monitoring
				time.Sleep(time.Millisecond * 500)
				// Run the monitoring again...
				go g.runMonitoring()
			}
		}
	}()

	// Declare the 'debug' log function
	_debug := func() *zerolog.Event {
		return g.LDebugF("runMonitoring")
	}
	/*_error := func() *zerolog.Event {
		return g.LErrorF("runMonitoring")
	}*/

	_debug().Msg("entering...")
	defer _debug().Msg("leaving...")

	// Run the monitoring in a separate goroutine
	go func() {
		// Check if it's running already, and if not, then Set as running
		// If it's running, it will exit the goroutine
		if g.isMonitoringRunning.IfFalseSetTrue() {
			_debug().Msg("monitoring is already running")
			return
		}

		// Create a defer function which will run when exiting the goroutine
		// We need to perform some checkups before exiting like:
		// - were there any panics?
		// - ...
		defer func() {
			// Set that monitoring is not running anymore...because we need to rerun itself
			g.isMonitoringRunning.False()
			// we should turn on the monitoring only if it died
			// we should not turn on the monitoring if the context cancel has being called
			if r := recover(); r != nil {
				log.Println("gor monitoring died from panic", r)
				// Check if the app is not terminating
				if !g.IsTerminating() {
					// Sleep a bit... and then rerun the monitoring from the beginning
					time.Sleep(time.Millisecond * 500)
					go g.runMonitoring()
				}
			}
		}()

		// Check if it's not terminating
		// Check if main goroutine it's not running already
		if !g.IsTerminating() && !g.isCallbackRunning.Get() {
			// Turn on for the first time in a separate goroutine!
			go g.runCallback()
		}

		// Enter loop
		for {
			// Check if the app is not terminating...
			if g.IsTerminating() {
				// Stop entirely if it's terminating...
				break
			}

			// We do select because it's the correct way of communicating with other goroutines and
			// to not load the system and lose performance and time by using sleep!
			select {
			// Maybe receive a channel bool from callback as defer if died, and react based on it?!
			case <-g.ctx.Context().Done():
				// The inside scope is Terminating, meaning the client should stop!
				break
			case <-g.parentCtx.Done():
				// the app or the parent object  is terminating
				break
			case notRunning := <-g.notRunning:
				// TODO: check if it should be run!
				// TODO: check if it's has being finished, if yes then we should not run anymore!

				// If it's not running anymore and the process hasn't finished well, we will run it again!
				if notRunning {
					// if it's not finished yet...
					if !g.isAllFinished.Get() {
						// Turn it on!
						go g.runCallback()
					} else {
						// it has finished
						// we should turn off everything?!
						// TODO: what else should we do here? cleanup everything?
						// TODO: what if later on we will want to rerun the object
						// TODO: set itself as terminating?! or simply break?
					}
				}
			}
		}
	}()
}
