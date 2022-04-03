package request_timing

import (
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"time"
)

func GetMiddleware(logger *model.Logger) gin.HandlerFunc {
	return GetHandlerFunc(logger)
}

func GetHandlerFunc(logger *model.Logger) gin.HandlerFunc {
	return func(gin *gin.Context) {
		t := time.Now()

		// Go to next middleware?!
		gin.Next()
		latency := time.Since(t).Milliseconds()
		// Write to log
		logger.Logger.Info().
			Int64("request_process_time_ms", latency).
			Msg(color.Style{color.LightGreen, color.OpBold}.Render("request process timing"))
	}
}
