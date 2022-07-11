package server

import "github.com/kyaxcorp/go-core/core/helpers/function"

func (h *Hub) runGetter() {
	// On Start callback
	if function.IsCallable(h.onStartGetter) {
		h.onStartGetter(h)
	}
	// TODO: we should adapt here in a cron or something like
	// 		that...
	for {
		if h.StopCalled.Get() {
			break
		}
		//select {
		//case <-h.stopGetter:
		//	break
		//}
		if function.IsCallable(h.getter) {
			h.getter(h)
		}
	}
}
