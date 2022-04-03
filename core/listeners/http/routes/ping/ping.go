package ping

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Ping(server *gin.Engine) {
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"now": time.Now(),
		})
	})
}
