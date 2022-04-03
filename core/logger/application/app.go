package application

import (
	configEvents "github.com/KyaXTeam/go-core/v2/core/config/events"
	"github.com/KyaXTeam/go-core/v2/core/logger"
	"github.com/KyaXTeam/go-core/v2/core/logger/application/vars"
	loggerConfig "github.com/KyaXTeam/go-core/v2/core/logger/config"
	loggerPaths "github.com/KyaXTeam/go-core/v2/core/logger/paths"
)

// Define variables
var applicationLoggerConfig loggerConfig.Config

func CreateAppLogger() {
	applicationLoggerConfig, _ = loggerConfig.DefaultConfig(&loggerConfig.Config{
		IsEnabled:   "yes",
		Name:        "application",
		ModuleName:  "Application",
		Description: "saving all application logs...",
		Level:       0,
		DirLogPath:  loggerPaths.GetApplicationLogsPath(),
		// We set to yes, because this is the main Application Logger from which others will extend
		IsApplication: "yes",
	})
	// This is the Application Logger, it will save all logs
	vars.ApplicationLogger = logger.New(applicationLoggerConfig)
}

func RegisterAppLogger() {
	var _, _ = configEvents.OnLoaded(func() {
		CreateAppLogger()
	})
}
