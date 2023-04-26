package lock

import (
	"github.com/gofrs/flock"
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/errors2/define"
	fsPath "github.com/kyaxcorp/go-core/core/helpers/filesystem/path"
	"github.com/kyaxcorp/go-core/core/helpers/folder"
	"github.com/kyaxcorp/go-core/core/helpers/hash"
	"github.com/kyaxcorp/go-core/core/logger/coreLog"
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

func getLockPath(lockName string) (string, error) {
	locksDir, locksDirErr := getLocksDir()
	if locksDirErr != nil {
		return "", locksDirErr
	}
	return locksDir + getLockName(lockName), nil
}

func getLocksDir() (string, error) {
	var pathErr error
	locksPath := config.GetConfig().Application.LocksPath
	coreLog.Info().Str("application_locks_path", locksPath).Msg("application locks path")

	locksPath, pathErr = fsPath.GenRealPath(locksPath, true)

	if pathErr != nil {
		return "", pathErr
	}

	if !folder.Exists(locksPath) {
		folder.MkDir(locksPath)
	}

	coreLog.Info().Str("locks_dir", locksPath).Msg("application generated real path")

	return locksPath, nil
}

func FLock(lockName string, wait bool) (bool, error) {
	coreLog.Info().
		Str("lock_name", lockName).
		Bool("wait", wait).
		Msg("FLock called")
	defer coreLog.Info().Msg("leaving...")
	lockNameHash := getLockName(lockName)
	coreLog.Info().
		Str("lock_name_hashed", lockNameHash).
		Msg("lock name hashed, getting lock path")
	lockPath, lockPathErr := getLockPath(lockName)
	if lockPathErr != nil {
		coreLog.Error().Err(lockPathErr).Msg("failed to get lock path")
		return false, lockPathErr
	}

	coreLog.Info().Str("lock_path", lockPath).Msg("lock path retrieved")

	// log.Println(lockPath)
	fileLock := flock.New(lockPath)
	var locked bool
	//var err interface{} = nil
	var lockErr error
	if wait {
		lockErr = fileLock.Lock()
		locked = true
	} else {
		locked, lockErr = fileLock.TryLock()
	}

	if lockErr != nil {
		// handle locking error
		return false, define.Err(0, "failed to lock the process -> ", lockErr.Error())
	}

	if !locked {
		//log.Println("Is Locked!")
		return false, nil
	}

	// Lock inside the main process and globally!
	locksLocker.Lock()
	// Save the object into the map!
	locks[lockNameHash] = fileLock
	// Unlock!
	locksLocker.Unlock()

	//locks = append(locks, fileLock)

	return true, nil
}

func FRelease(lockName string) {
	// Get lock name hashed form
	lockNameHash := getLockName(lockName)
	// Lock inside the Main Check! (Globally)
	locksLocker.Lock()
	// Check if there is a key with this lock name!
	if fileLock, ok := locks[lockNameHash]; ok {
		// Unlock the file lock!
		fileLock.Unlock()
	}
	// Unlock in the current process!
	locksLocker.Unlock()
}
