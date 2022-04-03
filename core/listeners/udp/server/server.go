package server

import (
	"context"
	"errors"
	"fmt"
	"net"
)

// maxBufferSize specifies the size of the buffers that
// are used to temporarily hold data from the UDP packets
// that we receive.
const DefaultMaxBufferSize = 1024
const DefaultListeningAddress = "0.0.0.0:33333"

type OnPacketFunc func(recvMessage RecvMessage)

type Server struct {
	// The max buffer size
	MaxBufferSize int
	// It also includes port
	ListeningAddress string
	// Context
	Ctx context.Context
	// The server itself!
	server net.PacketConn
	// On packet callback!
	onPacket OnPacketFunc
	// Run in Goroutine when we have received a packet
	onPacketAsync bool
}

// You can use the default constructor!
func New() *Server {
	server := &Server{
		MaxBufferSize:    DefaultMaxBufferSize,
		ListeningAddress: DefaultListeningAddress,
		Ctx:              nil,
		server:           nil,
		onPacket:         nil,
		onPacketAsync:    false,
	}
	return server
}

func (s *Server) Stop() (bool, error) {
	return true, nil
}

func (s *Server) OnPacketReceive(async bool, callback OnPacketFunc) *Server {
	s.onPacketAsync = async
	if callback != nil {
		s.onPacket = callback
	}
	return s
}

type RecvMessage struct {
	Text      string
	Bytes     []byte
	NrOfBytes int
	Length    int
	IPAddress string
	Server    *Server
}

// server wraps all the UDP echo server functionality.
// ps.: the server is capable of answering to a single
// client at a time.
func (s *Server) Start() (bool, error) {

	// ListenPacket provides us a wrapper around ListenUDP so that
	// we don't need to call `net.ResolveUDPAddr` and then subsequentially
	// perform a `ListenUDP` with the UDP address.
	//
	// The returned value (PacketConn) is pretty much the same as the one
	// from ListenUDP (UDPConn) - the only difference is that `Packet*`
	// methods and interfaces are more broad, also covering `ip`.

	var err error
	s.server, err = net.ListenPacket("udp", s.ListeningAddress)
	if err != nil {
		return false, errors.New("Failed to start listening!")
	}

	// `Close`ing the packet "connection" means cleaning the data structures
	// allocated for holding information about the listening socket.
	defer s.server.Close()

	doneChan := make(chan error, 1)
	buffer := make([]byte, s.MaxBufferSize)

	// Given that waiting for packets to arrive is blocking by nature and we want
	// to be able of canceling such action if desired, we do that in a separate
	// go routine.
	go func() {
		for {
			// By reading from the connection into the buffer, we block until there's
			// new content in the socket that we're listening for new packets.
			//
			// Whenever new packets arrive, `buffer` gets filled and we can continue
			// the execution.
			//
			// note.: `buffer` is not being reset between runs.
			//	  It's expected that only `n` reads are read from it whenever
			//	  inspecting its contents.

			n, addr, err := s.server.ReadFrom(buffer)

			if err != nil {
				doneChan <- err
				return
			}

			msg := RecvMessage{
				Text:      string(buffer[:n]),
				Bytes:     buffer[:n],
				NrOfBytes: n,
				Length:    n,
				IPAddress: addr.String(),
				Server:    s,
			}

			if s.onPacketAsync {
				// Run as Goroutine
				go s.onPacket(msg)
			} else {
				// Run in this Routin
				s.onPacket(msg)
			}

			// Setting a deadline for the `write` operation allows us to not block
			// for longer than a specific timeout.
			//
			// In the case of a write operation, that'd mean waiting for the send
			// queue to be freed enough so that we are able to proceed.

			/*deadline := time.Now().Add(*timeout)
			err = pc.SetWriteDeadline(deadline)
			if err != nil {
				doneChan <- err
				return
			}

			// Write the packet's contents back to the client.
			n, err = pc.WriteTo(buffer[:n], addr)
			if err != nil {
				doneChan <- err
				return
			}

			fmt.Printf("packet-written: bytes=%d to=%s\n", n, addr.String())*/
		}
	}()

	select {
	case <-s.Ctx.Done():
		fmt.Println("cancelled")
		err = s.Ctx.Err()
	case err = <-doneChan:
	}

	return true, nil
}
