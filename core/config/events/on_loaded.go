package events

import (
	"errors"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/function"
)

// here we store the callbacks
var onLoadedCallbacks = make(map[string]func())

func OnLoaded(callback func()) (string, error) {
	// Register the callback
	if function.IsCallable(callback) {
		//log.Println("registering...")
		// Generating a new unique id
		currentVal := conv.UInt64ToStr(callbackId.Inc(1))
		// Setting the callback
		onLoadedCallbacks[currentVal] = callback
		return currentVal, nil
	}
	return "", errors.New("invalid config callback on loaded")
}

// CallOnLoaded -> when everything is loaded, this function is being called
func CallOnLoaded() {
	//log.Println("CallOnLoaded")
	for _, callback := range onLoadedCallbacks {
		callback()
	}
}
