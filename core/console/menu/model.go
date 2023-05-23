package menu

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/spf13/cobra"
)

type Menu struct {
	cfgFile     string
	userLicense string

	RootCmd *cobra.Command

	isDaemon bool
	commands map[string]*cobra.Command

	parentCtx context.Context     // This is the Parent context!
	ctx       *_context.CancelCtx // This is the cancel context! -> which will cancel any execution!?
}
