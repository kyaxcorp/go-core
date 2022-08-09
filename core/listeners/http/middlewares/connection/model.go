package connection

import (
	"github.com/gin-gonic/gin"
	"github.com/kyaxcorp/go-core/core/logger/model"
)

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

	Logger *model.Logger

	// Connection context
	C *gin.Context
}
