package menu

import (
	"context"
	"fmt"
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/bootstrap/register_service"
	brokerClientService "github.com/kyaxcorp/go-core/core/clients/broker/services"
	"github.com/kyaxcorp/go-core/core/console/command"
	"github.com/kyaxcorp/go-core/core/console/commands/version"
	"github.com/kyaxcorp/go-core/core/console/working_stage"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem/lock"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"github.com/kyaxcorp/go-core/core/helpers/process/shutdown"
	"github.com/kyaxcorp/go-core/core/logger/appLog"
	"github.com/kyaxcorp/go-core/core/services/broker/console"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Menu struct {
	cfgFile     string
	userLicense string
	RootCmd     *cobra.Command
	versionCmd  *cobra.Command
	isDaemon    bool
	commands    map[string]*cobra.Command

	parentCtx context.Context     // This is the Parent context!
	ctx       *_context.CancelCtx // This is the cancel context! -> which will cancel any execution!?
}

func New(ctx context.Context) *Menu {
	// This is the executable checksum, it will help the user to know which config and appdata folder
	// the app owns
	execChecksum := hash.MD5(filepath.Base(os.Args[0]))

	return &Menu{
		cfgFile:     "",
		userLicense: "",
		RootCmd: &cobra.Command{
			Use:   "ARGUMENT",
			Short: "Main CLI -> " + execChecksum,
			Long:  `Main CLI -> ` + execChecksum,
		},
		parentCtx: ctx,                      // This is the Root Context
		ctx:       _context.WithCancel(ctx), // Create the cancel Context for this menu

		versionCmd: &cobra.Command{
			Use:   "version",
			Short: "Print the version of the APP",
			Long:  `Print the version of the APP`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("1.0.0")
			},
		},
		commands: make(map[string]*cobra.Command),
	}
}

// Execute executes the root Command.
func (m *Menu) Execute() error {
	m.init()
	//return m.RootCmd.Execute()
	// We are using cancel context here!
	return m.RootCmd.ExecuteContext(m.ctx.Context())
}

func (m *Menu) RunInternalCommand(arg ...string) (*exec.Cmd, error) {
	currentApp, _ := os.Executable()
	//command := exec.Command(currentApp, arg...)
	command := exec.CommandContext(_context.GetDefaultContext(), currentApp, arg...)
	// TODO: start as detached child?!...
	_err := command.Start()
	// TODO: how to call release to detach ?!
	if _err != nil {
		return command, _err
	}

	// TODO: -> https://stackoverflow.com/questions/23031752/start-a-process-in-go-and-detach-from-it
	command.Process.Release()
	return command, nil
}

func (m *Menu) IsDaemon() bool {
	return m.isDaemon
}

func (m *Menu) init() {
	// TODO: check if we need to use config!
	cobra.OnInitialize(m.initConfig)

	//Adding additional Core Commands
	m.AddCommands([]*command.AddCmd{
		console.StartBrokerServer,
		console.StopBrokerServer,
		version.ShowVersion,     // Show app version
		version.ShowVersionJSON, // Show app Version in JSON format
	})

	// Adding options
	m.RootCmd.PersistentFlags().BoolVarP(
		&m.isDaemon,
		"daemon",
		"d",
		false,
		"Run Command in background",
	)
}

func (m *Menu) initConfig() {
	if m.cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(m.cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
func (m *Menu) AddCommands(c []*command.AddCmd) *Menu {
	for _, cmd := range c {
		m.AddCommand(cmd)
	}
	return m
}

// AddCommand -> Adding commands
func (m *Menu) AddCommand(c *command.AddCmd) *Menu {

	// Define the cobra command
	cobraCmd := &cobra.Command{
		Use:   c.Cmd,
		Short: c.GetCmdShortName(),
		Long:  c.GetCmdLongName(),
		Run: func(cmd *cobra.Command, args []string) {
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

				appLog.Info().Msg("running in background")
				//log.Println("Running in background")
				_command, _err := m.RunInternalCommand(c.Cmd)
				if _err != nil {
					// TODO: we should handle if we can't start!
					appLog.Error().Err(_err).Msg("failed to start command... ")
					return
				}

				//log.Println("PID", command.Process.Pid)
				appLog.Info().Int("pid", _command.Process.Pid).Msg("getting pid")
			} else {
				if c.LockProcess && !lock.FLock(c.GetProcessLockName(), false) {
					// handle locking error
					//log.Println("Failed to lock the process!")
					appLog.Warn().Msg(color.Style{color.LightYellow}.Render("failed to lock the process"))
					return
				}

				// Destruct when leaving this...
				defer func() {
					c.Destructor()
				}()

				// Destruct on signal!
				go shutdown.MonitorSigMessages(func() {
					cmdCtx.Cancel() // Send Signal,Cancel by context!
					c.Destructor()
				})

				// Generate PID only if needed!
				if c.GeneratePID {
					c.GenPID()
				}

				// -------Run Services--------\\

				if c.EnableStartupServices {
					// Run Broker Clients
					if c.StartupCoreServices.BrokerClients {
						// Register the broker client service
						brokerClientService.RegisterBrokerService()
					}

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
		},
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
