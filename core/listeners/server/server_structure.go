package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

/*
This is a Standard format but it's not being used because we need to Export the Vars to be available in other structs!
DO NOT USE IT! YOU CAN USE IT ONLY AS A TEMPLATE
*/

type OnStart func(s *Server)
type OnStop func(s *Server)

type Server struct {
	connectionID  uint64
	genConnIDLock sync.Mutex
	// Starting time of the server
	startTime time.Time
	// Stop time of the server
	stopTime time.Time

	// Logger
	L            *logrus.Logger
	LoggingLevel logrus.Level
	// Enable logging
	enableLogging bool

	// Enables Server Status through HTTP
	enableServerStatus bool

	// It also includes port
	ListeningAddress string
	// Context
	Ctx context.Context

	onStart OnStart
	onStop  OnStop
}

func (s *Server) SetLoggingLevel(level logrus.Level) *Server {
	s.LoggingLevel = level
	s.L.Level = s.LoggingLevel
	return s
}

func (s *Server) EnableLogging(level logrus.Level) *Server {
	s.enableLogging = true
	s.SetLoggingLevel(level)
	return s
}

func (s *Server) DisableLogging() *Server {
	s.enableLogging = false
	// Only error level!
	s.SetLoggingLevel(logrus.ErrorLevel)
	return s
}
