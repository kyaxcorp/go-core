package lock

import (
	"github.com/gofrs/flock"
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"log"
	"sync"
)

// TODO: is it correct?!
// Or should we create some kind of structure and pointing to it?! kind of an OOP

// They should be locally here!
var locks = make(map[string]*flock.Flock)
var locksLocker sync.Mutex

func getLockName(lockName string) string {
	return hash.Sha256(lockName) + ".lock"
}

func getLockPath(lockName string) string {
	return getLocksDir() + getLockName(lockName)
}

func getLocksDir() string {
	var err error = nil
	locksPath := config.GetConfig().Application.LocksPath
	locksPath, err = fsPath.GenRealPath(locksPath, true)

	if err != nil {
		log.Println(err)
	}

	if !filesystem.Exists(locksPath) {
		filesystem.MkDir(locksPath)
	}

	return locksPath
}

func FLock(lockName string, wait bool) bool {
	lockNameHash := getLockName(lockName)
	lockPath := getLockPath(lockName)

	// log.Println(lockPath)
	fileLock := flock.New(lockPath)
	var locked bool
	var err interface{} = nil
	if wait {
		err = fileLock.Lock()
		locked = true
	} else {
		locked, err = fileLock.TryLock()
	}

	if err != nil {
		// handle locking error
		log.Println("Failed to lock the process!")
		return false
	}

	if !locked {
		log.Println("Is Locked!")
		return false
	}

	// Lock inside the main process and globally!
	locksLocker.Lock()
	// Save the object into the map!
	locks[lockNameHash] = fileLock
	// Unlock!
	locksLocker.Unlock()

	//locks = append(locks, fileLock)

	return true
}

func FRelease(lockName string) {
	// Get lock name hashed form
	lockNameHash := getLockName(lockName)
	// Lock inside the Main Process! (Globally)
	locksLocker.Lock()
	// Check if there is a key with this lock name!
	if fileLock, ok := locks[lockNameHash]; ok {
		// Unlock the file lock!
		fileLock.Unlock()
	}
	// Unlock in the current process!
	locksLocker.Unlock()
}
