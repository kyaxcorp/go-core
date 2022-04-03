package connection

import "github.com/gin-gonic/gin"

const HttpContextConnDetailsKey = "CONN_DETAILS"

// Connection Details
type ConnDetails struct {
	Host string
	// Called domain name
	DomainName      string
	ClientIPAddress string
	RemoteIP        string
	ClientPort      int
	UserAgent       string
	RemoteAddr      string
	RequestPath     string
	// Is it through SSL
	IsSecure bool
	Referer  string

	// Connection context
	C *gin.Context
}
