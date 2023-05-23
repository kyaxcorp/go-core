package menu

import (
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/bootstrap/register_service"
	"github.com/kyaxcorp/go-core/core/console/command"
	"github.com/kyaxcorp/go-core/core/console/working_stage"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/lock"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/kyaxcorp/go-core/core/helpers/process/shutdown"
	"github.com/kyaxcorp/go-core/core/logger/coreLog"
	"github.com/spf13/cobra"
	"strings"
)

func (m *Menu) AddCommands(c []*command.AddCmd) *Menu {
	for _, cmd := range c {
		m.AddCommand(cmd)
	}
	return m
}

// AddCommand -> Adding commands
func (m *Menu) AddCommand(c *command.AddCmd) *Menu {

	// Define the cobra command
	var cobraCmd *cobra.Command
	if c.Command != nil {
		cobraCmd = c.Command
	} else {
		cobraCmd = &cobra.Command{}
	}

	cobraCmd.Use = c.Cmd
	cobraCmd.Short = c.GetCmdShortName()
	cobraCmd.Long = c.GetCmdLongName()
	cobraCmd.Run = func(cmd *cobra.Command, args []string) {
		// Adding here the ref -> but it should  already have because we have added it before!
		//c.Command = cmd

		isDev := true
		if strings.Contains(c.Cmd, "--prod") {
			isDev = false
		}
		if strings.Contains(c.Cmd, "--dev") {
			isDev = true
		}

		// set the global stage
		working_stage.SetStage(isDev)

		// Set the arguments to the command
		c.Args = args

		// Setting here Menu/Cobra context!
		cmdCtx := _context.WithCancel(cmd.Context())
		c.SetCancelContext(cmdCtx)

		// Run Event OnRun
		if function.IsCallable(c.OnRun) {
			c.OnRun(c)
		}

		// Check if it's a daemon
		if m.IsDaemon() {
			if function.IsCallable(c.OnDaemon) {
				c.OnDaemon(c)
			}

			// TODO: what happens with StdIn and StdOut when launching the app in the background!?
			// TODO: the primary app remains online! and child program becomes dependent of the primary one!
			// TODO: it should not be dependent!

			coreLog.Info().Msg("running in background")
			//log.Println("Running in background")
			_command, _err := m.RunInternalCommand(InternalCommandOptions{
				Args:                 []string{c.Cmd},
				Release:              true,
				RunDaemonFromExecDir: c.RunDaemonFromExecDir,
			})
			if _err != nil {
				// TODO: we should handle if we can't start!
				coreLog.Error().Err(_err).Msg("failed to start command... ")
				return
			}

			//log.Println("PID", command.Check.Pid)
			coreLog.Info().Int("pid", _command.Process.Pid).Msg("getting pid")
		} else {
			if c.LockProcess {
				if isLockAcquired, lockErr := lock.FLock(c.GetProcessLockName(), false); !isLockAcquired || lockErr != nil {
					// handle locking error
					//log.Println("Failed to lock the process!")
					coreLog.Warn().Msg(color.Style{color.LightYellow}.Render("failed to lock the process"))
					return
				}
			}

			// Destruct when leaving this...
			defer func() {
				c.Destructor()
			}()

			// Destruct on signal!

			shutdown.OnShutdown(func() {
				cmdCtx.Cancel() // Send Signal,Cancel by context!
				c.Destructor()
			})
			go shutdown.MonitorSigMessages()

			// Generate PID only if needed!
			if c.GeneratePID {
				// TODO: handle error!
				c.GenPID()
			}

			// -------Run Services--------\\

			if c.EnableStartupServices {
				// Run Broker ClientsStatus
				//if c.StartupCoreServices.BrokerClients {
				// Register the broker client service
				//brokerClientService.RegisterBrokerService()
				//}

				// Run the registered services
				register_service.RunRegisteredServices()
			}
			// -------Run Services--------\\

			// Start Execution!
			if function.IsCallable(c.OnExecute) {
				//c.OnExecute(c, args)
				c.OnExecute(c)
			}
		}
	}

	c.Command = cobraCmd
	if function.IsCallable(c.OnCreate) {
		c.OnCreate(c)
	}
	m.commands[c.Cmd] = cobraCmd
	// Adding Command to Cobra
	m.RootCmd.AddCommand(cobraCmd)

	return m
}
