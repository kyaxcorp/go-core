package connection

import "github.com/kyaxcorp/go-core/core/logger/model"

func (c *ConnDetails) SetLogger(logger *model.Logger) *ConnDetails {
	c.Logger = logger
	return c
}
