package autoloader

import (
	"github.com/KyaXTeam/go-core/v2/core/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/filesystem"
	fsPath "github.com/KyaXTeam/go-core/v2/core/helpers/filesystem/path"
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
