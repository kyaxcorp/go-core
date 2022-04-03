package certs

import (
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"log"
)

func GetCertsFullPath() string {
	var err error = nil
	certsPath := config.GetConfig().Application.CertsPath
	certsPath, err = fsPath.GenRealPath(certsPath, true)

	if err != nil {
		log.Println(err)
	}

	if !filesystem.Exists(certsPath) {
		filesystem.MkDir(certsPath)
	}

	return certsPath
}

func GetCertsFullPathByScope(scope string) string {
	certsPath := GetCertsFullPath()
	if certsPath == "" {
		return ""
	}

	certsPath = certsPath + filesystem.DirSeparator() + scope
	if !filesystem.Exists(certsPath) {
		filesystem.MkDir(certsPath)
	}
	return certsPath
}
