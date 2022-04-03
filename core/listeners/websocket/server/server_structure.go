package server

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_map_string_interface"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_time"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_uint16"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_uint64"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/duration"
	"github.com/KyaXTeam/go-core/v2/core/listeners/http/middlewares/authentication"
	"github.com/KyaXTeam/go-core/v2/core/listeners/http/middlewares/connection"
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/config"
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type OnJsonError func(err error, message interface{})

type OnEvent func(...interface{})

type OnBeforeUpgrade func(s *Server)
type OnUpgrade func(c *Client, s *Server)

// On Connect it will be launched in a goroutine!
type OnConnect func(c *Client, s *Server)
type OnMessage func(message *ReceivedMessage, c *Client, s *Server)
type OnClose func(c *Client, s *Server)

// Stop
type OnStop func(s *Server)
type OnBeforeStop func(s *Server)
type OnStopped func(s *Server)

// Start
type OnStart func(s *Server)
type OnBeforeStart func(s *Server)
type OnStarted func(s *Server)

/*type Client struct {
	Message   string
	ginClient *gin.Context
}*/

type Server struct {
	Name        string
	Description string

	// This is the original config from the constructor
	config config.Config

	connectionID *_uint64.Uint64

	isStopCalled  *_bool.Bool
	isStopped     *_bool.Bool
	isStartCalled *_bool.Bool
	isStarted     *_bool.Bool

	// Starting time of the server
	startTime *_time.Time
	// Stop time of the server
	stopTime *_time.Time

	// This will be the main folder where we will store the logs
	LoggerDirPath string
	// This is the logger configuration!
	Logger *model.Logger

	// Enables Server Status through HTTP
	enableServerStatus *_bool.Bool
	// These are the server status credentials
	statusUsername string
	statusPassword string

	// enableUnsecure -> most of the time is readonly!
	enableUnsecure bool // Enable unsecure connections
	//
	enableSSL   bool
	sslCertPath string
	sslKeyPath  string

	// It also includes port
	ListeningAddresses    []string // This is for unencrypted
	ListeningAddressesSSL []string // This is for encrypted
	// Context
	parentCtx context.Context
	ctx       *_context.CancelCtx

	// This is GIN Server
	WSServer *gin.Engine
	// This is the registrationHub that registers the c...
	WSRegistrationHub *RegistrationHub
	// This is the registrationHub which sends broadcast messages!
	WSBroadcastHub *BroadcastHub
	// This is the upgrader which transfers from http to WebSocket
	WSUpgrader websocket.Upgrader

	// These settings are being set into the Upgrader
	readBufferSize    *_uint64.Uint64
	writeBufferSize   *_uint64.Uint64
	enableCompression *_bool.Bool

	// It enables automatic upgrade from http to ws
	// Usually the server will be behind a proxy server, so it's not necessary....
	EnableHttpToWSUpgrade *_bool.Bool

	// Events/Callbacks - they are common for all routes!
	onBeforeUpgrade *_map_string_interface.MapStringInterface
	onConnect       *_map_string_interface.MapStringInterface
	onMessage       *_map_string_interface.MapStringInterface
	onClose         *_map_string_interface.MapStringInterface

	// Stop
	onStop       *_map_string_interface.MapStringInterface
	onBeforeStop *_map_string_interface.MapStringInterface
	onStopped    *_map_string_interface.MapStringInterface

	// Start
	onStart       *_map_string_interface.MapStringInterface
	onBeforeStart *_map_string_interface.MapStringInterface
	onStarted     *_map_string_interface.MapStringInterface

	// ------Settings ---------\\
	// Time allowed to write a Message to the peser.
	//WriteWait time.Duration
	WriteWait *duration.Duration
	// Time allowed to read the next pong Message from the peer.
	PongWait *duration.Duration
	// send pings to peer with this period. Must be less than pongWait.
	PingPeriod *duration.Duration
	// Maximum Message size allowed from peer.
	MaxMessageSize *_uint64.Uint64

	// ------Settings ---------\\

	// Here we store the active/registered Clients (Connections)
	// c      map[*Client]bool
	// clientsIndex ClientsIndex
	c *clientsData

	// Here we store the created Hubs
	// The Hubs are being added here when they are started only!
	// There is no reference until they are started!
	Hubs map[*Hub]bool
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

func (s *Server) genConnID() uint64 {
	return s.connectionID.ResetMaxAndInc(1, 1000000, 0)
}

type ReceivedMessage struct {
	// There are few types...
	MessageType int8
	// Bytes?!!
	MessageLength uint
	Message       []byte
}

type Client struct {
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
	httpContext *gin.Context

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
