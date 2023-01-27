package lock

import (
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"sync"
	"time"
)

//var eventsLocks = make(map[string]sync.RWMutex)

type stringLock struct {
	lock      *sync.Mutex
	lastUsage time.Time
}

type String struct {
	lock  *sync.Mutex
	locks map[string]*stringLock
}

func NewString() *String {
	t := &String{
		lock:  &sync.Mutex{},
		locks: make(map[string]*stringLock),
	}

	// create a goroutine, which will monitor...
	go func() {
		for {
			select {
			case <-_context.GetDefaultContext().Done():
				// leave the routine...
				return
			case <-time.After(time.Second * 300):
				// do the cleaning...
				t.lock.Lock()
				// loop through all existing locks, check the last usage, compare, and delete...
				for lockName, l := range t.locks {
					if l.lastUsage.Unix()+300 < time.Now().Unix() {
						// Delete it!
						delete(t.locks, lockName)
					}
				}
				t.lock.Unlock()
			}
		}
	}()

	return t
}

func (t *String) Lock(name string) {
	t.lock.Lock()

	var l *stringLock
	var ok bool

	if l, ok = t.locks[name]; !ok {
		l = &stringLock{
			lock:      &sync.Mutex{},
			lastUsage: time.Now(),
		}
		// Write into the Map the lock
		t.locks[name] = l
	}
	// Unlock the map lock
	l.lastUsage = time.Now()
	t.lock.Unlock()
	// call the string lock!
	l.lock.Lock()
}

func (t *String) Unlock(name string) {
	t.lock.Lock()

	var l *stringLock
	var ok bool

	if l, ok = t.locks[name]; ok {
		// Unlock the map lock
		// TODO: be careful when we will be creating the TTL and garbage collector
		l.lastUsage = time.Now()
		t.lock.Unlock()
		// call the string Unlock!
		l.lock.Unlock()
	}
}
