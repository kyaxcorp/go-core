package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/authentication"
	"strconv"
	"time"
)

func (c *Client) GetConnectTime() time.Time {
	return c.connectTime
}

func (c *Client) GetConnectedTimeSeconds() int64 {
	// TODO : calculate
	return 1
}

func (c *Client) setAsClosed() {
	// Setting as connection closed!
	c.isClosed.True()
}

// TODO: review this function!
// DisconnectGracefully -> set 0 and "" for default values!
/*func (c *Client) DisconnectGracefully(code uint16, message string) {
	// We send to the client that we want to close the connection!
	// And we should receive response back, and after that the disconnect will be called!

	c.setAsClosed()

	if code > 0 {
		c.closeCode = code
	}
	if message != "" && len(message) > 0 {
		c.closeMessage = message
	}
	// send through channel!
	c.send <- []byte(strconv.Itoa(msg.Close))
}*/

// TODO: review this function!
// Disconnect the client forcefully!!
/*func (c *Client) Disconnect() error {
	info := func() *zerolog.Event {
		return c.LInfoF("Disconnect")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("Disconnect")
	}

	info().Msg("calling...")
	defer info().Msg("leaving...")

	if c.IsDisconnecting() {
		warn().Msg("already disconnecting...")
		return nil
	}
	c.setAsClosed()

	// On Close callback
	c.server.onClose.Scan(func(k string, v interface{}) {
		v.(OnClose)(c, c.server)
	})

	info().Msg("closing the client connection...")

	c.writeTicker.Stop()
	c.isDisconnecting.True()
	// This is a force Close!
	return c.conn.Close()
}*/

func (c *Client) IsDisconnecting() bool {
	return c.isDisconnecting.Get()
}

func (c *Client) GetConnectionID() uint64 {
	return c.connectionID
}

func (c *Client) GetHttpContext() *gin.Context {
	return c.httpContext
}

func (c *Client) GetDeviceID() string {
	return c.authDetails.DeviceDetails.DeviceID
}

func (c *Client) GetDeviceUUID() string {
	return c.authDetails.DeviceDetails.DeviceUUID
}

func (c *Client) GetUserID() string {
	return c.authDetails.UserDetails.UserID
}

func (c *Client) GetAuthToken() string {
	return c.authDetails.AuthTokenDetails.Token
}

/*func (c *Client) GetAuthTokenID() uint64 {
	return c.authDetails.AuthTokenDetails.TokenID
}*/

func (c *Client) GetIPAddress() string {
	return c.connDetails.ClientIPAddress
}

func (c *Client) GetRemoteIP() string {
	return c.connDetails.RemoteIP
}

func (c *Client) GetRequestPath() string {
	return c.connDetails.RequestPath
}

func (c *Client) GetTokenExpirationTime() time.Time {
	return c.authDetails.AuthTokenDetails.ExpireDate
}

func (c *Client) GetAuthDetails() *authentication.AuthDetails {
	return c.authDetails
}

// This generates an unique ID for the Message that will be sent!
func (c *Client) genPayloadID() string {

	c.randomPayloadID.Inc(1)
	if c.randomPayloadID.Get() > 65500 {
		// Reset it!
		c.randomPayloadID.Set(1)
	}
	id := c.randomPayloadID.Get()

	// Prefix "S" as Server + Connection ID + Random payload ID + Nano Time
	return "s_" + strconv.FormatUint(c.connectionID, 10) + "_" +
		strconv.Itoa(int(id)) + "_" +
		strconv.FormatInt(time.Now().UnixNano(), 10)
}

// Set custom Data to client connection!
func (c *Client) Set(key string, value interface{}) *Client {
	//c.customData[key] = value
	c.customData.Set(key, value)
	return c
}

// Get custom data from the client connection!
func (c *Client) Get(key string) interface{} {
	return c.customData.Get(key)
	/*if val, ok := c.customData[key]; ok {
		return val
	}
	return nil*/
}
