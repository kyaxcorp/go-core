package gor

import (
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"log"
	"time"
)

func (g *GInstance) runCallback() {
	// Check if the app or parent object is terminating...
	if g.IsTerminating() {
		// We will not run...
		return
	}

	if g.isCallbackRunning.IfFalseSetTrue() {
		// Already running...
		return
	}

	// If you'll ask: why we are not self healing or self turning on the process
	// when the goroutine dies, well the answer is that: i'm not sure if this from
	// arhitecture point of view it's a good way to do it, i mean to create self nested goroutine creation...
	// maybe it's a easy correct way, and that's the point of go, but, for precautions, i created a monitoring,
	// which handles the turning from there...
	// also, in future we can add multiple goroutines running as a pool here...?

	// This is the mark which shows us if the entire planned process has being finished
	g.isAllFinished.False()

	// Run as a separate goroutine!
	go func() {
		// Do some checks before leaving...
		defer func() {
			// Set that is not running anymore
			// Set that callback is not running anymore...
			g.isCallbackRunning.False()

			// Send notification that we are not running anymore! The monitoring will receive it as a trigger!
			g.notRunning <- true
			// Check if there was a panic... if yes, then try recovering from it!
			if r := recover(); r != nil {
				// fmt.Println("Recovered in f", r)
				// TODO: we should use a global function which will gather this logs in 1 location
				log.Println("gor function run panic -> recovered", r)
			}

			if function.IsCallable(g.config.OnRunFinished) {
				g.config.OnRunFinished(g)
			}
		}()

		// Set the initial counter of the running times
		runTimes := 0
		runPanics := 0

		// Start looping and calling the OnRun
		for {
			runTimes++

			isFinished := make(chan bool)
			// Run in a separate goroutine because it can contain panics
			// So we need to ensure it's isolated and it can recover
			// Also we want to have logic well arranged
			// We have defined that runClosure is a function because later on we should
			// Call this function from inside...
			var runClosure func()
			// Set the function
			runClosure = func() {
				defer func() {
					// TODO: should we add here some events like finished, or onPanic?
					// TODO: should it be a function attached to structure?!

					// Notify that has finished!

					if r := recover(); r != nil {
						runPanics++
						// TODO: we should use a global function which will gather this logs in 1 location
						log.Println("gor function run 'time' panic -> recovered", r)

						// Check if ReRun On Recovery is On
						if conv.ParseBool(g.config.ReRunOnRecover) {
							// TODO: how many times already had panics
							time.Sleep(g.config.ReRunOnRecoverOptions.RunAfterDuration)
							// Check if it's not infinite and it's higher than 0

							// set that it can't rerun...
							canRunAgain := false
							// Check if it's infinite or higher than 0...
							// also check if it's not higher than the specific max number of panics
							if g.config.ReRunOnRecoverOptions.MaxNrOfPanics == -1 ||
								(g.config.ReRunOnRecoverOptions.MaxNrOfPanics > 0 &&
									runPanics <= g.config.ReRunOnRecoverOptions.MaxNrOfPanics) {
								canRunAgain = true
							}
							// Check if can run again
							if canRunAgain {
								// Run itself again!
								go runClosure()
								// exit this closure, and don't say that has finished yet!
								return
							}
						}
					}
					// Set that it has finished the processing...
					isFinished <- true
				}()

				// Check if the OnRun function is callable!
				if function.IsCallable(g.config.OnRun) {
					// Run the OnRun function
					g.config.OnRun(g)
				}
			}
			// Run the Closure in a separate goroutine
			// we run it separately because of panics
			// we also have introduced some recovery mechanism for repeated strategy
			go runClosure()
			// Wait until finished
			<-isFinished

			// Check if it has successfully finished or not...

			// -1 it's infinite -> it will run infinitely only if not terminated
			// if it's higher than 0 then, till a specific nr of times...
			if g.config.RunTimes >= runTimes && g.config.RunTimes != -1 {
				// Set that the process has being finished successfully as planned
				g.isAllFinished.True()
				// Break from loop!
				break
			}
			// if it's terminating, the app or the parent
			if g.IsTerminating() {
				// Break from loop!
				break
			}
		}

		// If all finished as planned
		if g.isAllFinished.Get() {
			// Check if the callback has being set
			if function.IsCallable(g.config.OnAllFinished) {
				// Run the callback
				g.config.OnAllFinished(g)
			}
		}

	}()
}
