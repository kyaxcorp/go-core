package channel

import (
	mainConfig "github.com/KyaXTeam/go-core/v2/core/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err"
	"github.com/KyaXTeam/go-core/v2/core/logger"
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	loggerPaths "github.com/KyaXTeam/go-core/v2/core/logger/paths"
	"github.com/rs/zerolog"
)

type Config struct {
	ChannelName              string
	ReturnDefaultIfNotExists bool
}

// GetDefaultChannel -> gets the default logger based on the current configuration
// Finds the default one
// checks if the object is created in memory
// if yes then returns it, if not it creates it based on the configuration
func GetDefaultChannel() (*model.Logger, error) {
	cfg := mainConfig.GetConfig()
	if cfg.Logging.DefaultChannel == "" {
		msg := "logger default channel name is empty"
		l().Warn().Msg(msg)
		return nil, err.New(0, msg)
	}
	return GetChannel(Config{
		ChannelName: cfg.Logging.DefaultChannel,
	})
}

// GetChannel -> Get channel based on instance name (the key from the config)
// Check if there is a channel like this...
// If there is not, then return an error...
// If there is, return the logger based on the config found
func GetChannel(c Config) (*model.Logger, error) {
	cfg := mainConfig.GetConfig()
	// Check if exists
	if _, ok := cfg.Logging.Channels[c.ChannelName]; !ok {
		// Doesn't exist
		if c.ReturnDefaultIfNotExists {
			// Return the default one!
			return GetDefaultChannel()
		}

		msg := "logger channel doesn't exist"
		l().Warn().Str("logger_channel", c.ChannelName).Msg(msg)
		return nil, err.New(0, msg)
	}

	// Exists
	instanceConfig := cfg.Logging.Channels[c.ChannelName]

	// Setting default values if some of them are missing in the config!
	// Setting instance name by key
	if instanceConfig.Name == "" {
		instanceConfig.Name = c.ChannelName
	}
	// If DirLogPath is not defined, it will set the default folder!
	if instanceConfig.DirLogPath == "" {
		instanceConfig.DirLogPath = loggerPaths.GetLogsPathForChannels("websocket/" + instanceConfig.Name)
	}

	// Set Module Name
	if instanceConfig.ModuleName == "" {
		instanceConfig.ModuleName = instanceConfig.Name
	}

	return logger.New(instanceConfig), nil
}

// log -> This is for local use only
func l() *zerolog.Logger {
	return logger.GetAppLogger().Logger
}
