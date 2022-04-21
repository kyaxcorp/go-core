package _runtime

import (
	"runtime"
)

const (
	Windows = "windows"
	Linux   = "linux"
)

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func IsLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

func OS() string {
	return runtime.GOOS
}
