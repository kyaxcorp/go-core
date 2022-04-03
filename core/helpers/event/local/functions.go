package local

import (
	"errors"
	"github.com/kyaxcorp/go-core/core/helpers/function"
	"sort"
	"strconv"
	"strings"
)

func fmtListenerName(name string) string {
	return strings.ToLower(name)
}

func fmtEventName(eventName string) string {
	return strings.ToLower(eventName)
}

func createPriorityOrderIndex() map[uint16]map[string]*ListenOptions {
	return make(map[uint16]map[string]*ListenOptions)
}

func createListenersMap() map[string]*ListenOptions {
	return make(map[string]*ListenOptions)
}

func _listen(eventName string, listenOptions *ListenOptions) {

	//
	// Check if the event exists, if not then add it
	// Check if a sequence for this event exists, if not then add it as 0 and generate a new ID!
	// Use Write locks when working with this 2 vars!
	// When a process or someone registers for listening an event, we should also add
	// TTL
	// When the event destructs itself..

	if listenOptions.Callback == nil {
		if function.IsCallable(listenOptions.OnListenFailed) {
			listenOptions.OnListenFailed(errors.New("callback empty"))
		}
		return
	}

	eventName = fmtEventName(eventName)

	var _eventsStack *eventsStack
	events.ShardLock(eventName)         // TODO: TEST -> if we are getting a shard lock here... will it do a lock later?!
	defer events.ShardUnlock(eventName) // Unlock when finished
	if !events.HasNoLock(eventName) {
		// Create first time!
		_eventsStack = &eventsStack{
			ListenersMap:       createListenersMap(),
			PriorityOrderIndex: createPriorityOrderIndex(),
			eventSeq:           0,
		}
		events.SetNoLock(eventName, _eventsStack)
	} else {
		// Clear up the Call PriorityOrder and Index
		// Recreating the struct!
		currentEventStack, _ := events.GetNoLock(eventName)
		_eventsStack = &eventsStack{
			ListenersMap:       currentEventStack.(*eventsStack).ListenersMap,
			PriorityOrderIndex: createPriorityOrderIndex(),
			eventSeq:           currentEventStack.(*eventsStack).eventSeq,
		}
		events.SetNoLock(eventName, _eventsStack)
	}
	// Generate new ID
	_eventsStack.eventSeq = _eventsStack.eventSeq + 1
	id := _eventsStack.eventSeq

	// Save the listenOptions
	var strId string
	if listenOptions.Name == "" {
		strId = strconv.FormatUint(id, 10)
	} else {
		strId = fmtListenerName(listenOptions.Name)
	}

	// Save it
	_eventsStack.ListenersMap[strId] = listenOptions
	_eventsStack.generateEventOrder()

	// Give back the id!
	if function.IsCallable(listenOptions.OnEventRegisterFinish) {
		listenOptions.OnEventRegisterFinish(strId)
	}
}

func (ev *eventsStack) generateEventOrder() {
	var keysExists = make(map[int]bool)
	keys := ev.CallPriorityOrder
	listeners := ev.PriorityOrderIndex
	for k, v := range ev.ListenersMap {
		priority := v.ExecutionPriority
		priorityInt := int(v.ExecutionPriority)

		// check if thereis already and item like this, if yes, then don't add!
		if _, ok := keysExists[priorityInt]; !ok {
			keys = append(keys, priorityInt)
			keysExists[priorityInt] = true
		}

		// save the item to listeners
		if _, ok := listeners[priority]; !ok {
			listeners[priority] = make(map[string]*ListenOptions) // TODO: we should make smaller dynamic value
		}

		// Saving the Callback pointer with the priority key
		listeners[priority][k] = v
	}

	// Sort the keys...
	sort.Ints(keys)
	ev.CallPriorityOrder = keys
}
