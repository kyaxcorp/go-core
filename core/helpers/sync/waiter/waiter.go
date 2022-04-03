package waiter

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/sync/_uint64"
	"sync"
)

/*
We need to register
We need to be destroyed
We need to signal
The purpose of this object is to notify other that something has hapened?!...
Kind of a broadcast? But in mean time, it should also block the execution of the goroutines... meaning that
they will not sleep, they'll just wait for a signal and be blocked!

So: there are waiters, and there is the signaler
The signaler can work fast and should always receive the list of recepients,
The recepients should fast receive the signal and leave...
This can be done with cancelContext!
*/

type Waiter struct {
	currentID *_uint64.Uint64
	lock      sync.RWMutex
	// Cancel functions from WithCancel Context
	cancelFuncs map[uint64]context.CancelFunc
	// This are the contexts from WithCancel
	channels map[uint64]context.Context
	// Nr of waiters -> counting... when it will be 0, it will be deleted
	nrOfWaiters map[uint64]*_uint64.Uint64
	ctx         context.Context
}

func New(ctx context.Context) *Waiter {
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}
	w := &Waiter{
		currentID:   _uint64.NewVal(0),
		lock:        sync.RWMutex{},
		cancelFuncs: make(map[uint64]context.CancelFunc),
		channels:    make(map[uint64]context.Context),
		nrOfWaiters: make(map[uint64]*_uint64.Uint64),
		ctx:         ctx,
	}
	w.inc()
	return w
}

func (w *Waiter) Signal() {
	// Send the signal...
	// Get the current id!
	// Call the cancel function
	// the garbage collector should remove this channels from the map after some time...!
	// It can remove by the nr of waiters... the waiter should be counted!!
	// When signal has being sent, we can increment to next stage!
	// Also we can set as to be deleted...

	go func() {
		id := w.getCurrentID()
		// Call the cancel function...
		w.lock.RLock()
		// Call cancel function
		w.cancelFuncs[id]()
		w.lock.RUnlock()
		// Increment last!
		w.inc()
	}()
}

func (w *Waiter) getCurrentID() uint64 {
	return w.currentID.Get()
}

// It will change the current id!
// This should be called from the outside!
// It should we called only when signalling has being done!
func (w *Waiter) inc() uint64 {
	defer w.lock.Unlock()
	w.lock.Lock()
	// Check if higher for reset...
	if w.currentID.Get() >= 709551615 {
		// Reset the counter
		w.currentID.Set(0)
	}
	id := w.currentID.Inc(1)
	// Create new channel!

	ctx, cancelFunc := context.WithCancel(_context.GetDefaultContext())
	w.cancelFuncs[id] = cancelFunc
	w.channels[id] = ctx
	w.nrOfWaiters[id] = _uint64.New()

	return id
}

func (w *Waiter) Wait() {
	// Send the signal...

	// Get the current id!
	w.lock.RLock()
	id := w.getCurrentID()
	// Add +1 waiter
	w.nrOfWaiters[id].Inc(1)
	waitChannel := w.channels[id]
	w.lock.RUnlock()

	// The channel can be still alive for a couple of seconds... until the
	// garbage collector will not destroy it!

	defer func() {
		go func() {
			w.nrOfWaiters[id].Dec(1)
			if w.nrOfWaiters[id].Get() == 0 {
				// Destroy it...
				w.lock.Lock()
				delete(w.nrOfWaiters, id)
				delete(w.channels, id)
				delete(w.cancelFuncs, id)
				w.lock.Unlock()
			}
		}()
	}()

	select {
	case <-waitChannel.Done():
		// When the signal has being sent and received!
	case <-w.ctx.Done():
		// When the parent has done...
	}
}
