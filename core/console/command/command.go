package command

import (
	"context"
	"fmt"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/filesystem/lock"
	"github.com/KyaXTeam/go-core/v2/core/helpers/function"
	"github.com/KyaXTeam/go-core/v2/core/helpers/process"
	"github.com/spf13/cobra"
	"log"
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
	return c.Cmd
}

func (c *AddCmd) PIDDestroy() bool {
	return process.PIDDestroy(c.GetProcessName())
}

func (c *AddCmd) GenPID() bool {
	return process.GeneratePID(c.GetProcessName())
}

func (c *AddCmd) GetExistingPID() string {
	return process.GetExistingPID(c.GetProcessName())
}

func (c *AddCmd) StopProcess() bool {
	pid := c.GetExistingPID()
	log.Println(pid)
	if pid == "" {
		log.Println("PID Empty...")
		return false
	}

	// Call SIGTERM

	_pid, err := strconv.Atoi(pid)
	p, err := os.FindProcess(_pid)
	if err != nil {
		log.Println("Process not found!")
		return false
	}

	err = p.Signal(os.Interrupt)
	if err != nil {
		log.Println("Failed to Interrupt process!!")
		return false
	}
	fmt.Println("Process stopped!")
	return true
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
