package server

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
)

func (s *Server) SetContext(ctx context.Context) {
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}
	s.parentCtx = ctx
}
