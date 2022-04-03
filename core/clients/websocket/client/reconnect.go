package client

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/connection"
	"github.com/KyaXTeam/go-core/v2/core/helpers/conv"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"time"
)

/*
1. it should be added when a disconnect has happened
2. it should be enabled/disabled
3. it should trigger a function
4. it should have a specific nr of reconnects or unlimited....
5.
*/

// Reconnect -> it's initiating again the connection process but calling and other callbacks before it...
// The user can also call this function if handling manually or from other place this process!
// We have created this function to delimit the connect with reconnect as events, and for the user to know
// what's happened, and of course for easier understanding!
func (c *Client) Reconnect() {
	info := func() *zerolog.Event {
		return c.LInfoF("Reconnect")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("Reconnect")
	}
	_event := func(eventType string, eventName string) *zerolog.Event {
		return c.LEventF(eventType, eventName, "Reconnect")
	}

	info().Msg("entering...")
	defer info().Msg("leaving...")

	// Set that reconnect has being called!
	// Call the Reconnecting event
	// If success then call the reconnected event

	// TODO: should we do here a check that we already reconnect, and to
	// not reconnect multiple times?!!...

	// Set that reconnect function has being called!
	if c.isReconnecting.IfFalseSetTrue() {
		warn().Msg("already called...")
		return
	}
	// Increment the nr of reconnections..
	c.reconnectRound.Inc(1)

	// Calling the on reconnecting callbacks

	_event("start", "OnReconnecting").Msg("")
	c.onReconnecting.Scan(func(k string, v interface{}) {
		v.(OnReconnecting)()
	})
	_event("finish", "OnReconnecting").Msg("")

	//

	// Add to event listener -> even if added multiple times, it will be overwritten
	c.OnConnectSuccess("autoReconnect", func(connNr int, connection *connection.Connection) {
		info().Msg(color.LightGreen.Render("reconnected successfully"))

		// Here we have connected successfully

		// If Reconnecting wasn't called then return...
		// Because this event may remain still registered!
		if !c.isReconnecting.Get() {
			return
		}
		c.isReconnecting.False()
		// Is Reconnecting ON!
		// Reset the nr of retries
		c.reconnectRound.Set(0)
		// Set that we are not reconnecting anymore
		_event("start", "OnReconnected").Msg("")
		c.onReconnected.Scan(func(k string, v interface{}) {
			v.(OnReconnected)(connNr, connection)
		})
		_event("finish", "OnReconnected").Msg("")
		info().Msg("reconnection finished")
		/*if c.HasOnConnectSuccess("autoReconnect") {
			c.RemoveOnConnectSuccess("autoReconnect")
		}*/
	})

	//

	// If all connections have failed...
	c.OnConnectFailedAll("autoReconnect", func() {
		warn().Msg(color.LightRed.Render("reconnection failed..."))

		// No connections have succeeded
		// If Reconnecting wasn't called then return...
		// Because this event may remain still registered!
		if !c.isReconnecting.Get() {
			warn().Msg("reconnect not called...leaving...")
			return
		}
		// Set that we are not reconnecting anymore
		c.isReconnecting.False()
		info().Msg(color.Style{color.Yellow}.Render("retrying reconnection..."))
		go c.autoReconnect()
	})

	//

	// Calling the connect function
	info().Msg("calling connect function")
	c.connect()
}

// autoReconnect -> calls Reconnect after a specific time and by checking the nr. of tried reconnections
func (c *Client) autoReconnect() {
	info := func() *zerolog.Event {
		return c.LInfoF("autoReconnect")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("autoReconnect")
	}

	info().Msg("entering...")
	defer info().Msg("leaving...")
	// Launch the reconnect after a specific time...
	// Check max nr of times we have set to reconnect

	reconnect := conv.ParseBool(c.config.AutoReconnect)
	if reconnect {
		info().
			Int16("reconnect_max_retries", c.config.Reconnect.MaxRetries).
			Msg("auto reconnect is enabled")
		// If it's -1, it means infinite times!
		if c.config.Reconnect.MaxRetries != -1 {
			// Check the current rounds...
			info().
				Uint16("reconnect_nr_of_retries", c.reconnectRound.Get()).
				Msg("checking how many times have retried reconnection")
			if uint16(c.config.Reconnect.MaxRetries) <= c.reconnectRound.Get() {
				// Stop auto Reconnect
				// Reset the counter of reconnects
				c.reconnectRound.Set(0)
				reconnect = false

				// Call an event that reconnect has being stopped completely!
				c.LEvent("start", "OnReconnectFailed", nil)
				c.onReconnectFailed.Scan(func(k string, v interface{}) {
					v.(OnReconnectFailed)()
				})
				c.LEvent("finish", "OnReconnectFailed", nil)
			}
		} else {
			info().
				Uint16("reconnect_nr_of_retries", c.reconnectRound.Get()).
				Msg("there are infinite reconnect retries")
		}
	} else {
		warn().Msg("auto reconnect is disabled")
	}

	if reconnect {
		warn().Uint16("reconnect_timeout", c.config.Reconnect.TimeoutSeconds).
			Msg("calling reconnect function after a specific timeout")
		time.AfterFunc(time.Second*time.Duration(c.config.Reconnect.TimeoutSeconds), func() {
			warn().Msg("checking if AutoReconnect enabled")
			if conv.ParseBool(c.config.AutoReconnect) {
				warn().Msg("calling Reconnect")
				c.Reconnect()
			}
		})
	}
}
