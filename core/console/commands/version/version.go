package version

import (
	"fmt"
	"github.com/KyaXTeam/go-core/v2/core/console/command"
	"github.com/KyaXTeam/go-core/v2/core/helpers/json"
	"github.com/KyaXTeam/go-core/v2/core/helpers/version"
)

var ShowVersion = &command.AddCmd{
	GeneratePID: false,
	LockProcess: false,
	ProcessName: "version",
	Cmd:         "version",
	Name:        "Application Version",
	OnExecute: func(cmd *command.AddCmd) {
		// Print information !
		appVersion := version.GetAppVersion()
		fmt.Print(appVersion.Version)
	},
}

var ShowVersionJSON = &command.AddCmd{
	GeneratePID: false,
	LockProcess: false,
	ProcessName: "version:json",
	Cmd:         "version:json",
	Name:        "Application Version Json Output",
	OnExecute: func(cmd *command.AddCmd) {
		// Print information !
		appVersion := version.GetAppVersion()
		fmt.Print(json.Encode(appVersion))
	},
}
