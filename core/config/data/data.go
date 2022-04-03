package data

import (
	"github.com/kyaxcorp/go-core/core/config/model"
	"github.com/kyaxcorp/go-core/core/helpers/sync"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_bool"
	"github.com/spf13/viper"
)

var AutoLoaderLaunched sync.Mutex

// MainConfigDefaultsSetProcessing -> it shows if the default settings have being set to MainConfig
// It will be used in GetConfig function where we get the config...
var MainConfigDefaultsSetProcessing = _bool.NewVal(false)

// MainConfigDefaultsSetProcessed -> it's only for securing the processing function! When multiple goroutines access GetConfig
// And it's not yet processed!
var MainConfigDefaultsSetProcessed = _bool.NewVal(false)

//var MainConfig NonPtrObj
var MainConfig model.Model

// MainConfigJson -> here we store same MainConfig but in json format, why? for faster reading nested keys or
// for reading dynamically by using string paths
var MainConfigJson string

// MainConfigViper -> we store the instance of viper for the main config file!
var MainConfigViper *viper.Viper
