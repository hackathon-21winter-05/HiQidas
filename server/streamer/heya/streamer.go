package heya

import (
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type HeyaStreamer struct {
	heyaClients   map[uuid.UUID]*heyaClient
	receiveBuffer chan *heyaCliMessage
	ser           *service.Service
}

func NewHeyaStreamer(ser *service.Service) *HeyaStreamer {
	s := &HeyaStreamer{
		heyaClients:   map[uuid.UUID]*heyaClient{},
		receiveBuffer: make(chan *heyaCliMessage),
		ser:           ser,
	}

	return s
}

func (s *HeyaStreamer) Listen() {
	for {
		msg := <-s.receiveBuffer
		s.heyaWSHandler(msg)
	}
}
