package connection

import (
	"github.com/gin-gonic/gin"
	"github.com/kyaxcorp/go-core/core/logger/model"
)

func Middleware(logger *model.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		New().SetGinContext(context).SetLogger(logger).Process()
	}
}
