package bootstrap

import (
	"time"
)

type Bootstrap struct {
	// This is the time when has being bootstrapped, usually app start time
	startTime time.Time
}

func Start() *Bootstrap {
	return &Bootstrap{
		startTime: time.Now(),
		// These are the broker clients which are auto connecting to the broker servers
		//brokerClients: make(map[string]*client.Client),
	}
}

// StartForProcess -> This is the initial function which will create the bootstrap instance and register some info!
func StartForProcess() *Bootstrap {
	processBootstrap = Start()
	return processBootstrap
}

// GetProcessBootstrap -> this returns the bootstrap instance!
func GetProcessBootstrap() *Bootstrap {
	return processBootstrap
}

// GetStartTime -> return the time when started!
func (b *Bootstrap) GetStartTime() time.Time {
	return b.startTime
}

// GetRunningTime -> get the running time of the current app
func (b *Bootstrap) GetRunningTime() time.Duration {
	return time.Since(b.startTime)
}
