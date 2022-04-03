package autoloader

var globalConfigPath string

func OverrideGlobalConfigPath(newConfigPath string) {
	globalConfigPath = newConfigPath
}
