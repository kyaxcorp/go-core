package shutdown

import (
	"context"
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/kyaxcorp/go-core/core/logger"
	"github.com/kyaxcorp/go-core/core/logger/application/vars"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//var WaitFinished = make(chan bool)
// TODO: should be set here the default context instead of the background?!
var WaitFinished, WaitCancel = context.WithCancel(context.Background())

type OnShutdownCb func()

var onShutdownItems = []OnShutdownCb{}

func CallGracefulShutdown() {
	// We should tell to the global context that the application is shutting down!
	// context with 30s timeout
	_context.Cancel()
	// Running shutdown in a separate routine
	if len(onShutdownItems) > 0 {
		for _, onShutdown := range onShutdownItems {
			if function.IsCallable(onShutdown) {
				go onShutdown()
			}
		}
	}

	waitTime := config.GetConfig().Application.OnShutdownWaitSeconds
	if vars.ApplicationLogger != nil {
		logger.GetAppLogger().Info().Int("shutdown_wait_time", waitTime).Msg("waiting processes to finish")
	}
	time.Sleep(time.Second * time.Duration(waitTime))
	if vars.ApplicationLogger != nil {
		logger.GetAppLogger().Info().Msg("shutting down...")
	}
	//WaitFinished <- true
	WaitCancel()
	// Exiting the application!
	//os.Exit(0)
}

/*
func GracefulShutdown(
	callback func() error,
	teardown func(context.Context) error,
) error {
	term := make(chan os.Signal) // OS termination signal
	fail := make(chan error)     // Teardown failure signal

	go func() {
		signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
		<-term // waits for termination signal

		// context with 30s timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// all teardown process must complete within 30 seconds
		fail <- teardown(ctx)
	}()

	// listenAndServe blocks our code from exit, but will produce ErrServerClosed when stopped
	if err := callback(); err != nil && err != http.ErrServerClosed {
		return err
	}

	// after server gracefully stopped, code proceeds here and waits for any error produced by teardown() process @ line 26
	return <-fail
}*/

func OnShutdown(onShutdown ...OnShutdownCb) {
	onShutdownItems = append(onShutdownItems, onShutdown...)
}

// MonitorSigMessages -> receives the signal of termination, and reacts based on this
// It should call the Cancel Context, and the entire app should terminate gracefully
func MonitorSigMessages() {
	term := make(chan os.Signal) // OS termination signal
	// fail := make(chan error)     // Teardown failure signal

	go func() {
		signal.Notify(
			term,
			syscall.SIGINT,  // CTRL+C
			syscall.SIGTERM, // 15 Gracefull shutdown
		)

		<-term // waits for termination signal
		CallGracefulShutdown()
	}()
}
