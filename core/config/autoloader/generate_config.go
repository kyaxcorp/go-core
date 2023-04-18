package autoloader

import (
	"github.com/kyaxcorp/go-core/core/config/model"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/errors2"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/lock"
	"github.com/spf13/viper"
)

// Deprecated
// func GenerateConfig(c Config) error {
func GenerateConfig() error {
	configPath := GetConfigPath()
	if configPath == "" {
		return errors2.New(0, "config path is empty...")
	}

	configFilePath := GetConfigFilePath()
	if configFilePath == "" {
		return errors2.New(0, "config file path is empty...")
	}

	// We add the name as config path for uniqueness because multiple processes can read the same file!
	// the config can be modified by multiple processes at once if launched simultaneously!
	// This is why each process will do its work and after finishing it, the next process will do the same thing!
	// That will not degrade much in performance, but still will be a small slow down
	if isLockAcquired, lockErr := lock.FLock(configFilePath, true); !isLockAcquired || lockErr != nil {
		// Here we have some kind of error?!
		return errors2.New(0, "failed to lock config file -> ", lockErr.Error())
	}
	// Release the file lock on return
	defer lock.FRelease(configFilePath)

	v := viper.New()
	// NonPtrObj
	obj := &model.Model{}
	if _err := _struct.SetDefaultValues(obj); _err != nil {
		panic(_err)
	}
	// NonPtrObj of the custom config
	customObj := LoadedConfig.CustomConfigModel
	// Setting the default values
	if _err := _struct.SetDefaultValues(customObj); _err != nil {
		panic(_err)
	}

	// Setting in viper the main config
	v.Set("main", obj)
	// Setting in viper the custom config
	v.Set("custom", customObj)
	_err := v.SafeWriteConfigAs(configFilePath)
	if _err != nil {
		// log.Println("Failed to generate default config!")
		return errors2.New(0, "failed to generate default config -> "+_err.Error())
	}
	return nil
}

func GenerateConfigFromMemory() error {
	return SaveConfigFromMemory(LoadedConfig)
}
