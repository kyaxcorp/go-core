package application

import (
	configEvents "github.com/kyaxcorp/go-core/core/config/events"
	"github.com/kyaxcorp/go-core/core/logger"
	"github.com/kyaxcorp/go-core/core/logger/application/vars"
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
	loggerPaths "github.com/kyaxcorp/go-core/core/logger/paths"
)

// Define variables
var applicationLoggerConfig loggerConfig.Config

type MainLogOptions struct {
	Level int
}

func CreateAppLogger(o MainLogOptions) {
	applicationLoggerConfig, _ = loggerConfig.DefaultConfig(&loggerConfig.Config{
		IsEnabled:   "yes",
		Name:        "application",
		ModuleName:  "Application",
		Description: "saving all application logs...",
		Level:       o.Level,
		DirLogPath:  loggerPaths.GetApplicationLogsPath(),
		// We set to yes, because this is the main Application Logger from which others will extend
		IsApplication: "yes",
	})
	// This is the Application Logger, it will save all logs
	vars.ApplicationLogger = logger.New(applicationLoggerConfig)
}

func RegisterAppLogger() {
	var _, _ = configEvents.OnLoaded(func() {
		CreateAppLogger(MainLogOptions{})
	})
}
