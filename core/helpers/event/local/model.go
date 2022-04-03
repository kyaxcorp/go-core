package local

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/cmap"
	"time"
)

var events = cmap.New(cmap.MapConstructor{})

// These are the options for listening
type ListenOptions struct {
	// This is the listener name, a unique identifier,
	// even if the event will be recreated, it will be simply overwritten!
	// This param is Optional! If not defined, it will be replaced with an automated ID generator!
	Name                  string
	Callback              func(dispatchData *DispatchData)
	Async                 bool            // The callback will be run async or sync mode (goroutine or without)
	TTL                   uint64          // Time to live in the stack
	ExecutionPriority     uint16          // This is the execution priority in the stack
	createdAt             time.Time       // When it has being created
	executionPriorityId   string          // This is the ID that's being set back by the system when recreating the order!
	OnListenFailed        func(err error) // This is when it has failed to listen
	OnEventRegisterFinish func(id string) //This is the callback which will be called when the registration will be finished!
	// We will need it for deleting the Listener from the ordering map
}

// This is the data that comes from the dispatcher
type DispatchData struct {
	Time      time.Time // When it has being dispatched
	Id        uint64    // It's optional, and sometimes it can be empty! But it can be very useful!
	Data      interface{}
	EventName string // This is the event that has being dispatched
}

type eventsStack struct {
	ListenersMap      map[string]*ListenOptions // These are the listeners
	CallPriorityOrder []int                     // This is the order how the listeners are being called
	// In the priority index we save same listeners but getting them by order id, this is more for dispatching part
	PriorityOrderIndex map[uint16]map[string]*ListenOptions
	eventSeq           uint64
}
