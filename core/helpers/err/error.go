package err

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
	"github.com/KyaXTeam/go-core/v2/core/helpers/error_reporting"
	"github.com/KyaXTeam/go-core/v2/core/helpers/function"
	"github.com/KyaXTeam/go-core/v2/core/logger/application/vars"
)

var report = true
var reportHandlerSet = false
var callbackToReport func(e *define.GeneralError) bool

// New error
func New(code int, message ...interface{}) error {
	e := define.Err(code, message...)
	NewCustom(CustomError{
		// It will log to app logs!
		LogToApp:     true,
		GeneralError: e,
	})
	return e
}

// New error
func NewDefined(definedErr *define.GeneralError) error {
	return New(definedErr.Code, definedErr.Err.Error())
}

type CustomError struct {
	Code    int
	Message string
	// Log to application Logger
	LogToApp bool
	// Set instantly this error if you want to override the code & message
	GeneralError *define.GeneralError
}

func NewCustom(customError CustomError) error {
	var e *define.GeneralError
	if customError.GeneralError != nil {
		// Get the defined one
		e = customError.GeneralError
	} else {
		// Compose one
		e = define.Err(customError.Code, customError.Message)
	}
	//go func() {
	// Take the general value
	toReport := report
	// Check if there is a handler for this and take it if there is!
	if reportHandlerSet {
		toReport = callbackToReport(e)
	}

	// Check if object is created...because when the app initially starts, it depends on the configuration file!
	if customError.LogToApp && vars.ApplicationLogger != nil {
		vars.ApplicationLogger.Error().Err(e).Msg("")
		//debug.PrintStack()
	}
	//log.Println(toReport)
	if toReport {
		// Report
		// log.Println("reporting", toReport)
		error_reporting.CaptureException(e)
	}
	//}()
	return e
}

func DisableReport() {
	report = false
}

func EnableReport() {
	report = false
}

// BeforeReport decides to report or not the error to the error reporting manager
func BeforeReport(report func(e *define.GeneralError) bool) {
	if function.IsCallable(report) {
		reportHandlerSet = true
		callbackToReport = report
	}
}

// NewReport -> This will also report!
/*func NewReport(code int, message string) error {
	e := New(code, message)
	error_reporting.CaptureException(e)
	return e
}*/
