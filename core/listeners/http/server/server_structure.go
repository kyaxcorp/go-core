package server

import (
	"context"
	"github.com/gin-gonic/gin"
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

type OnJsonError func(err error, message interface{})

type OnEvent func(...interface{})

// On Connect it will be launched in a goroutine!
type OnRequest func(c *Client, s *Server)
type OnResponse func(c *Client, s *Server)

// Stop
type OnStop func(s *Server)
type OnBeforeStop func(s *Server)
type OnStopped func(s *Server)

// Start
type OnStart func(s *Server)
type OnBeforeStart func(s *Server)
type OnStarted func(s *Server)

type Server struct {
	Name        string
	Description string

	connectionID *_uint64.Uint64
	// Starting time of the server
	startTime *_time.Time
	// Stop time of the server
	stopTime *_time.Time

	isStopCalled  *_bool.Bool
	isStopped     *_bool.Bool
	isStartCalled *_bool.Bool
	isStarted     *_bool.Bool

	genConnIDLock sync.Mutex

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

	HttpServer *gin.Engine

	// Events/Callbacks - they are common for all routes!
	onRequest  *_map_string_interface.MapStringInterface
	onResponse *_map_string_interface.MapStringInterface

	// Stop
	onStop       *_map_string_interface.MapStringInterface
	onBeforeStop *_map_string_interface.MapStringInterface
	onStopped    *_map_string_interface.MapStringInterface

	// Start
	onStart       *_map_string_interface.MapStringInterface
	onBeforeStart *_map_string_interface.MapStringInterface
	onStarted     *_map_string_interface.MapStringInterface

	// Here we store the active/registered ClientsStatus (Connections)
	c *clientsData
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

type Client struct {
	// Logger -> it's specifically related to client
	// Logs will be written to client file, but not in the main websocket log file
	// If needed, this can be enabled
	Logger *model.Logger

	// connectTime -> when it has being connected , and it's read only...we don't change it later
	connectTime time.Time

	// connectionID -> Generated server connection id, it's read only!
	connectionID uint64

	// Auth Details containing (User Details, Device Details, Authentication Details)
	authDetails *authentication.AuthDetails
	connDetails *connection.ConnDetails

	// Gin Context
	httpContext *gin.Context

	// This is the server itself as a relation!
	server *Server

	writeTicker *time.Ticker

	// Buffered channel of outbound messages.
	send chan []byte

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

	/*// Here we store on response callbacks!
	payloadMessageCallbacks    map[string]TextPayloadOnResponse
	payloadMessageCallbackLock sync.Mutex*/

	randomPayloadID *_uint16.Uint16

	// This is Custom data array which can be accessed with Get/Set Methods
	//customData map[string]interface{}
	customData *_map_string_interface.MapStringInterface
}
