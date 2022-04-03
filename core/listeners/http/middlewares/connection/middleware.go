package connection

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/logger/model"
)

func GetMiddleware(logger *model.Logger) gin.HandlerFunc {
	conn := GenerateConnDetails()
	return conn.GetHandlerFunc(logger)
}

func (c *ConnDetails) GetHandlerFunc(logger *model.Logger) gin.HandlerFunc {
	return func(gin *gin.Context) {

		// Set the context to the struct!
		c.C = gin
		c.generateDetails()
		// Debug the connection

		logger.Logger.Info().
			Str("host", c.Host).
			Str("client_ip", c.ClientIPAddress).
			Int("client_port", c.ClientPort).
			Str("user_agent", c.UserAgent).
			Str("remote_addr", c.RemoteAddr).
			Str("request_path", c.RequestPath).
			Str("referer", c.Referer).
			Msg(color.Style{color.LightGreen, color.OpBold}.Render("new connection"))
	}
}
