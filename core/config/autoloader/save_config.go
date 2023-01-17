package autoloader

import (
	"github.com/kyaxcorp/go-core/core/helpers/file"
	// cassandraConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/cassandra/config"

	cfgData "github.com/kyaxcorp/go-core/core/config/data"
	"github.com/kyaxcorp/go-core/core/helpers/err"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	//brokerConfig "github.com/kyaxcorp/go-core/core/services/broker/config"
	"github.com/spf13/viper"
)

func SaveConfigFromMemory(cfg Config) error {
	c := viper.New()

	var _err error

	// Setting the main config
	c.Set("main", cfgData.MainConfig)
	// Setting the custom config
	c.Set("custom", cfg.CustomConfig)

	// TODO: save config only by comparing if it's different!
	// If it's diff, then overwrite it!
	configPath := GetConfigFilePath()
	if configPath == "" {
		return err.New(0, "config path is empty")
	}

	configTmpPath := GetConfigTmpFilePath()
	if configTmpPath == "" {
		return err.New(0, "config tmp path is empty")
	}
	// Save the temporary config file
	_err = c.WriteConfigAs(configTmpPath)
	if _err != nil {
		// log.Println("Failed to generate config!")
		return _err
	}

	// Compare the 2 configs
	tmpConfigHash, _err := hash.FileSha256(configTmpPath)
	// Delete the tmp config
	file.Delete(configTmpPath)
	if _err != nil {
		return _err
	}

	realConfigHash, _err := hash.FileSha256(configPath)
	if _err != nil {
		return _err
	}
	// log.Println(realConfigHash, tmpConfigHash)

	// Compare the 2 configs
	if tmpConfigHash == realConfigHash {
		// It's the same configuration!
		// log.Println("Same config!!! skipping save...")
		return nil
	}

	// Save the real config file
	_err = c.WriteConfigAs(configPath)
	if _err != nil {
		// log.Println("Failed to generate config!")
		return _err
	}
	return nil
}
