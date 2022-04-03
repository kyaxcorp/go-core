package instances

import (
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"github.com/kyaxcorp/go-core/core/listeners/http/server"
	"sync"
)

// Here we store the created instances...
var instances = make(map[string]*server.Server)

// This is the locker when writing and reading the instances
var instancesLock sync.RWMutex

/*
There are instances from configuration
But the user can also can create analogic instances or duplicates based on the config instances
All instances should be saved as reference in a global var
*/

func SaveInstance(instanceName string, server *server.Server) {
	instancesLock.Lock()
	if _, ok := instances[instanceName]; !ok {
		instances[instanceName] = server
	}
	instancesLock.Unlock()
}

func GetInstance(instanceName string) (*server.Server, error) {
	instancesLock.RLock()
	defer instancesLock.RUnlock()
	if instance, ok := instances[instanceName]; ok {
		// Return the existing instance
		return instance, nil
	}
	return nil, define.Err(0, "http server instance missing")
}
