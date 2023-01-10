package autoloader

import (
	"github.com/caarlos0/env/v6"
	"github.com/fsnotify/fsnotify"
	"github.com/kyaxcorp/go-core/core/config"
	cfgData "github.com/kyaxcorp/go-core/core/config/data"
	"github.com/kyaxcorp/go-core/core/config/events"
	"github.com/kyaxcorp/go-core/core/config/model"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/_struct/defaults"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/err"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/lock"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/folder"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"github.com/kyaxcorp/go-core/core/helpers/json"
	"github.com/kyaxcorp/go-core/core/helpers/process"
	timezone "github.com/kyaxcorp/go-core/core/helpers/time"
	"github.com/kyaxcorp/go-core/core/logger/application"
	loggingConfig "github.com/kyaxcorp/go-core/core/logger/config"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"time"
)

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

func GenerateConfig(c Config) error {
	configPath := GetConfigPath()
	if configPath == "" {
		return err.New(0, "config path is empty...")
	}

	configFilePath := GetConfigFilePath()
	if configFilePath == "" {
		return err.New(0, "config file path is empty...")
	}

	// We add the name as config path for uniqueness because multiple processes can read the same file!
	// the config can be modified by multiple processes at once if launched simultaneously!
	// This is why each process will do its work and after finishing it, the next process will do the same thing!
	// That will not degrade much in performance, but still will be a small slow down
	if isLockAcquired, lockErr := lock.FLock(configFilePath, true); !isLockAcquired || lockErr != nil {
		// Here we have some kind of error?!
		return err.New(0, "failed to lock config file -> ", lockErr.Error())
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
	customObj := c.CustomConfigModel
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
		return err.New(0, "failed to generate default config -> "+_err.Error())
	}
	return nil
}

