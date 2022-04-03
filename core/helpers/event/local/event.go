package local

import "time"

// Listen -> Start listening for an event
func Listen(eventName string, listenOptions *ListenOptions) {
	go _listen(eventName, listenOptions)
}

// RemoveListener -> Remove an existing listener
func RemoveListener(eventName string, id string) {
	events.ShardLock(eventName)
	defer events.ShardUnlock(eventName)
	var ev *eventsStack
	tmp, ok := events.GetNoLock(eventName)
	if !ok {
		return
	}
	ev = tmp.(*eventsStack)
	delete(ev.ListenersMap, id)
	ev.generateEventOrder()
}

// RemoveEvent -> It will remove event entirely! (all the registered listeners will be removed)
// This is easier if you want to purge all listeners for an event
func RemoveEvent(eventName string) {
	events.Remove(eventName)
}

// DispatchAsync ->
func DispatchAsync(eventName string, dispatchData *DispatchData) {
	go Dispatch(eventName, dispatchData)
}

// Dispatch ->
func Dispatch(eventName string, dispatchData *DispatchData) {
	// Use read locks when using events var!
	/*
		Check if exists any events
		Loop through founded listeners
		See who and how should be called and launched
		See for who the data needs to be with pointer or in original form before execution
		See who and how should be organized as execution, async first or sync... or just as registered!
		Priority it's also important for listeners... some modules may require to do something after
		Other events have finished their work!
	*/

	events.ShardRLock(eventName)

	defer events.ShardRUnlock(eventName)
	var ev *eventsStack
	tmp, ok := events.GetNoLock(eventName)
	if !ok {
		return
	}

	ev = tmp.(*eventsStack)
	// if there are no events
	if len(ev.ListenersMap) == 0 {
		//log.Println("listeners map empty...")
		return
	}

	// Loop through registered callbacks
	// TODO -> we should loop through cached by priorities
	for _, order := range ev.CallPriorityOrder {
		evs := ev.PriorityOrderIndex[uint16(order)]
		for _, e := range evs {
			call := func() {
				dispatchData.Time = time.Now()     // Time when it has being triggered
				dispatchData.EventName = eventName // Set the event name that has being triggered
				e.Callback(dispatchData)
			}
			// TODO: check TTL
			// If it has expirated... call function to be removed!

			if e.Async {
				go call()
			} else {
				call()
			}
		}
	}
}
