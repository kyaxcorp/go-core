package server

func (h *Hub) runController() {
	for {
		if h.StopCalled.Get() {
			break
		}
		select {
		case <-h.stopController:
			break
		case client := <-h.UnregisterClientChannel:
			h.UnRegisterClient(client)
		}
	}
}
