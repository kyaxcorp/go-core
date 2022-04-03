package client

import (
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/listeners/websocket/server/msg"
	"github.com/rs/zerolog"
	"strconv"
)

// Disconnect -> It will send the disconnect packet, but it will not wait for it! Forcefully Stops the connection!
// This function can also be called by the user! to terminate the connection!
func (c *Client) Disconnect() {
	// Connection Name
	wasConnectedTo := c.GetConnectedTo()

	warn := func() *zerolog.Event {
		return c.LWarnF("Disconnect").Int("connected_to", wasConnectedTo)
	}
	_error := func() *zerolog.Event {
		return c.LErrorF("Disconnect").Int("connected_to", wasConnectedTo)
	}
	_event := func(eventType string, eventName string) *zerolog.Event {
		return c.LEventF(eventType, eventName, "Disconnect").Int("connected_to", wasConnectedTo)
	}
	info := func() *zerolog.Event {
		return c.LInfoF("Disconnect").Int("connected_to", wasConnectedTo)
	}

	info().Msg("entering...")
	defer info().Msg("leaving...")

	// Set that we are disconnecting and set it to true
	if c.isDisconnecting.IfFalseSetTrue() {
		// It's already disconnecting
		warn().Msg("already disconnecting...")
		return
	}

	// Run as defer, because we don't want somewhere the process to be blocked...
	// defer somehow guarantee that will be run!
	defer func() {
		// Set that we are disconnected
		c.isConnected.False()
		// Set that we are not disconnecting anymore
		c.isDisconnecting.False()
		// Set that we are not connected anywhere anymore
		c.connectedTo.Set(-1) // -1 meaning that is not connected anywhere!
	}()

	// Connection config
	connectedTo := c.GetConnection(wasConnectedTo)

	info().Msg(color.Style{color.LightRed}.Render("disconnecting from conn nr. -> " + conv.IntToStr(wasConnectedTo)))

	// Call before disconnect
	_event("start", "OnBeforeDisconnect").Msg("")
	c.onBeforeDisconnect.Scan(func(k string, v interface{}) {
		v.(OnBeforeDisconnect)()
	})
	_event("finish", "OnBeforeDisconnect").Msg("")

	// Signal that we are closing the connection
	c.ctxDone.True()
	c.ctx.Cancel()

	// Wait for everyone to finish and then close the websocket connection properly!

	// Close the websocket connection
	_err := c.WSClient.Close()
	if _err != nil {
		_error().Err(_err).Msg("")
	}

	// Call all callbacks when disconnect finished
	_event("start", "OnDisconnect").Msg("")
	c.onDisconnect.Scan(func(k string, v interface{}) {
		v.(OnDisconnect)()
	})
	_event("finish", "OnDisconnect").Msg("")

	// We indicate for which specific connection disconnect occurred
	_event("start", "OnLinkDisconnect").Msg("")
	c.onLinkDisconnect.Scan(func(k string, v interface{}) {
		v.(OnLinkDisconnect)(wasConnectedTo, connectedTo)
	})
	_event("finish", "OnLinkDisconnect").Msg("")

	// If we have disconnected properly... now we should check if there is a reconnection process
	// If yes, then we should trigger/activate it!
}

// autoDisconnect -> it's being called when the disconnection process is done by other reasons
// The user cannot call this function, and it shouldn't!
// This function also includes the reconnection process!
func (c *Client) autoDisconnect() {
	info := func() *zerolog.Event {
		return c.LInfoF("autoDisconnect")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")
	info().Msg("calling Disconnect function")
	c.Disconnect()
	info().Msg("calling autoReconnect function")
	c.autoReconnect()
	info().Msg("leaving autoDisconnect function")
}

// DisconnectGracefully -> It will send the disconnect packet, but it will wait until it receives confirmation, and if
// it doesn't receives it, it will be refer to the timeout and other params configured on the client
// after everything has passed, it will disconnect
func (c *Client) DisconnectGracefully(code uint16, message string) {
	info := func() *zerolog.Event {
		return c.LInfoF("DisconnectGracefully")
	}
	info().Msg("entering...")
	defer info().Msg("leaving...")
	if code > 0 {
		c.closeCode = code
	}
	if message != "" && len(message) > 0 {
		c.closeMessage = message
	}
	// send through channel!
	c.writeChannel <- []byte(strconv.Itoa(msg.Close))
}
