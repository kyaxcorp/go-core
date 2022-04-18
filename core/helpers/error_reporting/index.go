package error_reporting

import (
	"github.com/getsentry/sentry-go"
	"github.com/kyaxcorp/go-core/core/helpers/version"
	"log"
	"time"
)

var isInitialized = false

func New(dsn string, debug bool) {
	v := version.GetAppVersion()
	e := sentry.Init(sentry.ClientOptions{
		// Specify a fixed sample rate:
		TracesSampleRate: 0.2,
		// Or provide a custom sampler:
		/*TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			return sentry.SampledTrue
		}),*/

		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		//Environment: "",
		Release: v.ProjectName + "@" + v.Version,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: debug,
	})
	if e != nil {
		log.Fatalf("sentry.Init: %s", e)
		// Exit...
	}
	isInitialized = true
}

func Recover() *sentry.EventID {
	//sentry.Recover()
	if err := recover(); err != nil {
		hub := sentry.CurrentHub()
		return hub.Recover(err)
	}
	return nil
}

func BeforeShutdown() {
	//Recover()
	if err := recover(); err != nil {
		log.Println("------ERROR/PANIC------")
		log.Println(err)

		// If you want to print stack trace try using:
		// https://github.com/go-errors/errors
		// or sentry
		//log.Println("------STACK TRACE------")

		// fmt.Println(errors.Wrap(err, 2).ErrorStack())
		//errors.Wrap(err, 2).ErrorStack()
		hub := sentry.CurrentHub()
		hub.Recover(err)
	}
	Flush()
}

func Flush() {
	if !isInitialized {
		return
	}
	sentry.Flush(2 * time.Second)
}

func CaptureException(exception error) {
	if !isInitialized {
		//log.Println(isInitialized)
		return
	}
	sentry.CaptureException(exception)
}
