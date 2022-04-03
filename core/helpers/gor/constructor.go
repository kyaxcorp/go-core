package gor

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_struct"
	"github.com/KyaXTeam/go-core/v2/core/helpers/conv"
	"github.com/KyaXTeam/go-core/v2/core/helpers/function"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/logger"
	"github.com/rs/zerolog"
)

// GO -> run a goroutine
func GO(ctx context.Context, config Config) (*GInstance, error) {
	// Creating the functions for logging

	_logger := config.Logger
	if _logger == nil {
		_logger = logger.GetAppLogger()
	}

	_debug := func() *zerolog.Event {
		return _logger.DebugF("new goroutine")
	}
	_error := func() *zerolog.Event {
		return _logger.ErrorF("new goroutine")
	}

	// Checking if the OnRun it's being defined...
	// If there is no OnRun, it can't run further...
	if !function.IsCallable(config.OnRun) {
		_error().Err(ErrOnRunIsNotAFunction).Msg("")
		return nil, ErrOnRunIsNotAFunction
	}

	_debug().Msg("setting defaults values to config")
	// Set defaults to config which are missing
	_err := _struct.SetDefaultValues(&config)
	if _err != nil {
		// If there was an error setting the default values, we should return back and give the
		// error to the user!
		_error().Err(ErrFailedToSetDefaultValuesForConfig).
			Str("error_message", _err.Error()).
			Msg("failed  to set default values for config")
		return nil, ErrFailedToSetDefaultValuesForConfig
	}

	// Set defaults for ReRunOnRecoverOptions
	_err = _struct.SetDefaultValues(&config.ReRunOnRecoverOptions)
	if _err != nil {
		// If there was an error setting the default values, we should return back and give the
		// error to the user!
		_error().Err(ErrFailedToSetDefaultValuesForConfig).
			Str("error_message", _err.Error()).
			Msg("failed  to set default values for ReRunOnRecoverOptions config")
		return nil, ErrFailedToSetDefaultValuesForConfig
	}

	// We haven't added context here... because we are not doing anything special
	// even if the program will terminate, the callback itself should have the context added
	// into it, it will simply remain with a status that is not running anymore!

	// Check if RunTime is not 0, RunTimes should be at least 1.
	// RunTimes tells how many times the OnRun should run!
	if config.RunTimes == 0 {
		_error().Err(ErrRunTimeIsZero).Msg("")
		return nil, ErrRunTimeIsZero
	}
	// -1 it's infinite! But if it's lower than -1 then it's not an ok value!
	if config.RunTimes < -1 {
		_error().Err(ErrRunTimesIsLowerThanOne).Msg("")
		return nil, ErrRunTimesIsLowerThanOne
	}

	// -1 it's infinite! But if it's lower than -1 then it's not an ok value!
	if config.ReRunOnRecoverOptions.MaxNrOfPanics < -1 {
		_error().Err(ErrReRunOnRecoverMaxNrOfPanicsIsLowerThanOne).Msg("")
		return nil, ErrReRunOnRecoverMaxNrOfPanicsIsLowerThanOne
	}

	// Get the root context if the one set by user it's nil!
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}

	// Create the instance
	instance := &GInstance{
		Logger: _logger,

		// if callback is running
		isCallbackRunning: _bool.New(),
		// notRunning is just a channel which will be used as notifier
		notRunning: make(chan bool),
		// isMonitoringRunning is default False!
		// It's turned on in the right moment...
		isMonitoringRunning: _bool.New(),
		// Save the config in the instance, later will be used!
		config: config,
		// Set the parent context, later will be used
		parentCtx: ctx,
		// is just defining if the process has being finished as planned
		isAllFinished: _bool.New(),

		// Locks
		isRunFuncRunning:  _bool.New(),
		isStopFuncRunning: _bool.New(),
	}

	// If config AutoRun it's true, then we should run the monitoring on create...
	if conv.ParseBool(config.AutoRun) {
		instance.Run()
	}

	return instance, nil
}
