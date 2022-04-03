package config

import (
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/rs/zerolog"
	"io"
)

type Logger interface {
	GetLogger() *zerolog.Logger
}

type Config struct {
	// IsEnabled -> enable/disable the logging (this is generally)
	// By default will be enabled, and the default level will be 4 -> which is Warn!
	IsEnabled string `yaml:"is_enabled" mapstructure:"is_enabled" default:"yes"`
	// Name -> this is the identifier of the instance
	Name string
	// ModuleName -> identifies the module in logs...Each component should set its own module name
	ModuleName string `yaml:"module_name" mapstructure:"module_name" default:""`

	// Description -> something about it
	Description string

	EncodeLogsAsJson string `yaml:"encode_logs_as_json" mapstructure:"encode_logs_as_json" default:"yes"`

	// CliIsEnabled -> If console logging is enabled
	ConsoleIsEnabled string `yaml:"console_is_enabled" mapstructure:"console_is_enabled" default:"yes"`

	ConsoleTimeFormat string `yaml:"console_time_format" mapstructure:"console_time_format" default:"002 15:04:05.000"`

	// Level -> the level which should be for the output, it's for Console and for the file log!
	// For more specific levels:
	/*
		1: zerolog.TraceLevel,
		2: zerolog.DebugLevel,
		3: zerolog.InfoLevel,
		4: zerolog.WarnLevel,
		5: zerolog.ErrorLevel,
		6: zerolog.FatalLevel,
		7: zerolog.PanicLevel,
	*/
	// The default one is INFO!
	// We will not set the default value... so it will be by default 0 -> meaning DEBUG, we don't set, because when it's 0
	// it's counted as initial value!
	// We will set the default value to 2 -> Debug Level
	Level int `yaml:"level" mapstructure:"level" default:"4"`

	// FileName -> is optional, and can be set only if the user wants to override the default name, or to override the full
	// path of the writing file... it can be user with DirLogPath
	FileName string `yaml:"file_name" mapstructure:"file_name" default:"-"`

	// FileIsEnabled -> if file logging is enabled
	FileIsEnabled string `yaml:"file_is_enabled" mapstructure:"file_is_enabled" default:"yes"`
	// LogPath -> where the logs should be saved
	DirLogPath string `yaml:"dir_log_path" mapstructure:"dir_log_path" default:""`
	// RotateMaxFiles -> how many files should be in the same folder, it will auto delete the old ones
	FileRotateMaxFiles int `yaml:"file_rotate_max_files" mapstructure:"file_rotate_max_files" default:"50"`
	// LogMaxSize -> the maximum size of a log file, if it's getting big, it will be created a new one -> default 5 MB
	FileLogMaxSize int `yaml:"file_log_max_size_mb" mapstructure:"file_log_max_size_mb" default:"5"`
	// MaxAge the max age in days to keep a logfile
	FileMaxAge int `yaml:"file_max_age_days" mapstructure:"file_max_age_days" default:"30"`
	// Is the main logger from which everyone will extend!? We don't need to export it!
	IsApplication string `yaml:"-" mapstructure:"-" default:"-"`

	// This is the parent... so the child can write to parent log
	//ParentLogger Config
	ParentWriter io.Writer `yaml:"-" mapstructure:"-" default:"-"`
	// WriteToParent -> also write to parent log file
	WriteToParent string `yaml:"write_to_parent" mapstructure:"write_to_parent" default:"yes"`

	// This is an interface to get to logger object directly
	Logger Logger `yaml:"-" mapstructure:"-" default:"-"`
}

// DefaultConfig -> it will return the default config with default values
func DefaultConfig(configObj *Config) (Config, error) {
	if configObj == nil {
		configObj = &Config{}
	}
	// Set the default values for the object!
	_err := _struct.SetDefaultValues(configObj)
	return *configObj, _err
}
