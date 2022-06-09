package stage

import "github.com/kyaxcorp/go-core/core/config"

func Get() string {
	return config.GetConfig().Application.Stage
}

func IsDev() bool {
	if Get() == "dev" {
		return true
	}
	return false
}

func IsProd() bool {
	if Get() == "prod" {
		return true
	}
	return false
}

func IsTest() bool {
	if Get() == "test" {
		return true
	}
	return false
}
