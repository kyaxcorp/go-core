package config

import (
	"github.com/kyaxcorp/go-core/core/config"
	configModel "github.com/kyaxcorp/go-core/core/config/model"
	"log"
)

func LoadConfigs() bool {
	if !config.AutoLoad() {
		log.Fatal("Failed to load Main configuration file!")
		return false
	}

	// We receive the pointer of the object!
	paramsData, status := config.PreAutoLoadCustom(Data, &ConfigModel{})
	// We convert back to our Type!
	if !status {
		return false
	}
	Data = paramsData.(configModel.CustomConfigParams).Config.(*ConfigModel)

	return true
}
