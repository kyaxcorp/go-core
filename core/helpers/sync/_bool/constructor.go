package _bool

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_map_string_interface"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/waiter"
	"sync"
)

func New() *Bool {
	return getDefaultConstructor(&Bool{
		value: false,
	})
}

func NewContext(ctx context.Context) *Bool {
	return getDefaultConstructor(&Bool{
		ctx:   ctx,
		value: false,
	})
}

func NewVal(v bool) *Bool {
	return getDefaultConstructor(&Bool{
		value: v,
	})
}
func NewValContext(v bool, ctx context.Context) *Bool {
	return getDefaultConstructor(&Bool{
		ctx:   ctx,
		value: v,
	})
}

func getDefaultConstructor(b *Bool) *Bool {
	// Set default context if missing...
	if b.ctx == nil {
		b.ctx = _context.GetDefaultContext()
	}

	b.eventTrigger = make(chan Bool)
	b.onTrueWaiter = waiter.New(b.ctx)
	b.onFalseWaiter = waiter.New(b.ctx)
	b.lock = sync.RWMutex{}
	b.onChange = _map_string_interface.New()
	b.onChangeAsync = _map_string_interface.New()
	return b
}
