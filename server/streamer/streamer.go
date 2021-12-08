package streamer

import (
	"log"

	"github.com/gofrs/uuid"
)

type Streamer struct {
	clients       map[uuid.UUID]*client
	receiveBuffer chan *cliMessage
}

func NewStreamer() *Streamer {
	s := &Streamer{
		clients:       map[uuid.UUID]*client{},
		receiveBuffer: make(chan *cliMessage),
	}

	return s
}

func (s *Streamer) listen() {
	for {
		msg := <-s.receiveBuffer
		log.Print(*msg)
	}
}

func (s *Streamer) Run() {
	go s.listen()
}
