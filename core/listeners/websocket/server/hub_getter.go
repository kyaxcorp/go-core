package server

import "github.com/KyaXTeam/go-core/v2/core/helpers/function"

func (h *Hub) runGetter() {
	// On Start callback
	if function.IsCallable(h.onStartGetter) {
		h.onStartGetter(h)
	}
	//
	for {
		if h.StopCalled.Get() {
			break
		}
		if function.IsCallable(h.getter) {
			h.getter(h)
		}
	}
}
