package autoloader

import (
	"github.com/kyaxcorp/go-core/core/config"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/folder"
)

func GetBackupFullPath() string {
	ConfigsBackupPath := config.GetConfig().Application.ConfigsBackupPath
	backupPath, _err := fsPath.GenRealPath(ConfigsBackupPath, true)
	if _err != nil {
		return ""
	}

	// Create the backup folder
	if !folder.Exists(backupPath) {
		folder.MkDir(backupPath)
	}

	return backupPath
}
