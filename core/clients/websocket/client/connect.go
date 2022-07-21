package client

import (
	"crypto/tls"
	"github.com/gookit/color"
	"github.com/gorilla/websocket"
	"github.com/kyaxcorp/go-core/core/clients/websocket/connection"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/base64"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/authentication"
	"github.com/rs/zerolog"
	"time"
)

// StopConnecting -> Simply stop connecting
func (c *Client) StopConnecting() {
	// TODO: check if it is connecting
	c.ctxConnect.Cancel()
}

// Connect -> connect to the host
func (c *Client) Connect() {
	warn := func() *zerolog.Event {
		return c.LWarnF("connect")
	}
	if c.isConnecting.Get() {
		// Is already connecting
		warn().Msg("already connecting...")
		return
	}
	go c.connect()
}

// connect -> error it's being returned only if the user wants to handle it... it's not mandatory...!
func (c *Client) connect() error {
	info := func() *zerolog.Event {
		return c.LInfoF("connect")
	}
	warn := func() *zerolog.Event {
		return c.LWarnF("connect")
	}
	_debug := func() *zerolog.Event {
		return c.LDebugF("connect")
	}
	_error := func() *zerolog.Event {
		return c.LErrorF("connect")
	}
	_event := func(eventType string, eventName string) *zerolog.Event {
		return c.LEventF(eventType, eventName, "connect")
	}

	info().Msg("entering...")
	defer info().Msg("leaving...")

	/*
		1. Check if there are any connections
		2. Check if multiple connections are allowed
		3. Check the current connection where it has stopped
		4. Count -> how many times the client tried to connection
		5. Remember when it has connected
		6. It can loop through all connections and try connecting!
		7. If it has failed on a connection, go to next one...
		8. There is a retry timeout!
	*/
	//If it's already connecting... return...

	if c.isConnected.Get() {
		warn().Msg("already connected...")
		return define.Err(0, "already connected...")
	}

	if c.isConnecting.IfFalseSetTrue() {
		// Is already connecting
		warn().Msg("already connecting...")
		return define.Err(0, "already connecting...")
	}
	info().Msg("connecting...")

	// Set Connect Start Time!
	c.connectStartTime.SetNow()
	// Set is connected as False!
	c.isConnected.False()
	c.ctxDone.False()

	// finalize as failed!
	terminateConn := func() {
		// We set the values before the event, because events can contain errors...
		c.isConnected.False()
		c.isConnecting.False()
		c.connectEndTime.SetNow()

		// Stop Reconnection values...
		c.isReconnecting.False()
		c.reconnectRound.Set(0)

		_event("start", "OnTerminate").Msg("")
		c.onTerminate.Scan(func(k string, v interface{}) {
			v.(OnTerminate)()
		})
		_event("finish", "OnTerminate").Msg("")
	}

	// Create the main context
	c.ctx = _context.WithCancel(c.parentCtx)

	// Create Cancel Context!
	c.ctxConnect = _context.WithCancel(c.ctx.Context())
	defer func() {
		// Destroy connect context... we don't need it anymore...
		c.ctxConnect = nil
	}()

	// c.Logger.Logger.Info().Interface("connections", c.connections)

	// Get all configures connections
	availableConn := len(c.config.Connections)
	// If there are no connections -> raise error and return!
	if availableConn == 0 {
		_error().Msg("no available connections...")
		/*
			1. Sa fac write singur in fisierul cu erori...
				- insa in cazul dat eroarea la sentry nu va fi
			2. Sa chem functia de la ErrorLogger
		*/

		_event("start", "OnConnectError").Msg("")
		c.onConnectError.Scan(func(k string, v interface{}) {
			v.(OnConnectError)(NoConnectionsDefined)
		})
		_event("finish", "OnConnectError").Msg("")
		terminateConn()
		return define.Err(0, "no available connections...")
	}

	// The iteration will be random! because of the map!

	// Set initial values
	isConnected := false // If it has connected
	isLastConn := false  // If it's the last connection from the defined ones

	// When stop connecting has being called
	stopConnectingCalled := false

	// If we need to terminate connection
	terminate := false
	usedConnections := 0 // Through how many connections we have iterated already!
	// Iterated through defined connections!

	logConnKey := "conn_nr"
	for connNr, conn := range c.config.Connections {
		infoConn := func() *zerolog.Event {
			return info().Int(logConnKey, connNr)
		}
		warnConn := func() *zerolog.Event {
			return warn().Int(logConnKey, connNr)
		}
		debugConn := func() *zerolog.Event {
			return _debug().Int(logConnKey, connNr)
		}
		errorConn := func() *zerolog.Event {
			return _error().Int(logConnKey, connNr)
		}
		eventConn := func(eventType string, eventName string) *zerolog.Event {
			return _event(eventType, eventName).Int(logConnKey, connNr)
		}
		infoConn().Msg("using connection nr -> " + conv.IntToStr(connNr))

		// Check if multiple connections is allowed to use!
		// If it's not allowed, then we will use only first available and then break even if there are multiple defined
		if !conv.ParseBool(c.config.UseMultipleConnections) && usedConnections == 1 {
			// It will use only 1 connection... and if it failed previously... then it will stop here!
			warnConn().Msg("multiple connections are not allowed, breaking")
			break
		}
		// Set to which connection we are connecting right now!
		c.connectingTo.Set(connNr) // Setting where we are connecting now...
		// inc the Counter
		usedConnections++ // Count how many connections have being TRIED!
		// Compare and check if this is the last Connection from the available range
		if availableConn == usedConnections {
			isLastConn = true
		}

		// Reset and set the nr of retries to 0
		nrOfRetries := 0
		// Get the max nr of retries
		maxRetries := int(conn.MaxRetries)
		if maxRetries == 0 { // 0 Means no retries, but technically it should connect 1 time
			// So we set it as 1 to be able to make the connection
			maxRetries = 1 // Set by default 1 if nothing defined
		}

		// Generate the URL where to connect!
		u := conn.GenerateURL()
		connectUrl := u.String()
		debugConn().Msg("generating url -> " + connectUrl)
		requestHeader := c.RequestHeader

		// Before Connect Event!

		eventConn("start", "OnBeforeConnectToServer").Msg("")
		c.onBeforeConnectToServer.Scan(func(k string, v interface{}) {
			v.(OnBeforeConnectToServer)()
		})
		eventConn("finish", "OnBeforeConnectToServer").Msg("")

		// Entering loop for retry
		for {
			// Incrementing the Retries
			nrOfRetries++

			infoRetry := func() *zerolog.Event {
				return infoConn().Int("retry_nr", nrOfRetries)
			}
			warnRetry := func() *zerolog.Event {
				return warnConn().Int("retry_nr", nrOfRetries)
			}
			debugRetry := func() *zerolog.Event {
				return debugConn().Int("retry_nr", nrOfRetries)
			}
			errorRetry := func() *zerolog.Event {
				return errorConn().Int("retry_nr", nrOfRetries)
			}
			eventRetry := func(eventType string, eventName string) *zerolog.Event {
				return eventConn(eventType, eventName).Int("retry_nr", nrOfRetries)
			}

			shouldBreak := false
			select {
			case <-c.ctx.Done(): // When it's stopped, or the main process or others have signaled to cancel
				// Stop the connection process!
				warnRetry().Msg("parent stop signal has being received")
				terminate = true
			case <-c.ctxConnect.Done(): // When StopConnect has being called
				// Stop the connection process!
				warnRetry().Msg("stop command signal has being received")
				stopConnectingCalled = true
				terminate = true
			default:
				if maxRetries == -1 {
					// It should loop infinitely
					// Until someone stops it!
				} else if nrOfRetries >= maxRetries+1 {
					infoRetry().
						Int("retries", nrOfRetries).
						Msg("nr of retries has being reached, breaking...")
					shouldBreak = true
					// We have reached the max nr. of retries, so we will break!
					break
				}
			}

			if shouldBreak {
				infoRetry().
					Msg("breaking...")
				break
			}

			// If terminate it's being called, we exit!
			if terminate {
				warnRetry().Msg("terminating (1)...")
				break
			}

			// Do the authentication by adding necessary data to HEADER, GET
			if conv.ParseBool(conn.EnableAuth) {
				operationMsg := "setting auth type"
				infoRetry().Msg("authentication is enabled -> " + operationMsg)
				switch conn.AuthOptions.AuthType {
				case connection.AuthByToken:
					debugRetry().
						Str("auth_type", "token").
						Str("token", conn.AuthOptions.Token).
						Msg(operationMsg)
					requestHeader.Set(authentication.DefaultHeaderAuthKey, conn.AuthOptions.Token)
				case connection.AuthByBearerToken:
					debugRetry().
						Str("auth_type", "bearer_token").
						Str("token", conn.AuthOptions.Token).
						Msg(operationMsg)
					requestHeader.Set("Authorization", "Bearer "+conn.AuthOptions.Token)
				case connection.AuthBasic:
					debugRetry().
						Str("auth_type", "basic").
						Str("username", conn.AuthOptions.Username).
						Str("password", hash.MD5(conn.AuthOptions.Password)). // md5 password for security reasons!
						Msg(operationMsg)
					requestHeader.Set("Authorization", "Basic "+base64.EncodeString(conn.AuthOptions.Username+":"+conn.AuthOptions.Password))
				case connection.AuthByGETParamToken:
					// TODO: usually the generated URL doesn't have GET Params, but if it does, we will have to
					// check for existence and then add!
					connectUrl = connectUrl + "?" + authentication.DefaultGETAuthKey + "=" + conn.AuthOptions.Token
					debugRetry().
						Str("auth_type", "get_param_token").
						Str("token", conn.AuthOptions.Token).
						Str("connect_url", connectUrl).
						Msg(operationMsg)
				}
			}

			// Connect to the host
			infoRetry().Msg(color.Style{color.LightGreen}.Render("dialing..."))

			// Get the default dialer
			dialer := websocket.DefaultDialer
			// Set TLS configuration settings
			dialer.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: conv.ParseBool(conn.AcceptCertificate),
			}

			ws, httpResponse, err := dialer.Dial(connectUrl, requestHeader)
			if err != nil {
				errorRetry().
					Int("max_retries", maxRetries).
					Int("nr_of_retries", nrOfRetries).
					Errs("errors", []error{err, ConnectionFailedToHost}).
					Msg(color.LightRed.Render("dial error"))
				// Create the error response for callbacks
				onConnectErrResponse := OnConnectErrorResponse{
					MaxRetries:    uint16(maxRetries),
					RetryNr:       uint16(nrOfRetries),
					MainError:     ConnectionFailedToHost,
					OriginalError: err,
					Connection:    conn,
					ConnectionNr:  uint16(connNr),
				}

				// TODO: handle error
				eventRetry("start", "OnError").Msg("")
				c.onError.Scan(func(k string, v interface{}) {
					v.(OnError)() // TODO: we should indicate the error type by some code!
				})
				eventRetry("finish", "OnError").Msg("")
				// Sending the main error code, and also sending the original error!
				eventRetry("start", "OnConnectFailed").Msg("")
				c.onConnectFailed.Scan(func(k string, v interface{}) {
					v.(OnConnectFailed)(onConnectErrResponse)
				})
				eventRetry("finish", "OnConnectFailed").Msg("")
				// If it's not last connection and last retry then timeout?!
				// Also if it's a single connection
				if !isLastConn && nrOfRetries != maxRetries || availableConn == 1 {
					// Between connections i don't see a necessity to have timeouts... but who knows...
					// if it's not latest connection, we should check the timeouts

					// 100 milliseconds
					sleepTime := time.Millisecond * 100
					infoRetry().
						Dur("sleep_duration_ms", sleepTime).
						Msg("sleeping...")

					select {
					case <-c.ctx.Done(): // When it's stopped, or the main process or others have signaled to cancel
						// Stop the connection process!
						terminate = true
					case <-c.ctxConnect.Done(): // When StopConnect has being called
						// Stop the connection process!
						terminate = true
					case <-time.After(time.Second * time.Duration(conn.RetryTimeout)):
						// DO nothing...
					}

					/*_time.Sleep(
						sleepTime,
						uint64(conn.RetryTimeout*10),
						// Totally 100 ms * 50 = 5000 ms = 5 seconds
						func(sleepStatus *_time.SleepStatus) {
							select {
							case <-c.ctx.Done(): // When it's stopped, or the main process or others have signaled to cancel
								// Stop the connection process!
								sleepStatus.Break = true
								terminate = true
							case <-c.ctxConnect.Done(): // When StopConnect has being called
								// Stop the connection process!
								sleepStatus.Break = true
								terminate = true
							default:
								// DO nothing...
							}
						},
					)*/
					if terminate {
						warnRetry().Msg("terminating (2)...")
						break
					}
				}

				infoRetry().
					Msg(color.Style{color.LightYellow}.Render("continuing to next retry or connection..."))

				// Continue tu next connection if that's enabled!
				continue
			}

			connectTimeMs := time.Since(c.connectStartTime.Get()).Milliseconds()
			infoRetry().
				Int64("connect_latency_ms", connectTimeMs).
				Msg(color.LightGreen.Render("connected successfully"))

			// Set to which connection we have being connected!
			c.connectedTo.Set(connNr)
			// Set locally that we have connected
			isConnected = true
			// Set the object to structure
			c.WSClient = ws
			// Set the Http Response to Structure
			c.HttpResponse = httpResponse

			infoRetry().
				Msg("breaking retrying loop")

			break // break from retry loop
		}
		// If terminate it's being called, we exit!
		if terminate {
			infoConn().Msg("terminating (3)...")
			break
		}
		// If it's connected we should break from loop!
		if isConnected {
			infoConn().Int(logConnKey, connNr).
				Msg("breaking connections loop ")
			break
		}
	}

	// If terminate it's being called, we exit!
	if terminate {
		warn().Msg("terminating (4)...")
		if stopConnectingCalled {
			_event("start", "OnStopConnectingFinish").Msg("")
			c.onStopConnectingFinish.Scan(func(k string, v interface{}) {
				v.(OnStopConnectingFinish)()
			})
			_event("finish", "OnStopConnectingFinish").Msg("")
		}
		terminateConn()
		return define.Err(0, "terminated...")
	}

	if isConnected {
		info().Msg("running reader & writer goroutines")

		// Set as being connected
		c.isConnected.True()

		// Turn on the routines

		// This is the reader!
		go c.runReader()
		// This is the writer!
		go c.runWriter()

		// When we have successfully connected to a host
		// TODO: indicate to which host
		// at what time

		// Call event
		connNr := c.connectedTo.Get()
		connConf := c.GetConnection(connNr)
		_event("start", "OnConnectSuccess").Int(logConnKey, connNr).Msg("")
		c.onConnectSuccess.Scan(func(k string, v interface{}) {
			v.(OnConnectSuccess)(connNr, connConf)
		})
		_event("finish", "OnConnectSuccess").Int(logConnKey, connNr).Msg("")
	}

	// Set that it's not connecting anymore!
	c.isConnecting.False()
	// Set when procedure has being finished
	c.connectEndTime.SetNow()
	// Failed to connect

	if !isConnected {
		// When all connections have failed to connect!
		_event("start", "OnConnectFailedAll").Msg("")
		c.onConnectFailedAll.Scan(func(k string, v interface{}) {
			v.(OnConnectFailedAll)()
		})
		_event("finish", "OnConnectFailedAll").Msg("")
		// we should launch here autoReconnect?!
	}

	if !isConnected && !c.HasOnConnectFailedAll("autoReconnect") {
		// Launching autoReconnect in a separate goroutine
		// The Callbacks will be registered afterwards...
		info().Msg("launching autoReconnect")
		go c.autoReconnect()
	}

	if isConnected {
		return nil
	} else {
		return define.Err(0, "failed to connect")
	}
}