func StartAutoLoader(c Config) error {
	// Check if it's locked!
	// If yes then return

	// Set autoloader as called

	if !cfgData.AutoLoaderLaunched.TryLock() {
		return err.New(0, "failed to lock the loader")
	}

	// We should set the default values for the configuration!
	// Even when using file lock, it will be using some existing configuration values...
	// That's why before doing something, we should set the default values that are being defined
	config.ProcessConfig()

	// We set the default values for the custom config!
	if _err := _struct.SetDefaultValues(c.CustomConfig); _err != nil {
		return err.New(0, "failed to set default values for CustomConfig -> ", _err)
		//panic(_err)
	}

	// Check if autoloader launched!
	// If launched already, don't relaunch again!
	// We should add a watcher with viper to monitor for config changes!
	// The autoloader should run in a separate process!

	// Create a global lock for this configuration path! the lock should be based on the path?!
	// If other configuration path is provided, then, the lock should be for the new configuration path!
	// The other configuration path should be read also here... from the arguments provided in  the app!?
	configPath := GetConfigPath()
	if configPath == "" {
		return err.New(0, "config path is empty...")
	}

	configFilePath := GetConfigFilePath()
	if configFilePath == "" {
		return err.New(0, "config file path is empty...")
	}
	// We add the name as config path for uniqueness because multiple processes can read the same file!
	// the config can be modified by multiple processes at once if launched simultaneously!
	// This is why each process will do its work and after finishing it, the next process will do the same thing!
	// That will not degrade much in performance, but still will be a small slow down
	if isLockAcquired, lockErr := lock.FLock(configFilePath, true); !isLockAcquired || lockErr != nil {
		// Here we have some kind of error?!
		return err.New(0, "failed to lock config file -> ", lockErr.Error())
	}
	// Release the file lock on return
	defer lock.FRelease(configFilePath)

	// Check if the configuration exists...
	if !IsConfigExists() {
		// if config doesn't exist... what should we do?!
		// it means we don't load anything?
	} else {
		// Do a backup of the current config file before launching anything else!
		// This ensures that we will have a copy of the previous file before doing any automatic changes
		// Check if the config exists

		backupFullPath := GetBackupFullPath()
		if backupFullPath == "" {
			return err.New(0, "config Backup Path Folder is empty...")
		}

		// Keep a max nr of backups in a folder!
		// For that we will use a special function which will clean the map by time order
		// We will store backups for 30 days period
		folder.ScanAndClean(backupFullPath, 30, nil, nil, nil)

		// Create additional map folder separated by day month and year
		backupFullPath += filesystem.DirSeparator() + time.Now().Format("02.01.2006")
		// We need the backup folder path with time...
		backupFolderPath := backupFullPath

		// Create the folder...
		if !folder.Exists(backupFullPath) {
			folder.MkDir(backupFullPath)
		}

		// The backup will be created based on the comparison with the previous existing file from the existing folder
		// If the current config and the prev file from backup are different, the app will create the backup!

		currentConfigChecksum, _err := hash.FileSha256(configFilePath)
		if _err != nil {
			return err.New(0, "failed to generate checksum for "+configFilePath+" -> "+_err.Error())
		}
		createBackup := false

		// Find the last backup file from that directory
		backups, _err := ioutil.ReadDir(backupFolderPath)
		if _err != nil {
			// log.Fatal(_err)
			return err.New(0, "failed to read the backups folder -> "+_err.Error())
		}

		//log.Println(backupFolderPath)
		//log.Println(backups)

		if len(backups) > 0 {
			// Let's get the last file
			lastBackupFile := backups[len(backups)-1]
			// Generate the full path of the backup
			lastBackupFullPath := backupFolderPath + filesystem.DirSeparator() + lastBackupFile.Name()
			// Generate the checksum of the backup
			backupConfigChecksum, _err := hash.FileSha256(lastBackupFullPath)
			if _err != nil {
				return err.New(0, "failed to generate checksum for "+backupConfigChecksum+" -> "+_err.Error())
			}

			// Compare the 2 files and if are different then create the backup!
			if currentConfigChecksum != backupConfigChecksum {
				createBackup = true
			}
		} else {
			// There are no backups there... so we will create one
			createBackup = true
		}

		if createBackup {
			//backupFullPath += filesystem.DirSeparator() + conv.Int64ToStr(time.Now().UnixNano()) + "_" + conv.IntToStr(process.GetCurrentProcessPID()) + ".backup"
			// Month day year hour minute second nanosecond pid
			backupFullPath += filesystem.DirSeparator() +
				time.Now().Format("02.01.2006_15.04.05.999999999") +
				"_" + conv.IntToStr(process.GetCurrentProcessPID()) + ".backup.yaml"

			_, _err := file.Copy(configFilePath, backupFullPath)
			if _err != nil {
				return err.New(0, "failed to create a backup of the current config file! -> "+_err.Error())
			}
		}
	}

	// Now we should launch the viper instance which will monitor for changes and load the config into it
	// Load the current config in memory!
	cfgData.MainConfigViper = viper.New()

	cfgData.MainConfigViper.AddConfigPath(configPath)
	cfgData.MainConfigViper.SetConfigName(GetConfigFileName())
	cfgData.MainConfigViper.SetConfigType(GetConfigFileType())
	// c.SetDefault("main", NonPtrObj{})
	// c.AutomaticEnv()

	// We should indicate him the structure! so he could compare and unmarshal properly...
	// c.SetDefault("main", NonPtrObj{})
	//c.Set("main", defaultConfig)

	_err := cfgData.MainConfigViper.ReadInConfig()
	if _err != nil {
		return err.New(0, "failed to read config in viper -> "+_err.Error())
	}

	//log.Println(c.Get("main"))

	// ==================== DEFAULTS ====================\\
	// This is the Standard Config Structure
	obj := &model.Model{}
	if _err := defaults.Set(obj); _err != nil {
		panic(_err)
	}
	// I'm not sure if i am doing right over here!... but it works... (13.05.2021)
	cfgData.MainConfig = *obj

	// This is the Custom Config Structure
	objCustom := c.CustomConfigModel
	if _err := defaults.Set(objCustom); _err != nil {
		panic(_err)
	}
	// ==================== DEFAULTS ====================\\

	//

	// =================== VIPER SET ======================\\
	// c.Sub("main")
	_err = cfgData.MainConfigViper.UnmarshalKey("main", &cfgData.MainConfig)
	if _err != nil {
		return err.New(0, "failed to decode 'main' key from config -> "+_err.Error())
	}

	_err = cfgData.MainConfigViper.UnmarshalKey("custom", c.CustomConfig)
	if _err != nil {
		return err.New(0, "failed to decode 'custom' key from config -> "+_err.Error())
	}
	// =================== VIPER SET ======================\\

	//

	// We should save the configuration only if has being changed!
	// But this can be made by saving in other temporary location, and after that comparing the contents of the both files!

	// Save again the config with the newly added/removed keys based on the app structure!
	if _err := SaveConfigFromMemory(c); _err != nil {
		return err.New(0, "failed to save config from memory -> "+_err.Error())
	}

	// ===================== ENV ========================\\
	if _err = env.Parse(&cfgData.MainConfig); _err != nil {
		return err.New(0, "failed to set env variables for MainConfig -> "+_err.Error())
	}

	if _err = env.Parse(c.CustomConfig); _err != nil {
		return err.New(0, "failed to set env variables for CustomConfig -> "+_err.Error())
	}
	// ===================== ENV ========================\\

	// Launch config watcher... if something changes, we will notify the others
	cfgData.MainConfigViper.WatchConfig()
	cfgData.MainConfigViper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		// If a change occured, we should check if the config is ok!
		// we should set the new config to the memory var

		// Check if viper has no errors reading the config file...!
		_err = cfgData.MainConfigViper.UnmarshalKey("main", &cfgData.MainConfig)
		if _err != nil {
			log.Println(_err)
			// return false
		}

		_err = cfgData.MainConfigViper.UnmarshalKey("custom", c.CustomConfig)
		if _err != nil {
			log.Println(_err)
			// return false
		}

		// ===================== ENV ========================\\
		if _err = env.Parse(&cfgData.MainConfig); _err != nil {
			log.Println(_err)
			//return err.New(0, "failed to set env variables for MainConfig -> "+_err.Error())
		}

		if _err = env.Parse(c.CustomConfig); _err != nil {
			log.Println(_err)
			//return err.New(0, "failed to set env variables for CustomConfig -> "+_err.Error())
		}
		// ===================== ENV ========================\\

		// log.Println(cfgData.MainConfig)

		// TODO: we can create other triggers from here...
		// But it's not easy to reload configurations for multiple services...
		// This is why each service should monitor a specific part of the configuration!
	})

	// save in memory same config but as JSON
	cfgData.MainConfigJson, _err = json.Encode(cfgData.MainConfig)
	if _err != nil {
		log.Println(_err)
		panic("failed to convert MainConfig to json and save it...")
	}

	// Everything is ok...
	// Let's call the events...
	application.CreateAppLogger(application.MainLogOptions{
		Level: cfgData.MainConfig.Logging.AppLogLevel,
	})

	/*// Register the broker client service
	brokerClientService.RegisterBrokerService()
	// Run the registered services
	register_service.RunRegisteredServices()*/

	events.CallOnLoaded()
	go events.CallOnLoadedAsync()

	// Let's call other additional functions
	if !conv.ParseBool(config.GetConfig().Application.DisableTimezoneOverride) {
		timezone.OverrideLocalTimezone(config.GetConfig().Application.TimeZone)
	}

	cwd := config.GetConfig().Application.CurrentWorkingDirectory
	if cwd == "" {
		if conv.ParseBool(config.GetConfig().Application.IfEmptyCWDSetToExecPath) {
			_err := os.Chdir(path.Root())
			if _err != nil {
				panic(_err)
			}
		} else {
			// it will remain as it is!
		}
	} else {
		// set to this path!
		// but first, we should check if this path really exists!
		if folder.Exists(cwd) {
			// let's set it!
			_err := os.Chdir(cwd)
			if _err != nil {
				panic(_err)
			}
		}
	}

	return nil
}

func StopAutoLoader() {
	// Stop Viper?!
}
