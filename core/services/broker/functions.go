package broker

import "github.com/KyaXTeam/go-core/v2/core/listeners/websocket/server"

func (b *Broker) GetNrOfPipes() int {
	defer b.pipesLock.RUnlock()
	b.pipesLock.RLock()
	return len(b.Pipes)
}

func (b *Broker) createPipes() *Broker {
	b.Pipes = make(map[string]*server.Hub)
	return b
}
