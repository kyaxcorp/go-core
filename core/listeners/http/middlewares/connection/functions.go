package connection

import (
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/slice"
)

func (c *ConnDetails) getClientIP() string {
	ip := c.C.ClientIP()
	if includes, _ := slice.Includes([]string{"::1", "127.0.0.1"}, ip); includes {
		// Get another ip!
		xRealIP := c.C.Request.Header.Get("X-Real-IP")
		if xRealIP != "" {
			return xRealIP
		}
		xForwardedFor := c.C.Request.Header.Get("X-Forwarded-For")
		if xForwardedFor != "" {
			return xForwardedFor
		}
	}
	return ip
}

func (c *ConnDetails) generateDetails() {
	// Saving the authentication details into Http Connection Context
	//a.C.Request.TLS

	remoteIP := ""
	if remIP, gotIP := c.C.RemoteIP(); gotIP {
		remoteIP = conv.BytesToStr(remIP)
	}

	c.DomainName = "" // TODO:
	c.Host = c.C.Request.Host
	c.ClientIPAddress = c.getClientIP()
	c.RemoteIP = remoteIP
	c.ClientPort = 0 // TODO: we should search for a possibility
	c.UserAgent = c.C.Request.UserAgent()
	c.RemoteAddr = c.C.Request.RemoteAddr
	c.RequestPath = c.C.Request.RequestURI
	c.IsSecure = false // TODO:
	c.Referer = c.C.Request.Referer()
	// TODO Latency?!

	c.C.Set(HttpContextConnDetailsKey, c)
}

func GenerateConnDetails() *ConnDetails {
	connection := &ConnDetails{}
	return connection
}
