package client

import (
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/listeners/websocket/server"
	"github.com/kyaxcorp/go-core/core/listeners/websocket/server/msg"
	"github.com/rs/zerolog"
)

func (c *Client) runReader() {
	info := func() *zerolog.Event {
		return c.LInfoF("runReader")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("runReader")
	}
	_debug := func() *zerolog.Event {
		return c.LDebugF("runReader")
	}
	_error := func() *zerolog.Event {
		return c.LErrorF("runReader")
	}
	_event := func(eventType string, eventName string) *zerolog.Event {
		return c.LEventF(eventType, eventName, "runReader")
	}

	info().Msg("running...")
	defer info().Msg("leaving...")
	// receive message
	// The reader should not handle pings from server, it's already done by another handle
	// in the gorilla library
	terminate := false
	for {
		select {
		case <-c.ctx.Done():
			// We set it for all, because, some of them, may not be in the right moment of receiving it
			c.ctxDone.True()
			warn().Msg("terminating...")
			// TODO: stop Reader!
			terminate = true
		default:
			messageType, message, _err := c.WSClient.ReadMessage()
			if _err != nil {
				// handle error -> onReadError
				_error().Err(_err).Msg(color.Style{color.LightRed}.Render("failed to read message"))

				if !c.ctxDone.Get() {
					_event("start", "OnReadError").Msg("")
					c.onReadError.Scan(func(k string, v interface{}) {
						v.(OnReadError)(_err)
					})
					_event("finish", "OnReadError").Msg("")

					// Calling Disconnect
					info().Msg("calling autoDisconnect")
					go c.autoDisconnect()
				}

				terminate = true
				break
			}

			// TODO: should we capture PING messages and respond to them?!

			messageLength := uint(len(message))
			// Save for statistics
			c.receivedBytes.Inc(uint64(messageLength))
			// Save nr of received messages
			c.nrOfReceivedMessages.Inc(1)

			_debug().
				Int("message_type", messageType).
				Uint("message_length_bytes", messageLength).
				Msg("message received")

			// Without a pointer
			receivedMessage := &server.ReceivedMessage{
				MessageType:   int8(messageType),
				MessageLength: messageLength,
				Message:       message,
			}

			// Call events by message type
			// Be careful when using these...
			// They are not async, they can degrade performance?!
			switch messageType {
			case msg.Text:
				// Save for statistics
				c.nrOfReceivedTextMessages.Inc(1)
				c.receivedTextBytes.Inc(uint64(messageLength))
				// Call the events
				c.onText.Scan(func(k string, v interface{}) {
					v.(OnText)(receivedMessage, c)
				})
			case msg.Binary:
				// Save for statistics
				c.nrOfReceivedBinaryMessages.Inc(1)
				c.receivedBinaryBytes.Inc(uint64(messageLength))
				// Call the events
				c.onBinary.Scan(func(k string, v interface{}) {
					v.(OnBinary)(receivedMessage, c)
				})
			case msg.Close:
				_debug().
					Msg(color.Style{color.LightYellow}.Render("close code has being received"))
				// TODO: what we do next?!...
			}

			// TODO: maybe later add support for async
			c.onReceive.Scan(func(k string, v interface{}) {
				v.(OnReceive)(receivedMessage, c)
			})
		}
		if terminate {
			break
		}
	}
}
