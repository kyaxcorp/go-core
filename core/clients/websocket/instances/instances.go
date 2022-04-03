package instances

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/websocket/client"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
	"sync"
)

// Here we store the created instances...
var instances = make(map[string]*client.Client)

// This is the locker when writing and reading the instances
var instancesLock sync.RWMutex

func SaveInstance(instanceName string, server *client.Client) {
	instancesLock.Lock()
	if _, ok := instances[instanceName]; !ok {
		instances[instanceName] = server
	}
	instancesLock.Unlock()
}

func GetInstance(instanceName string) (*client.Client, error) {
	instancesLock.RLock()
	defer instancesLock.RUnlock()
	if instance, ok := instances[instanceName]; ok {
		// Return the existing instance
		return instance, nil
	}
	return nil, define.Err(0, "websocket client instance missing")
}
