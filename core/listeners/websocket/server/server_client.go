package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/authentication"
	"github.com/kyaxcorp/go-core/core/listeners/websocket/server/msg"
	"github.com/rs/zerolog"
	"io"
	"strconv"
	"time"
)

// readPump pumps messages from the websocket connection to the registrationHub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	info := func() *zerolog.Event {
		return c.LInfoF("readPump")
	}
	debug := func() *zerolog.Event {
		return c.LDebugF("readPump")
	}
	_error := func() *zerolog.Event {
		return c.LErrorF("readPump")
	}

	// Defer
	defer func() {
		_err := c.Disconnect()
		_error().Err(_err).Msg("disconnect error")
	}()

	info().Msg("running...")
	defer info().Msg("closing...")

	var _err error
	var messageType int
	var message []byte

	// Set options

	// Set the maximum message size for reading...
	c.conn.SetReadLimit(int64(c.server.MaxMessageSize.Get()))

	// Set the specific time when the read is not possible anymore and the connection
	// should be considered dead/corrupt
	_err = c.conn.SetReadDeadline(time.Now().Add(c.server.PongWait.Get()))
	// The error from read deadline can be ignored...
	if _err != nil {
		_error().
			Err(_err).
			Msg("set read deadline")
	}

	// Set the pong handler... this handler will be called on some reader will try to
	// call the Read/ReadMessage/NextReader...
	// if we receive a Pong Frame, this handler will be the one reacting on it!
	c.conn.SetPongHandler(func(string) error {
		// if we received a pong, it means, that the client reacts to pings

		// Set again the read deadline after receiving this frame...
		_err = c.conn.SetReadDeadline(time.Now().Add(c.server.PongWait.Get()))

		// The error from read deadline can be ignored...
		if _err != nil {
			_error().
				Err(_err).
				Msg("set read deadline")
		}

		c.lastReceivedPongTime.SetNow()
		c.nrOfReceivedPongs.Inc(1)
		debug().Uint64("pong_nr", c.nrOfReceivedPongs.Get()).Msg("pong received...")
		return nil
	})

	// This is when we receive a ping handler from the client...
	// in this case the server should also respond to these messages...
	c.conn.SetPingHandler(func(string) error {
		debug().Msg("we have received a ping frame from the client, responding with pong frame")
		_err = c.conn.SetWriteDeadline(time.Now().Add(c.server.WriteWait.Get()))
		if _err != nil {
			_error().Err(_err).Msg("SetPingHandler -> set write deadline")
		}

		// Sending the frame
		if _err = c.conn.WriteMessage(msg.Pong, nil); _err != nil {
			// Failed
			c.nrOfFailedSendPongs.Inc(1)
			_error().Err(_err).Msg("SetPingHandler -> failed to send pong")
			return nil
		}
		c.lastSentPongTime.SetNow()
		c.nrOfSentPongs.Inc(1)
		debug().Uint64("pong_nr", c.nrOfSentPongs.Get()).Msg("pong sent success")
		return nil
	})

	// Loop/Wait
	for {
		messageType, message, _err = c.conn.ReadMessage()
		// log.Println("Read a Message...")

		if _err != nil {
			_error().
				Err(_err).
				Int("message_type", messageType).
				Str("received_message", string(message[:])).
				Bytes("received_bytes", message).
				Msg("some error?! closing the connection!?")
			c.setAsClosed()
			if websocket.IsUnexpectedCloseError(_err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				_error().
					Msg("unexpected close error,expected other codes like 1001, 1006")
			}
			break
		}

		// Without a pointer
		receivedMessage := &ReceivedMessage{
			MessageType:   int8(messageType),
			MessageLength: uint(len(message)),
			Message:       message,
		}

		// TODO: should we process here the pong messages?!!

		// Message = bytes.TrimSpace(bytes.Replace(Message, newline, space, -1))

		// On Message Callback!
		// We decide in the callback how fast we want to process the incoming messages!

		// Declared on the server
		c.server.LEvent("start", "OnMessage", nil)
		c.server.onMessage.Scan(func(k string, v interface{}) {
			v.(OnMessage)(receivedMessage, c, c.server)
		})
		c.server.LEvent("finish", "OnMessage", nil)

		// Declared in the upgrader
		if function.IsCallable(c.onMessage) {
			c.onMessage(receivedMessage, c, c.server)
		}

		// We are sending the Message through Channel to broadcast! The registrationHub is another goroutine which sends to other
		// c the Message!

		// c.server.WSRegistrationHub.broadcast <- Message
	}
}

