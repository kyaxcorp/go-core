package autoloader

import loggingConfig "github.com/kyaxcorp/go-core/core/logger/config"

var LoadedConfig Config

type Config struct {
	// This is where the Custom Config data is saved... it's a pointer!
	CustomConfig interface{}
	// This is the structure of the Custom Config, from which the yaml is generated
	CustomConfigModel interface{}
	// Here we define Additional logging channels if required... for example we know that our app
	// has some functionality which requires logging channeling, so everything is isolated and understood
	// so, we also need to generate a config file which already has that channels without a user intervention
	// or having a ready template file! So in this case we dictate from here additional channels
	AdditionalLoggingChannels map[string]loggingConfig.Config
}
