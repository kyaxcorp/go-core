package client

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/config"
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/connection"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_int"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_map_string_interface"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_time"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_uint16"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_uint64"
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/server"
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	"github.com/gorilla/websocket"
	"net/http"
)

/*
Should be created 2 way of sending/receive messages in an async way!
1 routine for reading messages!
1 routine for sending messages

- Communication with these routines should be made through channels
- Should be a mechanism of auto reconnect
- Should check if ping/pong works...!?
- Should we handle or buffer incoming requests?! and store them for later when reconnect will be done!?
-

*/

// These client can connect to multiple destinations if one server is not responding!

// OnReceive -> when we receive any message (text or binary)
type OnReceive func(recvMessage *server.ReceivedMessage, c *Client)
type OnConnect func()

// OnText -> when we receive a text message
type OnText func(recvMessage *server.ReceivedMessage, c *Client)

// OnBinary -> when we receive a binary message
type OnBinary func(recvMessage *server.ReceivedMessage, c *Client)

type OnSend func()
type OnSendError func()
type OnReadError func(error)
type OnError func()
type OnBeforeDisconnect func()

// OnLinkDisconnect -> When we have lost connection for a specific connection
type OnLinkDisconnect func(connNr int, connection *connection.Connection)

type OnDisconnect func()

// OnReconnecting -> it's being called when reconnect function it's being called!
type OnReconnecting func()

// OnReconnected -> it's being called when reconnection has being called, and the client has successfully connected to
// the host
type OnReconnected func(connNr int, connection *connection.Connection)

// OnReconnectFailed -> When reconnection has totally failed and it has stopped...
// Usually if reconnect max retries has being set to -1 meaning infinite... this event will not occur!
type OnReconnectFailed func()

// OnBeforeConnectToServer -> it's being called before each connection call
type OnBeforeConnectToServer func()

// OnConnectError -> it's being called when connection has failed to be initiated
type OnConnectError func(error)

// OnConnectFailed -> It's being called each time when a connection has failed to connect
type OnConnectFailed func(OnConnectErrorResponse)

// OnConnectFailedAll -> it's being called When all connections have failed to connect!
type OnConnectFailedAll func()

// OnConnectSuccess -> it's being called when the client has being connected successfully to the host
type OnConnectSuccess func(connNr int, connection *connection.Connection)

// OnTerminate -> it's being called when connection has being totally terminated or finalized
type OnTerminate func()

// OnStopConnectingFinish -> it's being used when stopConnecting function it's being called!
// And when the process of connecting is stopped and finished, this event it's being called
type OnStopConnectingFinish func()

// OnConnectErrorResponse -> it's a response structure for OnConnectFailed
type OnConnectErrorResponse struct {
	// Other details
	MaxRetries uint16 // how many retries will be
	RetryNr    uint16 // the attempt
	// Error Details
	MainError     error
	OriginalError error
	// Connection Details
	ConnectionNr uint16
	Connection   *connection.Connection
}

