package command

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/lock"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/kyaxcorp/go-core/core/helpers/process"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

type CmdOnDaemon func(*AddCmd)
type CmdOnExecute func(cmd *AddCmd)
type CmdOnRun func(*AddCmd)
type CmdOnCreate func(*AddCmd)
type OnGetProcessName func(*AddCmd) string

// StartupCoreServices -> Here we add defined/core services, don't know how right is to do it here... but anyway...
type StartupCoreServices struct {
	// Will launch the broker clients on this command!
	BrokerClients bool
}

type AddCmd struct {
	//Callbacks
	// That's when daemon it's being called!
	OnDaemon CmdOnDaemon
	// That's when execution it's being called!
	OnExecute CmdOnExecute

	// That's when we enter Run Mode
	OnRun CmdOnRun
	// That's in create mode
	OnCreate         CmdOnCreate
	OnGetProcessName OnGetProcessName

	Args       []string
	ArgsPolicy string

	// This will be the process name based on what the lock will be made!
	ProcessName string
	// Command -> websocket:start
	Cmd string
	//Names
	ShortName string
	LongName  string
	Name      string

	// it will run the executable from its current location! ./APP
	RunDaemonFromExecDir bool

	// This enables startup services
	EnableStartupServices bool
	StartupCoreServices   StartupCoreServices

	Command *cobra.Command

	// should the process be locked!?
	LockProcess bool
	// Generate PID based on this Command!?
	GeneratePID bool

	ctx *_context.CancelCtx
}

func (c *AddCmd) SetCancelContext(ctx *_context.CancelCtx) {
	c.ctx = ctx
}

func (c *AddCmd) GetContext() context.Context {
	return c.ctx.Context()
}

func (c *AddCmd) GetCancelFunc() context.CancelFunc {
	return c.ctx.Cancel
}

func (c *AddCmd) GetProcessName() string {
	if function.IsCallable(c.OnGetProcessName) {
		return c.OnGetProcessName(c)
	}
	return c.ProcessName
}

func (c *AddCmd) GetCmdShortName() string {
	if c.ShortName != "" {
		return c.ShortName
	}
	if c.LongName != "" {
		return c.LongName
	}
	if c.Name != "" {
		return c.Name
	}
	return ""
}

func (c *AddCmd) GetCmdLongName() string {
	if c.LongName != "" {
		return c.LongName
	}
	if c.ShortName != "" {
		return c.ShortName
	}
	return ""
}

func (c *AddCmd) GetProcessLockName() string {
	// Shouldn't we somehow add more info here for the lock?
	return "console_command_" + c.Cmd
}

func (c *AddCmd) PIDDestroy() (bool, error) {
	return process.PIDDestroy(c.GetProcessName())
}

func (c *AddCmd) GenPID() (bool, error) {
	return process.GeneratePID(c.GetProcessName())
}

func (c *AddCmd) GetExistingPID() (string, error) {
	return process.GetExistingPID(c.GetProcessName())
}

func (c *AddCmd) StopProcess() (bool, error) {
	pid, _err := c.GetExistingPID()
	if _err != nil {
		return false, _err
	}

	//log.Println(pid)
	if pid == "" {
		//log.Println("PID Empty...")
		return false, define.Err(0, "pid empty")
	}

	// Call SIGTERM

	_pid, _err := strconv.Atoi(pid)
	p, _err := os.FindProcess(_pid)
	if _err != nil {
		//log.Println("Process not found!")
		return false, define.Err(0, "process not found -> ", _err.Error())
	}

	_err = p.Signal(os.Interrupt)
	if _err != nil {
		//log.Println("Failed to Interrupt process!!")
		return false, define.Err(0, "failed to interrupt process -> ", _err.Error())
	}
	//log.Println("Process stopped!")
	return true, nil
}

func NewCmd() *AddCmd {
	return &AddCmd{}
}

func (c *AddCmd) Destructor() {
	// log.Println("Destructor...")
	if c.GeneratePID {
		c.PIDDestroy()
	}
	if c.LockProcess {
		lock.FRelease(c.GetProcessLockName())
	}
}