// writePump pumps messages from the registrationHub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	info := func() *zerolog.Event {
		return c.LInfoF("writePump")
	}
	debug := func() *zerolog.Event {
		return c.LDebugF("writePump")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("writePump")
	}
	_error := func() *zerolog.Event {
		return c.LErrorF("writePump")
	}
	// Defer
	defer func() {
		_err := c.Disconnect()
		_error().Err(_err).Msg("disconnect error")
	}()

	info().Msg("running...")
	defer info().Msg("closing...")

	// launch a ticker which will ping...
	// we should also check if pong received... if not... then close the connection
	c.pingTicker = time.NewTicker(c.server.PingPeriod.Get())
	calledClose := false

	var _err error
	var messageType int
	var w io.WriteCloser

	for {
		select {
		// We called the close pump!
		case calledClose = <-c.closeWritePump:
			if calledClose {
				break
			}
		case message, ok := <-c.send:
			// log.Println("Message received to channel!")

			// Well, if we send an empty... message it can give false! so in this case the connection will be closed...
			// it should always be called!
			_err = c.conn.SetWriteDeadline(time.Now().Add(c.server.WriteWait.Get()))

			if _err != nil {
				_error().Err(_err).Msg("set write deadline")
			}

			if !ok {
				// The registrationHub closed the channel.
				// TODO: we should see how we handle this type of error!!! this is an internal problem!
				// We should somehow close the connection!
				// In this case, if we are sending close message, the connection will be closed gracefully!
				// And even if not closed gracefully, the ping/pong mechanism will handle the link and it will close it after all!

				closeCode := int(c.closeCode)
				closeMsg := c.closeMessage
				if c.closeCode == 0 {
					closeCode = 1011 // Internal server error
				}
				if closeMsg == "" {
					closeMsg = "ws server has an internal error"
				}

				_err = c.conn.WriteMessage(
					msg.Close,
					websocket.FormatCloseMessage(closeCode, closeMsg),
				)
				if _err != nil {
					_error().Err(_err).Msg("failed to send close code")
				}
				return
			}

			// TODO: check if the the type is right!

			messageType, _err = strconv.Atoi(string(message[0]))

			switch messageType {
			case msg.Close:
				c.setAsClosed()
				info().Msg("sending close message for graceful disconnect")
				// https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent

				closeCode := int(c.closeCode)
				closeMsg := c.closeMessage
				if c.closeCode == 0 {
					closeCode = 1000 // Normal closure
				}
				if closeMsg == "" {
					closeMsg = "you have being disconnected gracefully"
				}

				if _err = c.conn.WriteMessage(
					msg.Close,
					websocket.FormatCloseMessage(closeCode, closeMsg),
				); _err != nil {
					return
				}
				info().Msg("continuing...")
				continue
			}

			if _err != nil {
				warn().Msg("incorrect message type")
				return
			}

			//log.Println("Getting the writer!", Message[0], messageType)
			// log.Println("send Message ", Message)

			// Choose a Writer by type! (we have defined that first byte is the Message Type!)
			w, _err = c.conn.NextWriter(messageType)
			if _err != nil {
				return
			}
			// log.Println("Writing the Message")

			// Count
			c.nrOfSentMessages.Inc(1)

			// Writing the message to the client
			// We exclude the first byte because it includes the messageType
			_, _err = w.Write(message[1:])
			if _err != nil {
				_error().Err(_err).Msg("failed write bytes")
				// Count
				c.nrOfSentFailedMessages.Inc(1)
				// Even if we failed to write... we should close the writer!
				//return
			} else {
				// Count
				c.nrOfSentSuccessMessages.Inc(1)
			}

			// Closing the writer
			if _err = w.Close(); _err != nil {
				_error().Err(_err).Msg("failed to close the writer")
				return
			}
		case <-c.pingTicker.C:
			// Check when was the last time we have received a pong message...
			// If we haven't received for a specific amount of time the pong message
			// we should close automatically the connection!
			debug().Msg("preparing to send ping")
			if c.lastReceivedPongTime.Get().Unix()-180 > c.lastSentPingTime.Get().Unix() {
				// It's dead...or the client simply doesn't send any pongs...
				// close the connection...
				_error().Err(_err).Msg("ping ticker -> haven't receive for 3 minutes any pong messages from the client... closing now...")
				_err = c.Disconnect()
				warn().Msg(_err.Error())
				return
			}

			// it should always be called!
			_err = c.conn.SetWriteDeadline(time.Now().Add(c.server.WriteWait.Get()))
			if _err != nil {
				_error().Err(_err).Msg("ping ticker -> set write deadline")
			}

			c.lastSendPingTry.SetNow()
			if _err = c.conn.WriteMessage(msg.Ping, nil); _err != nil {
				// Failed
				_error().Err(_err).Msg("ping ticker -> failed to send ping")
				c.nrOfFailedSendPings.Inc(1)
				return
			}

			// Success
			c.lastSentPingTime.SetNow()
			c.nrOfSentPings.Inc(1)
			debug().Uint64("ping_nr", c.nrOfSentPings.Get()).Msg("ping sent success")
		}

		if calledClose {
			//log.Println("WE WANT TO CLOSE!!!")
			break
		}
	}
}

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

// DisconnectGracefully -> set 0 and "" for default values!
func (c *Client) DisconnectGracefully(code uint16, message string) {
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
}

// Disconnect the client forcefully!!
func (c *Client) Disconnect() error {
	info := func() *zerolog.Event {
		return c.LInfoF("Disconnect")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("Disconnect")
	}
	/*_error := func() *zerolog.Event {
		return c.LErrorF("Disconnect")
	}*/
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

	c.pingTicker.Stop()
	c.closeWritePump <- true // Close the write pump!
	c.isDisconnecting.True()
	c.server.WSRegistrationHub.unregister <- c
	// This is a force Close!
	return c.conn.Close()
}

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

func (c *Client) GetAuthTokenID() string {
	return c.authDetails.AuthTokenDetails.TokenID
}

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
