package application

import (
	configEvents "github.com/kyaxcorp/go-core/core/config/events"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/logger"
	"github.com/kyaxcorp/go-core/core/logger/application/vars"
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
	loggerPaths "github.com/kyaxcorp/go-core/core/logger/paths"
	"os"
)

// Define variables
var applicationLoggerConfig loggerConfig.Config
var coreLoggerConfig loggerConfig.Config

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

func CreateCoreLogger() bool {
	logLevel := os.Getenv("GO_CORE_LOG_LEVEL")
	var lvl int
	if logLevel == "" {
		lvl = 4
	} else {
		lvl = conv.StrToInt(logLevel)
	}

	coreLoggerConfig, _ = loggerConfig.DefaultConfig(&loggerConfig.Config{
		IsEnabled:   "yes",
		Name:        "core",
		ModuleName:  "Core",
		Description: "saving all core logs...",
		Level:       lvl, // take from the environment

		FileIsEnabled:    "no",
		ConsoleIsEnabled: "yes",
	})
	// This is the Application Logger, it will save all logs
	vars.CoreLogger = logger.New(coreLoggerConfig)
	return true
}

var _ = CreateCoreLogger()
