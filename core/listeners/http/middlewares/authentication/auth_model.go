package authentication

import "github.com/gin-gonic/gin"

type Auth struct {
	// The one that has being found!
	authToken string
	// By Cookie, By Get Param, By header Value
	authType uint8
	// This is the key name where the value has been detected/found/extracted
	authTypeKeyName string
	// These are the keys through which are searched for the token!!
	authHeaderKeys []string
	authGetKeys    []string
	authCookieKeys []string
	onTokenValid   OnTokenValid
	onTokenInvalid OnTokenInvalid
	// This is the context
	C *gin.Context
}
