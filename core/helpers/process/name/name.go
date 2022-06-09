package name

import (
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"os"
	"path/filepath"
	"strings"
)

var cachedCleanAppName string

// GetCurrentProcessExecName -> it will return executable name of the current process
func GetCurrentProcessExecName() string {
	return filepath.Base(os.Args[0])
}

func GetCurrentProcessCleanExecName() string {
	if cachedCleanAppName != "" {
		return cachedCleanAppName
	}

	appFullName := GetCurrentProcessExecName()
	appName := ""
	extensionSep := "."
	if strings.Contains(appFullName, extensionSep) {
		appNameSplit := strings.Split(appFullName, extensionSep)
		appName = strings.Join(appNameSplit[:len(appNameSplit)-1], extensionSep)
	} else {
		appName = appFullName
	}
	cachedCleanAppName = appName
	return appName
}

func GetCurrentProcessCleanMD5ExecName() string {
	return hash.MD5(GetCurrentProcessCleanExecName())
}
