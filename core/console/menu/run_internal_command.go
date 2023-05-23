package menu

import (
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"os"
	"os/exec"
	"path/filepath"
)

type InternalCommandOptions struct {
	Args                 []string
	Release              bool
	RunDaemonFromExecDir bool
}

// func (m *Menu) RunInternalCommand(arg ...string) (*exec.Cmd, error) {
func (m *Menu) RunInternalCommand(options InternalCommandOptions) (*exec.Cmd, error) {
	var _err error
	var currentApp string
	var execPath string
	execPath, _err = os.Executable()
	currentApp = execPath
	// let's change the working directory!
	if options.RunDaemonFromExecDir {
		appDirPath := filepath.Dir(currentApp)
		_err = os.Chdir(appDirPath)
		if _err != nil {
			return nil, _err
		}
		currentApp = "./" + filepath.Base(execPath)
	}

	var _cmd *exec.Cmd
	if options.Release {
		_cmd = exec.Command(currentApp, options.Args...)
	} else {
		_cmd = exec.CommandContext(_context.GetDefaultContext(), currentApp, options.Args...)
	}
	// TODO: start as detached child?!...
	_err = _cmd.Start()
	// TODO: how to call release to detach ?!
	if _err != nil {
		return _cmd, _err
	}

	// TODO: -> https://stackoverflow.com/questions/23031752/start-a-process-in-go-and-detach-from-it

	if options.Release {
		_cmd.Process.Release()
	}
	return _cmd, nil
}
