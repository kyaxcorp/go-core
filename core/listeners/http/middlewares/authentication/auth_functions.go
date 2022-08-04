package authentication

import (
	"github.com/gin-gonic/gin"
)

// ByHeaderKeys > Search in header specific keys
func (a *Auth) ByHeaderKeys(authHeaderKeys []string) *Auth {
	a.authHeaderKeys = authHeaderKeys
	return a
}

// ByCookies -> Search in Cookies
func (a *Auth) ByCookies(authCookieKeys []string) *Auth {
	a.authCookieKeys = authCookieKeys
	return a
}

// ByGetParams -> Search in GET method existing key params
func (a *Auth) ByGetParams(authGetKeys []string) *Auth {
	a.authGetKeys = authGetKeys
	return a
}

// OnTokenValid -> On Valid!
func (a *Auth) OnTokenValid(callback OnTokenValid) *Auth {
	if callback != nil {
		a.onTokenValid = callback
	}
	return a
}

// OnTokenInValid -> On Invalid
func (a *Auth) OnTokenInValid(callback OnTokenInvalid) *Auth {
	if callback != nil {
		a.onTokenInvalid = callback
	}
	return a
}

func (a *Auth) GetToken() string {
	return a.authToken
}

func (a *Auth) GetAuthType() uint8 {
	return a.authType
}

func (a *Auth) GetAuthTypeKeyName() string {
	return a.authTypeKeyName
}

func (a *Auth) SetAuthDetails(details *AuthDetails) {
	// Saving the authentication details into Http Connection Context
	a.C.Set(HttpContextAuthDetailsKey, details)
	//ctx := context.WithValue(a.C.Request.Context(), HttpContextAuthDetailsKey, details)
	//a.C.Request = a.C.Request.WithContext(ctx)
}

func (a *Auth) Abort(code int, httpCode int, msg string) {
	a.C.JSON(httpCode, gin.H{
		//"message": "You are not authorized to use this resource!",
		// TODO: check response structure! should be the same as in PHP!
		"status":  false,
		"code":    code,
		"message": msg,
		"data":    nil,
	})
	a.C.Abort()
}
