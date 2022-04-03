package console

import (
	"github.com/kyaxcorp/go-core/core/console/command"
	"github.com/spf13/cobra"
)

var StopBrokerServer = &command.AddCmd{
	OnExecute: func(cmd *command.AddCmd) {
		cmd.StopProcess()
	},
	OnGetProcessName: func(cmd *command.AddCmd) string {
		brokerName := getInstanceName(cmd.Args)
		return cmd.ProcessName + "_" + brokerName
	},
	OnCreate: func(cmd *command.AddCmd) {
		cmd.Command.Args = cobra.MaximumNArgs(1) // Only 1 arguments -> broker name
	},
	ProcessName: "broker",
	Cmd:         "broker:stop",
	Name:        "Stop - Broker Server",
}
