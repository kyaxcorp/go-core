package client

import "github.com/kyaxcorp/go-core/core/listeners/websocket/server/msg"

func (c *Client) WriteText(message string) *Client {
	if message == "" {
		return c
	}
	go func() {
		c.writeChannel <- msg.TextToBytes(message)
	}()
	return c
}

func (c *Client) SendText(message string) *Client {
	return c.WriteText(message)
}

// WriteJSON - It sends Any structure to the client encoded as JSON!
func (c *Client) WriteJSON(message interface{}, onJsonError OnJsonError) *Client {
	go func() {
		encoded, err := msg.JsonToBytes(message)
		if err != nil {
			if onJsonError != nil {
				onJsonError(err, message)
			}
			return
		}
		c.writeChannel <- encoded
	}()
	return c
}
func (c *Client) SendJSON(message interface{}, onJsonError OnJsonError) *Client {
	return c.WriteJSON(message, onJsonError)
}

func (c *Client) WriteBinary() {

}

func (c *Client) SendBinary() {

}
