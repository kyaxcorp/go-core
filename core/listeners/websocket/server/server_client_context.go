package server

import "github.com/kyaxcorp/go-core/core/helpers/_context"

func (c *Client) GetCancelContext() *_context.CancelCtx {
	return c.ctx
}
