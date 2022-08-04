package authentication

import "github.com/gin-gonic/gin"

const DefaultHeaderAuthKey = "Auth-Token"
const DefaultGETAuthKey = "AuthToken"
const DefaultCookieAuthKey = "authtoken"

var DefaultCookieAuthKeys = []string{DefaultCookieAuthKey, "authtoken", "authToken", "auth-token", "auth_token"}
var DefaultGETAuthKeys = []string{DefaultGETAuthKey, "authtoken", "authToken", "auth-token", "auth_token"}

type OnTokenValid func(*Auth)
type OnTokenInvalid func(*Auth)

const HttpContextAuthDetailsKey = "AUTH_DETAILS"

const ByHeader = 1
const ByGetParam = 2
const ByCookie = 3

// New -> This is the Constructor or first function to call!
func New() *Auth {
	auth := &Auth{}
	return auth
}

func checkGETKey(tmpKey []string) bool {
	if tmpKey != nil && len(tmpKey) > 0 && tmpKey[0] != "" {
		return true
	}
	return false
}

func checkHeaderKey(tmpKey []string) bool {
	if tmpKey != nil && len(tmpKey) > 0 && tmpKey[0] != "" {
		return true
	}
	return false
}

func GetAuthDetailsFromCtx(c *gin.Context) *AuthDetails {
	var _authDetails *AuthDetails
	authDetails, ifExists := c.Get(HttpContextAuthDetailsKey)
	if ifExists && authDetails != nil {
		// Set the data into the Client NonPtrObj
		_authDetails = authDetails.(*AuthDetails)
	} else {
		_authDetails = &AuthDetails{}
	}
	return _authDetails
}
