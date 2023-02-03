package autoloader

import (
	"github.com/kyaxcorp/go-core/core/config/model"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"github.com/kyaxcorp/go-core/core/helpers/process/name"
	"os"
	"path/filepath"
	"strings"
)

var cachedAppEnvConfigPathParamName string

func GetConfigPath() string {
	// Check if there is an argument having the -c=path or --config=path
	// Or check if there is an ENVIRONMENT PARAM
	// the param should based on the app name
	appName := GetCleanAppFileName()
	if cachedAppEnvConfigPathParamName == "" {
		cachedAppEnvConfigPathParamName = strings.ReplaceAll("-", "_", strings.ToUpper(appName)) + "_" + "CONFIG_PATH"
	}
	// Check if exists
	envConfigPath := os.Getenv(cachedAppEnvConfigPathParamName)
	if envConfigPath != "" {
		return envConfigPath
	}

	if globalConfigPath != "" {
		return globalConfigPath
	}

	// Get the config full path
	configFilePath := GetConfigFilePath()
	// Get only the base dir without file
	return filepath.Dir(configFilePath) + filepath.FromSlash("/")
}

func GetConfigFilePath() string {
	//path := GetConfigPath()
	// TODO: we should add here from the arguments introduced for the process!
	// There are 3 variants:
	// The default one is from the root directory
	// From the argument list from the process
	// From the OS Default config path: /etc/.... or Windows somewhere...

	path := fsPath.Root()

	if path != "" {
		path += GetConfigFullFileName()
	}
	return path
}

// GetConfigFileName
// We should have the same config name if the app name is not changed (doesn't matter if it's on windows or Linux)
// we should remove the file extension!
func GetConfigFileName() string {
	return model.ConfigFileName + "_" + hash.MD5(GetCleanAppFileName())
}

func GetCleanAppFileName() string {
	return name.GetCurrentProcessCleanExecName()
}

func GetConfigFileType() string {
	return model.ConfigFileType
}

func GetConfigFullFileName() string {
	return GetConfigFileName() + "." + GetConfigFileType()
}

func GetTmpConfigFileName() string {
	return model.ConfigTmpFileName + "_" + hash.MD5(GetCleanAppFileName()) + "." + model.ConfigFileType
}

// GetConfigTmpFilePath -> this is the temporary file path of the config... it's only for comparation purpose
func GetConfigTmpFilePath() string {
	path := GetConfigPath()
	if path != "" {
		path += GetTmpConfigFileName()
	}
	return path
}

func GetCustomConfigFilePath() string {
	// TODO: try reading from environment values or input arguments the path of the config!
	path := GetConfigPath()
	if path != "" {
		path = path + model.CustomConfigFileName + "." + model.CustomConfigFileType
	}
	return path
}

func IsConfigExists() bool {
	path := GetConfigFilePath()
	if path == "" {
		return false
	}
	return file.Exists(path)
}

func IsCustomConfigExists() bool {
	path := GetCustomConfigFilePath()
	if path == "" {
		return false
	}
	return file.Exists(path)
}

func IsConfigValid() bool {
	// TODO: load through viper and check if loaded!
	return true
}
