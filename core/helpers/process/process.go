package process

import (
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/folder"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"os"
	"path/filepath"
	"strconv"
)

func getPIDsFolderPath() (string, error) {
	var _err error = nil
	pidsPath := config.GetConfig().Application.PIDsPath
	pidsPath, _err = fsPath.GenRealPath(pidsPath, true)

	if _err != nil {
		//log.Println(err)
		return "", _err
	}

	if !folder.Exists(pidsPath) {
		folder.MkDir(pidsPath)
	}

	return pidsPath, nil
}

func getPidPath(name string) (string, error) {
	pidsFolder, _err := getPIDsFolderPath()
	if pidsFolder == "" || _err != nil {
		return "", _err
	}
	return pidsFolder + filesystem.DirSeparator() + hash.Sha256(name) + ".pid", nil
}

func PIDDestroy(name string) (bool, error) {
	pidFilePath, _err := getPidPath(name)
	if pidFilePath == "" || _err != nil {
		return false, _err
	}
	return file.Unlink(pidFilePath)
}

func GeneratePID(name string) (bool, error) {
	pidFilePath, _err := getPidPath(name)
	if pidFilePath == "" || _err != nil {
		return false, _err
	}

	processPid := os.Getpid()
	_err = file.PutContents(pidFilePath, strconv.FormatInt(int64(processPid), 10))
	if _err != nil {
		return false, _err
	}
	return true, nil
}

func GetCurrentProcessPID() int {
	return os.Getpid()
}

func GetExistingPID(name string) (string, error) {
	pidFilePath, _err := getPidPath(name)
	if pidFilePath == "" || _err != nil {
		return "", _err
	}

	exists, _err := file.ExistsErr(pidFilePath)
	if _err != nil || !exists {
		return "", _err
	}

	//log.Println(pid_file_path)

	pid, _err := file.GetContents(pidFilePath)
	if _err != nil {
		//log.Println(_ere)
		return "", _err
	}
	return pid, nil
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
