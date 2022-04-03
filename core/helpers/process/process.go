package process

import (
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func getPIDsFolderPath() string {
	var err error = nil
	pidsPath := config.GetConfig().Application.PIDsPath
	pidsPath, err = fsPath.GenRealPath(pidsPath, true)

	if err != nil {
		log.Println(err)
	}

	if !filesystem.Exists(pidsPath) {
		filesystem.MkDir(pidsPath)
	}

	return pidsPath
}

func getPidPath(name string) string {
	return getPIDsFolderPath() + filesystem.DirSeparator() + hash.Sha256(name) + ".pid"
}

func PIDDestroy(name string) bool {
	pidFilePath := getPidPath(name)
	if pidFilePath == "" {
		return false
	}

	return filesystem.Unlink(pidFilePath)
}

func GeneratePID(name string) bool {
	pidFilePath := getPidPath(name)
	if pidFilePath == "" {
		return false
	}

	processPid := os.Getpid()
	filesystem.FilePutContents(pidFilePath, strconv.FormatInt(int64(processPid), 10))
	return true
}

func GetCurrentProcessPID() int {
	return os.Getpid()
}

func GetExistingPID(name string) string {
	pidFilePath := getPidPath(name)
	if pidFilePath == "" {
		return ""
	}

	if !filesystem.FileExists(pidFilePath) {
		return ""
	}

	//log.Println(pid_file_path)

	pid, err := filesystem.FileGetContents(pidFilePath)
	if err != nil {
		log.Println(err)
	}
	return pid
}

func GetCurrentProcessFolder() (string, error) {
	dir, _err := filepath.Abs(filepath.Dir(os.Args[0]))
	if _err != nil {
		return "", _err
	}
	return dir, nil
}

// GetCurrentProcessExecName -> it will return executable name of the current process
func GetCurrentProcessExecName() string {
	return filepath.Base(os.Args[0])
}
