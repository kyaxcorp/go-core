package sender

import "time"

func (s *Sender) Start() {
	go s.start()
}

func (s *Sender) start() {
	for {
		calledStop := false
		select {
		case calledStop = <-s.stopSender:
			if calledStop {
				break
			}
			// TODO: maybe wait for a notification
		default:
			// TODO: SELECT FROM DB!?
			// TODO: it would be good to subscribe to a table and receive notifications!
			time.Sleep(time.Second)
		}
		if calledStop {
			break
		}
	}
}