type Client struct {
	config config.Config // This is the configuration
	// RequestHeader -> can be configured directly
	RequestHeader http.Header

	connectStartTime    *_time.Time // When the connect command has being called
	connectEndTime      *_time.Time // When the process has finished successfully or with a failure
	disconnectStartTime *_time.Time // When the disconnect command has being called
	disconnectEndTime   *_time.Time // When disconnect process has finished successfully or with failure

	// This is the outbound messages
	writeChannel chan []byte
	//connections       map[string]*connection.Connection // These are the connections where the client can connect!
	failedConnections map[string]*connection.Connection // These are the failed connections to which couldn't  connect

	// This is the Logger of the websocket client, we make it public, for easier interaction!
	Logger *model.Logger

	// Properties of events
	connectedTo  *_int.Int   // Connection Name (where it's connected now)
	connectingTo *_int.Int   // Connection Name (where it's connecting right now)
	isConnected  *_bool.Bool // If it's connected!
	isConnecting *_bool.Bool // If it's connecting

	isDisconnecting *_bool.Bool // If it's disconnecting right now

	isReconnecting *_bool.Bool     // If it's Reconnecting
	reconnectRound *_uint16.Uint16 // What time(round) it's connecting...

	// That's when we are disconnecting/closing gracefully!
	closeCode    uint16
	closeMessage string

	// Statistics
	// Sent
	nrOfSentMessages        *_uint64.Uint64
	nrOfSentTextMessages    *_uint64.Uint64
	nrOfSentBinaryMessages  *_uint64.Uint64
	nrOfSentFailedMessages  *_uint64.Uint64
	nrOfSentSuccessMessages *_uint64.Uint64
	sentBytes               *_uint64.Uint64
	sentBinaryBytes         *_uint64.Uint64
	sentTextBytes           *_uint64.Uint64

	// Received
	nrOfReceivedMessages       *_uint64.Uint64
	nrOfReceivedTextMessages   *_uint64.Uint64
	nrOfReceivedBinaryMessages *_uint64.Uint64
	receivedBytes              *_uint64.Uint64
	receivedBinaryBytes        *_uint64.Uint64
	receivedTextBytes          *_uint64.Uint64

	// Other Statistics
	nrOfDisconnections *_uint64.Uint64

	// Events
	onReceive               *_map_string_interface.MapStringInterface // When it has received a message
	onText                  *_map_string_interface.MapStringInterface // When we receive a Text Message
	onBinary                *_map_string_interface.MapStringInterface // When we receive a Binary Message
	onSend                  *_map_string_interface.MapStringInterface // When it sends something (before sending)
	onSendError             *_map_string_interface.MapStringInterface // When tried to send, but received an error from the Writer
	onReadError             *_map_string_interface.MapStringInterface // When tried to read, but received an error from the Reader
	onError                 *_map_string_interface.MapStringInterface // Generally, if an error occurred anywhere!
	onBeforeDisconnect      *_map_string_interface.MapStringInterface // When the client calls Disconnect function manually!
	onDisconnect            *_map_string_interface.MapStringInterface // When the client it's disconnected
	onLinkDisconnect        *_map_string_interface.MapStringInterface // When someone (server/client) or even link has being disconnected! No other function has being called to call the disconnect!
	onConnect               *_map_string_interface.MapStringInterface // When the client has connected
	onReconnecting          *_map_string_interface.MapStringInterface // When the clients was connected before, but it has AutoReconnect -> True and it has reconnected to the host!
	onReconnected           *_map_string_interface.MapStringInterface // When the clients was connected before, but it has AutoReconnect -> True and it has reconnected to the host!
	onReconnectFailed       *_map_string_interface.MapStringInterface // When it has failed totally when no more retries
	onBeforeConnectToServer *_map_string_interface.MapStringInterface
	onConnectError          *_map_string_interface.MapStringInterface // When there is an error inside the connect function, some conditions are not met...
	onConnectFailed         *_map_string_interface.MapStringInterface // When tried to connect to the server, but it failed
	onConnectFailedAll      *_map_string_interface.MapStringInterface // When tried to connect to all servers, and all of them failed!
	onConnectSuccess        *_map_string_interface.MapStringInterface // When connected to any server with success
	onTerminate             *_map_string_interface.MapStringInterface // When the app it's being closed, or close it's being called and it should terminate!
	onStopConnectingFinish  *_map_string_interface.MapStringInterface // On Stop connecting it's being called, and the process of connecting has finished

	// This is for stopping entirely the client... meaning to disconnect, stop the connection process etc...!
	parentCtx context.Context
	ctx       *_context.CancelCtx
	// THis is for canceling the connection process
	ctxConnect *_context.CancelCtx
	// Here we just set if it's done or not!
	ctxDone *_bool.Bool

	//Objects
	WSClient     *websocket.Conn // This is the Gorilla Web Socket Client
	HttpResponse *http.Response  // This is the response from the server!
}
