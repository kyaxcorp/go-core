package _bool

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_map_string_interface"
	"github.com/kyaxcorp/go-core/core/helpers/sync/waiter"
	"sync"
)

type Bool struct {
	lock  sync.RWMutex
	value bool
	// TODO: we can add here also OnSet, OnGet events

	// Parent Context...
	ctx context.Context

	// it's needed for WaitFor True & False
	eventTrigger chan Bool

	// Here we store the events
	onChange      *_map_string_interface.MapStringInterface
	onChangeAsync *_map_string_interface.MapStringInterface

	onTrueWaiter  *waiter.Waiter
	onFalseWaiter *waiter.Waiter
}
