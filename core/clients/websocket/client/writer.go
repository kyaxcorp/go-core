package client

import (
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/server/msg"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"strconv"
)

func (c *Client) runWriter() {
	info := func() *zerolog.Event {
		return c.LInfoF("runWriter")
	}
	/*warn := func() *zerolog.Event {
		return c.LWarnF("runWriter")
	}*/
	_debug := func() *zerolog.Event {
		return c.LDebugF("runWriter")
	}
	_error := func() *zerolog.Event {
		return c.LErrorF("runWriter")
	}

	info().Msg("running...")
	defer info().Msg("leaving...")

	terminate := false
	for {
		select {
		case message, ok := <-c.writeChannel:
			// This is an Internal Client Error...
			// TODO: we should see how we gonna handle this type of error! It means the channel has being closed
			// But channel can be closed even before this or in concurrently when disconnect happends?!...
			if !ok {
				_debug().Msg("sending close message, for graceful disconnect (1)")

				// TODO: we should handle the close of the entire connection! we should call the cancel function!? or close function!
				_err := c.WSClient.WriteMessage(msg.Close, []byte{}) // closing the connection gracefully!
				if _err != nil {
					_error().Err(_err).Msg("failed to send close message flag")
				}
				// Calling disconnect!
				go c.autoDisconnect()
				return
			}

			messageType, _err := strconv.Atoi(string(message[0]))
			messageLength := len(message[1:]) // In bytes...

			_debug().
				Int("message_type", messageType).
				Int("message_length_bytes", messageLength).
				Msg("sending message")

			// If Close message has being sent!
			switch messageType {
			case msg.Close:
				_debug().
					Msg("sending close message, for graceful disconnect (2)")

				// https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent
				// Sending message that we are closing the connection
				closeCode := int(c.closeCode)
				closeMsg := c.closeMessage
				if c.closeCode == 0 {
					closeCode = 1001
				}
				if closeMsg == "" {
					closeMsg = "ws client is disconnecting gracefully"
				}

				if _err = c.WSClient.WriteMessage(
					msg.Close,
					websocket.FormatCloseMessage(closeCode, closeMsg),
				); _err != nil {
					_error().Err(_err).Msg("failed to send close message flag")
					return
				}
				continue
			}

			if _err != nil {
				_error().Err(_err).Msg("incorrect message type")
				//c.server.L.Infoln("Write Pump - incorrect Message type")
				return
			}

			// Choose a Writer by type! (we have defined that first byte is the Message Type!)
			w, _err := c.WSClient.NextWriter(messageType)
			if _err != nil {
				_error().Err(_err).Msg("failed to get next writer")
				return
			}
			// log.Println("Writing the Message")

			// Writing message
			// We exclude the first byte because it includes the messageType
			_, _err = w.Write(message[1:])
			if _err != nil {
				_error().Err(_err).Msg("failed to write message")
				// Failed to send/write message
				c.nrOfSentFailedMessages.Inc(1)
			} else {
				// Message Sent
				// Saving for statistics
				c.nrOfSentMessages.Inc(1)
				// Count
				c.nrOfSentSuccessMessages.Inc(1)
				// Count Sent Bytes
				c.sentBytes.Inc(uint64(messageLength))

				switch messageType {
				case msg.Binary:
					// Saving for statistics
					c.nrOfSentBinaryMessages.Inc(1)
					c.sentBinaryBytes.Inc(uint64(messageLength))
				case msg.Text:
					// Saving for statistics
					c.nrOfSentTextMessages.Inc(1)
					c.sentTextBytes.Inc(uint64(messageLength))
				}
			}

			// Closing the writer
			if _err = w.Close(); _err != nil {
				_error().Err(_err).Msg("failed to close writer")
				// Failed to close the writer
				return
			}
		case <-c.ctx.Done():
			// TODO: stop Writer!
			c.ctxDone.True()
			_debug().Msg("sending close message, for graceful disconnect (3)")

			terminate = true

			if _err := c.WSClient.WriteMessage(
				msg.Close,
				//websocket.FormatCloseMessage(int(c.closeCode), c.closeMessage),
				// https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent/code
				websocket.FormatCloseMessage(1001, "ws client is shutting down..."),
			); _err != nil {
				_error().Err(_err).Msg("failed -> sending close message, for graceful disconnect (4)")
				return
			}
			// TODO: send a message of GRACEFUL DISCONNECT!
		}
		if terminate {
			break
		}
	}
}
