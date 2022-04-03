package appLog

import (
	"github.com/KyaXTeam/go-core/v2/core/logger"
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	"github.com/rs/zerolog"
)

func getApp() *model.Logger {
	return logger.GetAppLogger()
}

//--------------------------\\

func Info() *zerolog.Event {
	return getApp().Info()
}

func Warn() *zerolog.Event {
	return getApp().Warn()
}

func Error() *zerolog.Event {
	return getApp().Error()
}

func Debug() *zerolog.Event {
	return getApp().Debug()
}

func Fatal() *zerolog.Event {
	return getApp().Fatal()
}

func Panic() *zerolog.Event {
	return getApp().Panic()
}

//--------------------------\\

func InfoF(functionName string) *zerolog.Event {
	return getApp().InfoF(functionName)
}

func WarnF(functionName string) *zerolog.Event {
	return getApp().WarnF(functionName)
}

func ErrorF(functionName string) *zerolog.Event {
	return getApp().ErrorF(functionName)
}

func DebugF(functionName string) *zerolog.Event {
	return getApp().DebugF(functionName)
}

func FatalF(functionName string) *zerolog.Event {
	return getApp().FatalF(functionName)
}

func PanicF(functionName string) *zerolog.Event {
	return getApp().PanicF(functionName)
}
