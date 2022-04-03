package client

// GetConnectedTo -> Where has being connected to
func (c *Client) GetConnectedTo() int {
	return c.connectedTo.Get()
}

// GetConnectingTo -> Where it's trying to connect to!
func (c *Client) GetConnectingTo() int {
	return c.connectedTo.Get()
}

// IsConnected ->  If is connected right now!
func (c *Client) IsConnected() bool {
	return c.isConnected.Get()
}

// IsConnecting ->  If is connecting right now!
func (c *Client) IsConnecting() bool {
	return c.isConnecting.Get()
}
