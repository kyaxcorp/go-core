package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kyaxcorp/go-core/core/helpers/info"
	"strings"
)

type FullStatus struct {
	Name                  string
	ListeningAddresses    []string
	ListeningAddressesSSL []string
	CurrentConnectionID   uint64
	NrOfClients           uint
	SystemStatus          info.SystemStatus
}

type SystemStatus struct {
	SystemStatus info.SystemStatus
}

type NrOfClientsStatus struct {
	CurrentConnectionID uint64
	NrOfClients         uint
}

type Status struct {
	Name                  string
	ListeningAddresses    []string
	ListeningAddressesSSL []string
	CurrentConnectionID   uint64
	NrOfClients           uint
}

func (s *FullStatus) Collect() {

}

func (s *Server) Status(onCollected func(status FullStatus)) {
	go func() {
		status := FullStatus{
			Name:                  "",
			ListeningAddresses:    s.ListeningAddresses,
			ListeningAddressesSSL: s.ListeningAddressesSSL,
			CurrentConnectionID:   s.connectionID.Get(),
			NrOfClients:           s.GetNrOfClients(),
			SystemStatus:          info.GetSystemStatus(),
		}

		if onCollected != nil {
			onCollected(status)
		}
	}()
}

func (s *Server) StatusSystemStatus(onCollected func(status SystemStatus)) {
	go func() {
		status := SystemStatus{
			SystemStatus: info.GetSystemStatus(),
		}

		if onCollected != nil {
			onCollected(status)
		}
	}()
}

func (s *Server) ServerStatus(onCollected func(status Status)) {
	go func() {
		status := Status{
			Name:                  "",
			ListeningAddresses:    s.ListeningAddresses,
			ListeningAddressesSSL: s.ListeningAddressesSSL,
			CurrentConnectionID:   s.connectionID.Get(),
			NrOfClients:           s.GetNrOfClients(),
		}

		if onCollected != nil {
			onCollected(status)
		}
	}()
}

func (s *Server) StatusNrOfClients(onCollected func(status NrOfClientsStatus)) {
	go func() {
		status := NrOfClientsStatus{
			CurrentConnectionID: s.connectionID.Get(),
			NrOfClients:         s.GetNrOfClients(),
		}

		if onCollected != nil {
			onCollected(status)
		}
	}()
}

func (s *Server) startServerStatus() *Server {
	// TODO: add authentication details
	/*
		TODO: create a group
		1. add different listeners for specific statuses
		2. add clients
		3. add hubs
		4. different detailed info...
	*/

	// This function it's being called when Accessing through Http Method!!
	getStatus := func(context *gin.Context) {
		// Creating a channel for awaiting a response from a goroutine
		awaitStatus := make(chan interface{})
		// Calling the status function, which returns us a FullStatus Object! This object we afterwards convert to JSON

		exploded := strings.Split(context.Request.RequestURI, "/")

		switch exploded[len(exploded)-1] {
		case "server":
			s.ServerStatus(func(status Status) {
				// We have received the status, and we return through channel the response!
				awaitStatus <- status
			})
		case "system_status":
			s.StatusSystemStatus(func(status SystemStatus) {
				// We have received the status, and we return through channel the response!
				awaitStatus <- status
			})
		case "nr_of_clients":
			s.StatusNrOfClients(func(status NrOfClientsStatus) {
				// We have received the status, and we return through channel the response!
				awaitStatus <- status
			})
		default:
			s.Status(func(status FullStatus) {
				// We have received the status, and we return through channel the response!
				awaitStatus <- status
			})
		}

		//log.Println(context.Request.RequestURI)

		// Here we receive the status
		status := <-awaitStatus
		// We are sending the response to the client!
		context.IndentedJSON(200, status)
	}

	authorized := s.HttpServer.Group("/", gin.BasicAuth(gin.Accounts{
		s.statusUsername: s.statusPassword,
	}))

	serverStatus := authorized.Group("/server_status")
	{
		serverStatus.GET("/", getStatus)
		serverStatus.GET("/server", getStatus)
		serverStatus.GET("/nr_of_clients", getStatus)
		serverStatus.GET("/system_status", getStatus)
	}
	return s
}

func (s *Server) SetStatusCredentials(username string, password string) *Server {
	s.statusUsername = username
	s.statusPassword = password
	return s
}

func (s *Server) stopServerStatus() *Server {
	// TODO: remove from route!
	return s
}

func (s *Server) EnableServerStatus() *Server {
	s.enableServerStatus.Set(true)
	s.startServerStatus()
	return s
}

func (s *Server) DisableServerStatus() *Server {
	s.enableServerStatus.Set(false)
	s.stopServerStatus()
	return s
}
