package console

import (
	"github.com/KyaXTeam/go-core/v2/core/console/command"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/services/broker"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var StartPushNotificationsServer = &command.AddCmd{
	GeneratePID: false,
	LockProcess: false,
	ProcessName: "push_notifications_server",
	// Run the broker client service
	EnableStartupServices: true,
	StartupCoreServices: command.StartupCoreServices{
		BrokerClients: true,
	},

	OnGetProcessName: func(cmd *command.AddCmd) string {
		brokerName := getBrokerName(cmd.Args)
		return cmd.ProcessName + "_" + brokerName
	},
	Cmd:  "broker:start",
	Name: "Start - Broker Server",
	OnExecute: func(cmd *command.AddCmd) {
		brokerName := getBrokerName(cmd.Args)
		brk, err := broker.GetBroker(_context.GetDefaultContext(), brokerName)
		if err != nil {
			log.Fatalln("failed to generate broker...", err)
		}

		err = brk.Start()
		if err != nil {
			log.Fatalln("failed to start broker...", err)
		}
		for {
			time.Sleep(time.Second)
		}
	},
	OnCreate: func(cmd *command.AddCmd) {
		cmd.Command.Args = cobra.MaximumNArgs(1) // Only 1 arguments -> broker name
	},
}
