package connection

import "github.com/gin-gonic/gin"

func (c *ConnDetails) SetGinContext(ctx *gin.Context) *ConnDetails {
	c.C = ctx
	return c
}
