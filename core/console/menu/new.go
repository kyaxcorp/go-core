package menu

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/process/name"
	"github.com/spf13/cobra"
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

	return &Menu{
		cfgFile:     "",
		userLicense: "",
		commands:    make(map[string]*cobra.Command),

		parentCtx: ctx,                      // This is the Root Context
		ctx:       _context.WithCancel(ctx), // Create the cancel Context for this menu

		RootCmd: rootCmd,
	}
}
