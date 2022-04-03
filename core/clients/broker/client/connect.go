package client

func (c *Client) Connect() {
	// We can create the layer for the broker only if we will add additional functionality...
	// In that case, we should rename the function to Start
	// In that case multiple goroutines will be running....
	// But in this case only websocket will be running... so it's not necessary to create additional functionality here

	c.WSClient.Connect()
}
