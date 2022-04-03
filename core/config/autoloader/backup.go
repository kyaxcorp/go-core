package autoloader

import (
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
)

func GetBackupFullPath() string {
	ConfigsBackupPath := config.GetConfig().Application.ConfigsBackupPath
	backupPath, _err := fsPath.GenRealPath(ConfigsBackupPath, true)
	if _err != nil {
		return ""
	}

	// Create the backup folder
	if !filesystem.Exists(backupPath) {
		filesystem.MkDir(backupPath)
	}

	return backupPath
}
