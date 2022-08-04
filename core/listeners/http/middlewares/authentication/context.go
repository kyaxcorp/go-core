package authentication

import "github.com/gin-gonic/gin"

func (a *Auth) SetGinContext(ctx *gin.Context) *Auth {
	a.C = ctx
	return a
}
