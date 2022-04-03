package server

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/filesystem"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetNrOfClients() uint {
	return s.c.GetNrOfClients()
}

func (s *Server) GetWSServer() *gin.Engine {
	return s.WSServer
}

func (s *Server) GetClients() map[*Client]bool {
	return s.c.GetClients()
}

// GetClientsLogPath -> returns the path where the logs for clients are stored
func (s *Server) GetClientsLogPath() string {
	// Creating clients path
	return filesystem.FilterPath(s.LoggerDirPath + filesystem.DirSeparator() + "clients" + filesystem.DirSeparator())
}
