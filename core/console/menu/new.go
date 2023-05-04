package menu

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/process/name"
	"github.com/spf13/cobra"
	"go.szostok.io/version/extension"
)

func New(ctx context.Context) *Menu {
	// This is the executable checksum, it will help the user to know which config and appdata folder
	// the app owns
	execChecksum := name.GetCurrentProcessCleanMD5ExecName()

	rootCmd := &cobra.Command{
		Use:   "ARGUMENT",
		Short: "Main CLI -> " + execChecksum,
		Long:  `Main CLI -> ` + execChecksum,
	}

	// Add version info
	rootCmd.AddCommand(
		// 1. Register the 'version' command
		extension.NewVersionCobraCmd(
			// 2. Explicitly enable upgrade notice
			extension.WithUpgradeNotice("repo-owner", "repo-name"),
		),
	)

	return &Menu{
		cfgFile:     "",
		userLicense: "",
		commands:    make(map[string]*cobra.Command),

		parentCtx: ctx,                      // This is the Root Context
		ctx:       _context.WithCancel(ctx), // Create the cancel Context for this menu

		RootCmd: rootCmd,
	}
}
