package config

import (
	"github.com/kyaxcorp/go-core/core/config/model"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/tidwall/gjson"
)
import "github.com/kyaxcorp/go-core/core/config/data"

func GetConfig() model.Model {
	ProcessConfig()
	return data.MainConfig
}

func GetConfigByKey(keyPath string) gjson.Result {
	ProcessConfig()
	return gjson.Get(data.MainConfigJson, keyPath)
}

func ProcessConfig() {
	// Check if Config is being processed right now! If multiple goroutines access it, they should wait
	// Until the processing has being finished!
	if !data.MainConfigDefaultsSetProcessing.IfFalseSetTrue() {
		// Start processing...

		if _err := _struct.SetDefaultValues(&data.MainConfig); _err != nil {
			panic(_err)
		}
		// Set to true!
		data.MainConfigDefaultsSetProcessed.True()
		// Set that is not processing anymore!
		data.MainConfigDefaultsSetProcessing.False()
	} else {
		// Everyone will wait until is processed!
		data.MainConfigDefaultsSetProcessed.WaitUntilTrue()
	}
}
