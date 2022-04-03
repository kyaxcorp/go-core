package broker

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/authentication"
	"github.com/rs/zerolog"
)

func (b *Broker) CheckIsAuthenticated() gin.HandlerFunc {
	return authentication.
		CheckIsAuthenticated().
		//ByHeaderKeys([]string{"Auth-Token"}).
		//ByGetParams([]string{"AuthToken"}).
		OnTokenInValid(func(a *authentication.Auth) {
			warn := func() *zerolog.Event {
				return b.LWarnF("CheckIsAuthenticated")
			}
			warn().Msg("authentication token is not available")

			a.Abort(
				1000,
				403,
				"Authentication token is not available!",
			)
		}).
		OnTokenValid(func(a *authentication.Auth) {
			// Check here
			info := func() *zerolog.Event {
				return b.LInfoF("CheckIsAuthenticated")
			}

			token := a.GetToken()
			if token != b.config.AuthToken {
				info().Str("received_token", token).Msg(color.Style{color.LightRed, color.Bold}.Render("authentication failed"))
				a.Abort(
					2000,
					403,
					"Authentication Token is Invalid!",
				)
			}
			info().Msg(color.Style{color.LightGreen, color.Bold}.Render("authentication success"))

		}).
		GetHandlerFunc()
}
