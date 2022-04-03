package filter

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
)

func (f *Input) SetContext(ctx context.Context) *Input {
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}
	f.ctx = ctx
	return f
}

func (f *Input) checkContext() *Input {
	if f.ctx == nil {
		f.ctx = _context.GetDefaultContext()
	}
	return f
}
