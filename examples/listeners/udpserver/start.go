package udpserver

import (
	"context"
	"github.com/kyaxcorp/go-core/core/listeners/udp/server"
	"log"
)

var GlobalContext = context.Background()

func create() {
	s := server.New()
	s.Ctx = GlobalContext
	s.ListeningAddress = "0.0.0.0:33333"
	s.OnPacketReceive(func(msg string, bytes int, addr string, ss *server.Server) {
		log.Println(msg)
	})
	s.Start()
}
