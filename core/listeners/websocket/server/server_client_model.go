package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_bool"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_map_string_interface"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_time"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_uint16"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_uint64"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/authentication"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/connection"
	"github.com/kyaxcorp/go-core/core/logger/model"
	"sync"
	"time"
)

type Client struct {
	// parentCtx -> e WS Server-ul
	parentCtx context.Context
	// ctx asta e Client-ul!
	ctx *_context.CancelCtx

	// Logger -> it's specifically related to client
	// Logs will be written to client file, but not in the main websocket log file
	// If needed, this can be enabled
	Logger *model.Logger

	// connectTime -> when it has being connected , and it's read only...we don't change it later
	connectTime time.Time

	// connectionID -> Generated server connection id, it's read only!
	connectionID uint64

	// Ping information
	// Ping Send
	lastSendPingTry     *_time.Time     // What was the last time it tried to send a ping
	lastSentPingTime    *_time.Time     // what was the last successful time to send a ping
	nrOfSentPings       *_uint64.Uint64 // nr of successful pings!
	nrOfFailedSendPings *_uint64.Uint64 // nr of failed pings
	// Pong Receive

	lastSentPongTime     *_time.Time
	lastReceivedPongTime *_time.Time
	nrOfReceivedPongs    *_uint64.Uint64
	nrOfSentPongs        *_uint64.Uint64
	nrOfFailedSendPongs  *_uint64.Uint64

	// Auth Details containing (User Details, Device Details, Authentication Details)
	authDetails *authentication.AuthDetails
	connDetails *connection.ConnDetails

	// Gin Context
	httpContext     *gin.Context
	safeHttpContext *gin.Context

	//registrationHub *Hub

	// Client Specific on Message
	onMessage OnMessage

	// This is the server itself as a relation!
	server *Server

	registrationHub *RegistrationHub
	broadcastHub    *BroadcastHub

	// The websocket connection.
	conn *websocket.Conn

	pingTicker *time.Ticker

	// Buffered channel of outbound messages.
	send chan []byte

	// This is the channel where the WritePump
	closeWritePump chan bool

	// It shows if the connection is closed!
	isClosed *_bool.Bool

	// In case of Close call we define the code and reason!
	// closeCode -> it's mostly read only! it's used only once on graceful disconnect
	closeCode uint16
	// closeMessage -> it's mostly read only! it's used only once on graceful disconnect
	closeMessage string

	// If someone has called disconnect function!
	isDisconnecting *_bool.Bool

	// Message ID - is the nr. of messages sent to the client!
	nrOfSentMessages        *_uint64.Uint64
	nrOfSentFailedMessages  *_uint64.Uint64
	nrOfSentSuccessMessages *_uint64.Uint64

	// Here we store on response callbacks!
	payloadMessageCallbacks    map[string]TextPayloadOnResponse
	payloadMessageCallbackLock sync.Mutex

	randomPayloadID *_uint16.Uint16

	// This is Custom data array which can be accessed with Get/Set Methods
	//customData map[string]interface{}
	customData *_map_string_interface.MapStringInterface
}

// Here we store reverse map of the connections!
type ClientsIndex struct {
	// TODO: see later maybe we will use sync.Map for better sync... that's only if register/unregister will perform multiple
	// Goroutines at once!

	// These are locks for reading/writing to/form indexes
	usersLock       sync.RWMutex
	devicesLock     sync.RWMutex
	connectionsLock sync.RWMutex
	authTokensLock  sync.RWMutex
	ipAddressesLock sync.RWMutex
	requestPathLock sync.RWMutex

	// Indexes
	Users       map[string]map[uint64]*Client
	Devices     map[string]map[uint64]*Client
	Connections map[uint64]*Client
	AuthTokens  map[string]map[uint64]*Client
	IPAddresses map[string]map[uint64]*Client
	RequestPath map[string]map[uint64]*Client
}
