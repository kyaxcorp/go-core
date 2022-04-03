package model

import (
	"github.com/rs/zerolog"
	"io"
)
import "github.com/KyaXTeam/go-core/v2/core/logger/config"

type Logger struct {
	// This is where the entire config lies...
	Config config.Config

	// Usually we would like to create logs to: StdOut, In some file, or event output elsewhere!
	// Logrus has only 1 output,  so we decided to create multiple of them!
	// We should create a default channel with stdout
	// And additional if the user wants to other different outputs
	Logger *zerolog.Logger

	// Main Writer
	MainWriter io.Writer

	// Context is possibly not needed here... because we are not doing any goroutines
	//parentCtx context.Context
	//ctx       _context.CancelCtx
}
