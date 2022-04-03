package client

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/connection"
	"github.com/KyaXTeam/go-core/v2/core/helpers/slice"
)

// AddConnection -> adds additional Connection Config to current ones
/*func (c *Client) AddConnection(connectionNr int, connection *connection.Connection) {
	// TODO: check if everything is completed and ok!
	// TODO: LOCKS! -> we don't need locks.. becaues we will not add in realtime different connections, but who knows...

	// TODO: we should append to this slice! also add lock!
	// If item index exists... then replace it!, if not them append it,
	c.config.Connections[connectionNr] = connection
}*/

// RemoveConnection -> Removes a current/existing connection
/*func (c *Client) RemoveConnection(name string) {
	if _, ok := c.config.Connections[name]; ok {
		delete(c.config.Connections, name)
	}
}*/

func (c *Client) GetConnection(connectionNr int) *connection.Connection {
	if ifExists, _ := slice.IndexExists(c.config.Connections, connectionNr); ifExists {
		return c.config.Connections[connectionNr]
	}
	return nil
}

func (c *Client) GetConnections() []*connection.Connection {
	// Returns the pointer of the map
	return c.config.Connections
}
