package custom_http

import (
	"github.com/kyaxcorp/go-core/core/helpers/_runtime"
	"os"
)
import "os/exec"

func Start() bool {
	/*
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
	*/

	switch _runtime.OS() {
	case _runtime.Windows:
		// TODO: create a simple function copy embed file!

		//CustomHttpStorage.Open()
		fsData, _err := CustomHttpStorage.ReadFile("storage/win/generic/httpserver.exe")
		if _err != nil {
			// TODO: handle this error!?...
			return false
		}
		// TODO: get app's directory path
		filePath := "aaaa.exe"
		fs, _err := os.Create(filePath) // TODO: get app's name
		if _err != nil {
			// TODO: handle this error!?...
			return false
		}
		fileSize := len(fsData)
		sizeWritten, _err := fs.Write(fsData)
		if _err != nil {
			// TODO: handle this error!?...
			return false
		}
		if fileSize != sizeWritten {
			// TODO: handle this error!?...
			return false
		}

		// TODO: check what arguments should be called
		command := exec.Command(filePath)
		// TODO: handle stdin & stdout -> send it to null!
		_err := command.Start()
		if _err != nil {
			// TODO: handle this error!?...
			return false
		}
		command.Process.Pid

		// Check if it's running, get the pid
	case _runtime.Linux:

	default:
		return false
	}

	return true
}
