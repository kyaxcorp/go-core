package waiter

import (
	"context"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
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
	//currentID *_uint64.Uint64
	currentID uint64

	//lock sync.RWMutex
	lock sync.Mutex
	// Cancel functions from WithCancel Context
	cancelFunctions map[uint64]context.CancelFunc
	// This are the contexts from WithCancel
	channels map[uint64]context.Context
	// Nr of waiters -> counting... when it will be 0, it will be deleted
	//nrOfWaiters map[uint64]*_uint64.Uint64
	nrOfWaiters map[uint64]uint64
	ctx         context.Context
}

func New(ctx context.Context) *Waiter {
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}
	w := &Waiter{
		//currentID:       _uint64.NewVal(0),
		currentID: 0,
		//lock:            sync.RWMutex{},
		lock:            sync.Mutex{},
		cancelFunctions: make(map[uint64]context.CancelFunc),
		channels:        make(map[uint64]context.Context),
		//nrOfWaiters:     make(map[uint64]*_uint64.Uint64),
		nrOfWaiters: make(map[uint64]uint64),
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
		w.lock.Lock()
		id := w.getCurrentID()
		// Call the cancel function...
		// Call cancel function
		w.cancelFunctions[id]()
		w.inc()
		w.lock.Unlock()
		// Increment last!
	}()
}

func (w *Waiter) getCurrentID() uint64 {
	//return w.currentID.Get()
	return w.currentID
}

// It will change the current id!
// This should be called from the outside!
// It should we called only when signalling has being done!
func (w *Waiter) inc() uint64 {
	//w.lock.Lock()
	//defer w.lock.Unlock()
	// Check if higher for reset...
	/*if w.currentID.Get() >= 709551615 {
		// Reset the counter
		w.currentID.Set(0)
	}
	id := w.currentID.Inc(1)*/
	if w.currentID >= 709551615 {
		// Reset the counter
		w.currentID = 0
	}
	w.currentID += 1
	id := w.currentID

	// Create new channel!

	// Create the with cancel, receive the cancel function and context
	ctx, cancelFunc := context.WithCancel(_context.GetDefaultContext())
	// save the cancel function
	w.cancelFunctions[id] = cancelFunc
	// save the context
	w.channels[id] = ctx
	// create the nr. of waiters var
	//w.nrOfWaiters[id] = _uint64.New()
	w.nrOfWaiters[id] = 0

	// return the current id
	return id
}

func (w *Waiter) Wait() {

	// Acquire the lock...
	w.lock.Lock()
	// Get the current id!
	id := w.getCurrentID()
	// Add +1 waiter
	//w.nrOfWaiters[id].Inc(1)
	w.nrOfWaiters[id] += 1
	waitChannel := w.channels[id]
	w.lock.Unlock()

	// The channel can be still alive for a couple of seconds... until the
	// garbage collector will not destroy it!

	defer func() {
		go func() {
			// say that this goroutine is out... -1

			w.lock.Lock()
			//w.nrOfWaiters[id].Dec(1)
			w.nrOfWaiters[id] -= 1
			// now let's check if there are any more waiters... if not
			// then let's clean the stack
			//if w.nrOfWaiters[id].Get() == 0 {
			// Lower it can't be, but anyway...
			if w.nrOfWaiters[id] <= 0 {
				// Destroy it...
				delete(w.nrOfWaiters, id)
				delete(w.channels, id)
				delete(w.cancelFunctions, id)
			}
			w.lock.Unlock()
		}()
	}()

	select {
	case <-waitChannel.Done():
		// When the signal has being sent and received!
	case <-w.ctx.Done():
		// When the parent has done...
	}
}
