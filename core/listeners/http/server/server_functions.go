package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
)

func (s *Server) GetNrOfClients() uint {
	return s.c.GetNrOfClients()
}

func (s *Server) GetHttpServer() *gin.Engine {
	return s.HttpServer
}

func (s *Server) GetClients() map[*Client]bool {
	return s.c.GetClients()
}

// GetClientsLogPath -> returns the path where the logs for clients are stored
func (s *Server) GetClientsLogPath() string {
	// Creating clients path
	return file.FilterPath(s.LoggerDirPath + filesystem.DirSeparator() + "clients" + filesystem.DirSeparator())
}

func (s *Server) GetClientsOrderedByConnectionID() map[int64]*Client {
	return s.c.GetClientsOrderedByConnectionID()
}
