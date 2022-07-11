package server

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
)

func (s *Server) SetContext(ctx context.Context) {
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}
	s.parentCtx = ctx

	s.NewCancelContext()
}

func (s *Server) NewCancelContext() *Server {
	s.ctx = _context.WithCancel(s.parentCtx)
	return s
}
