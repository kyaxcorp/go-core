package authentication

import "github.com/gin-gonic/gin"

func (a *Auth) GetHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the context to the struct!
		a.C = c

		receivedAuthToken := ""
		authByType := 0
		authTypeKeyName := ""
		// Start first by checking in headers

		if len(a.authHeaderKeys) == 0 {
			// If there is nothing defined, we set the default one!
			a.authHeaderKeys = []string{DefaultHeaderAuthKey}
		}

		if len(a.authHeaderKeys) > 0 {
			// If there are defined multiple auth header keys!
			for _, authKey := range a.authHeaderKeys {

				if tmpKey, ok := c.Request.Header[authKey]; ok {
					if checkHeaderKey(tmpKey) {
						receivedAuthToken = tmpKey[0]
						authTypeKeyName = authKey
						authByType = ByHeader
						break
					}
				}
			}
		}

		// If nothing found check in GET params, but check if it's a GET Method

		reqMethod := c.Request.Method

		if (reqMethod == "" || reqMethod == "GET") && receivedAuthToken == "" {
			// Try searching here
			getParams := c.Request.URL.Query()
			if getParams != nil {
				if len(a.authGetKeys) == 0 {
					// If there is nothing defined, we set the default one!
					a.authGetKeys = DefaultGETAuthKeys
				}

				if len(a.authGetKeys) > 0 {
					// If there are defined multiple auth header keys!
					for _, authKey := range a.authGetKeys {
						if tmpKey, ok := getParams[authKey]; ok {
							if checkGETKey(tmpKey) {
								receivedAuthToken = tmpKey[0]
								authTypeKeyName = authKey
								authByType = ByGetParam
								break
							}
						}
					}
				}
			}
		}

		// By Cookies
		if receivedAuthToken == "" {
			// Check by cookie
			if len(a.authCookieKeys) == 0 {
				// If there is nothing defined, we set the default one!
				a.authCookieKeys = DefaultCookieAuthKeys
			}

			if len(a.authCookieKeys) > 0 {
				// If there are defined multiple auth header keys!
				for _, authKey := range a.authCookieKeys {
					cookie, err := c.Request.Cookie(authKey)
					if err != nil {
						continue
					}

					if cookie.Value == "" {
						continue
					}
					authByType = ByCookie
					authTypeKeyName = authKey
					receivedAuthToken = cookie.Value
					break
				}
			}
		}

		if receivedAuthToken == "" {
			a.onTokenInvalid(a)
			return
		}
		// Token it's ok!
		a.authToken = receivedAuthToken
		a.authTypeKeyName = authTypeKeyName
		a.authType = uint8(authByType) // By what type it has authenticated
		a.onTokenValid(a)

		// log.Println(receivedAuthToken)
	}
}
