package broker

import (
	"github.com/KyaXTeam/go-core/v2/core/listeners/websocket/server"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"time"
)

// This function monitors existing pipes
func (b *Broker) hubsMonitoring() {
	info := func() *zerolog.Event {
		return b.LInfoF("hubsMonitoring")
	}
	/*warn := func() *zerolog.Event {
		return b.LWarnF("hubsMonitoring")
	}
	_error := func() *zerolog.Event {
		return b.LErrorF("hubsMonitoring")
	}*/
	info().Msg("entering...")
	defer info().Msg("leaving...")

	shutdownCalled := false
	counter := time.Now()
	for {
		select {
		case isShutdownCalled := <-b.shutdownHubMonitoring:
			if isShutdownCalled {
				info().Msg("shutting down...")
				shutdownCalled = true
			}
		default:
			// 1 time at 10 minutes?! //
			// It will close the hubs that don't have connections for 10 minutes
			time.Sleep(time.Second)
			if time.Now().Unix() < counter.Unix()+600 {
				break // Breaking from select!
			}
			counter = time.Now() // Reset the counter!

			nrOfHubs := b.GetNrOfPipes()
			if nrOfHubs > 0 {
				now := time.Now()
				cleanPipes := make(map[string]*server.Hub)
				b.pipesLock.RLock()
				for pipeName, pipe := range b.Pipes {
					if pipe.GetNrOfClients() != 0 {
						continue
					}

					// 600 seconds -> 10 minutes
					if now.Unix() > pipe.GetCreatedTime().Unix()+600 {
						cleanPipes[pipeName] = pipe
					}
				}
				b.pipesLock.RUnlock()
				if len(cleanPipes) > 0 {
					// We should shutdown the Hubs! and then delete them from Pipes var!
					// Also, we should mark them as in deletion, or simply do the lock!
					b.pipesLock.Lock()
					for pipeName, pipe := range cleanPipes {
						info().Str("pipe_name", pipeName).Msg(color.Style{color.LightYellow}.Render("closing pipe/hub"))
						go pipe.Stop()            // Stopping the Hub
						delete(b.Pipes, pipeName) // Deleting
					}
					b.pipesLock.Unlock()
				}
			}
		}
		if shutdownCalled {
			break
		}
	}
}
