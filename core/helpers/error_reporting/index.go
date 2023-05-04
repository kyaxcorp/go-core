package error_reporting

import (
	"github.com/getsentry/sentry-go"
	"go.szostok.io/version"
	"log"
	"time"
)

var isInitialized = false

type Options struct {
	DSN   string
	Debug bool

	TracesSampleRate float64
}

func New(o Options) {
	v := version.Get()
	dsn := o.DSN
	tracesSampleRate := o.TracesSampleRate
	if tracesSampleRate == 0 {
		tracesSampleRate = 0.2
	}

	//if working_stage.IsDev() {
	//	dsn = o.DevDSN
	//	if dsn == "" {
	//		dsn = o.ProdDSN
	//	}
	//} else {
	//	dsn = o.ProdDSN
	//	if dsn == "" {
	//		dsn = o.DevDSN
	//	}
	//}

	if dsn == "" {
		panic("no DSN defined for sentry")
	}

	err := sentry.Init(sentry.ClientOptions{
		// Specify a fixed sample rate:
		TracesSampleRate: tracesSampleRate,
		// Or provide a custom sampler:
		/*TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			return sentry.SampledTrue
		}),*/

		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		//Environment: "",
		//Release: v. + "@" + v.Version,
		// TODO: add PROJECT NAME
		Release: "@" + v.Version,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: o.Debug,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
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

		// TODO: turn on/off printing the stacktrace...
		// TODO: print the stack?!
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
