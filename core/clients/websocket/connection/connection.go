package connection

import (
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"net/url"
	"strconv"
)

// NewConnection -> will generate a connection having default values!
func New() *Connection {
	connection := &Connection{}
	if _err := _struct.SetDefaultValues(&connection); _err != nil {
		return nil
	}
	return connection
}

func (c *Connection) GenerateURL() url.URL {
	// If it's encrypted
	scheme := "ws"
	if conv.ParseBool(c.IsSecure) {
		scheme = "wss"
	}
	// Uri path
	uriPath := "/"
	if c.UriPath != "" {
		uriPath = c.UriPath
	}

	// Setting up host
	host := c.Host
	if c.Port != 0 {
		host = host + ":" + strconv.Itoa(int(c.Port))
	}

	u := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   uriPath,
	}

	return u
}
