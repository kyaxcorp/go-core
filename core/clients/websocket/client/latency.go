package client

func (c *Client) GetLatency() {
	// Start Time
	// Send a payload id
	// Wait for response
	// End time
	// Calculate the timing...
	// It's important to use Before Send and after Send...
	// Also the buffer may be busy... so
	// the best way to find out latency is it to use a single connection without message overload!
	// In this way it will work perfectly
}
