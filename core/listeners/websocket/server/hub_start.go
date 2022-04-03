package server

// You can start the Hub before the websocket server has being started!
// This is the function which starts the hub!
func (h *Hub) Start() *Hub {
	if h.isRunning.Get() {
		return h
	}
	h.StopCalled.False()

	// We add this hub to the server stack in this moment!
	h.server.Hubs[h] = true

	// Start first the broadcast
	go h.run()
	// Start after that the reader!
	go h.runGetter()
	// Start Controller
	go h.runController()

	// On Start callback
	if h.onStart != nil {
		h.onStart(h)
	}

	return h
}
