package server

import (
	"github.com/kyaxcorp/go-core/core/helpers/function"
)

func (s *Server) OnStart(name string, callback OnStart) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onStart.Set(name, callback)
	return true
}

func (s *Server) OnBeforeStart(name string, callback OnStart) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onBeforeStart.Set(name, callback)
	return true
}

func (s *Server) OnStarted(name string, callback OnStart) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onStarted.Set(name, callback)
	return true
}

func (s *Server) OnStartRemove(name string) {
	s.onStart.Del(name)
}

func (s *Server) OnBeforeStartRemove(name string) {
	s.onBeforeStart.Del(name)
}

func (s *Server) OnStartedRemove(name string) {
	s.onStarted.Del(name)
}

func (s *Server) OnBeforeStop(name string, callback OnStop) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onBeforeStop.Set(name, callback)
	return true
}

func (s *Server) OnStop(name string, callback OnStop) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onStop.Set(name, callback)
	return true
}

func (s *Server) OnStopped(name string, callback OnStop) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onStopped.Set(name, callback)
	return true
}

func (s *Server) OnBeforeStopRemove(name string) {
	s.onBeforeStop.Del(name)
}

func (s *Server) OnStopRemove(name string) {
	s.onStop.Del(name)
}

func (s *Server) OnStoppedRemove(name string) {
	s.onStopped.Del(name)
}

func (s *Server) OnRequest(name string, callback OnRequest) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onRequest.Set(name, callback)
	return true
}

func (s *Server) OnRequestRemove(name string) {
	s.onRequest.Del(name)
}

func (s *Server) OnResponse(name string, callback OnResponse) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	s.onResponse.Set(name, callback)
	return true
}

func (s *Server) OnResponseRemove(name string) {
	s.onResponse.Del(name)
}
