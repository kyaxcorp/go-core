package model

import (
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/rs/zerolog"
)

const funcName = "function_name"

func (l *Logger) GetLogger() *zerolog.Logger {
	return l.Logger
}

// Debug -> 0
func (l *Logger) Debug() *zerolog.Event {
	return l.Logger.Debug()
}

// DebugF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) DebugF(functionName string) *zerolog.Event {
	return l.Debug().Str(funcName, functionName)
}

// Info -> 1
func (l *Logger) Info() *zerolog.Event {
	return l.Logger.Info()
}

// InfoF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) InfoF(functionName string) *zerolog.Event {
	return l.Info().Str(funcName, functionName)
}

// Warn -> 2
func (l *Logger) Warn() *zerolog.Event {
	return l.Logger.Warn()
}

// WarnF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) WarnF(functionName string) *zerolog.Event {
	return l.Warn().Str(funcName, functionName)
}

// Error -> 3
func (l *Logger) Error() *zerolog.Event {
	return l.Logger.Error()
}

// ErrorF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) ErrorF(functionName string) *zerolog.Event {
	return l.Error().Str(funcName, functionName)
}

// Fatal -> 4
func (l *Logger) Fatal() *zerolog.Event {
	return l.Logger.Fatal()
}

func (l *Logger) FatalF(functionName string) *zerolog.Event {
	return l.Fatal().Str(funcName, functionName)
}

func (l *Logger) PanicF(functionName string) *zerolog.Event {
	return l.Panic().Str(funcName, functionName)
}

// Panic -> 5
func (l *Logger) Panic() *zerolog.Event {
	return l.Logger.Panic()
}

// InfoEvent -> when executing some events
func (l *Logger) InfoEvent(
	eventType string,
	eventName string,
	beforeMsg func(event *zerolog.Event),
) {
	if eventType == "" {
		eventType = "no_event_type"
	}
	if eventName == "" {
		eventName = "no_event_name"
	}
	info := l.Info()
	if function.IsCallable(beforeMsg) {
		beforeMsg(info)
	}
	info.Str("event_type", eventType).
		Str("event_name", eventName).
		Msg("event.execution")
}

// InfoEventF -> with function name
func (l *Logger) InfoEventF(
	eventType string,
	eventName string,
	functionName string,
) *zerolog.Event {
	return l.InfoEventCustom(eventType, eventName).Str(funcName, functionName)
}

// InfoEventCustom -> when executing some events
func (l *Logger) InfoEventCustom(
	eventType string,
	eventName string,
) *zerolog.Event {
	if eventType == "" {
		eventType = "no_event_type"
	}
	if eventName == "" {
		eventName = "no_event_name"
	}
	info := l.Info()
	return info.
		Str("event_type", eventType).
		Str("event_name", eventName).
		Str("action", "event.execution")
}
