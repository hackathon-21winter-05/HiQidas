package heya

import (
	"log"

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

func (hs *HeyaStreamer) Listen() {
	for {
		msg := <-hs.receiveBuffer

		err := hs.heyaWSHandler(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (hs *HeyaStreamer) sendToHeya(heyaID uuid.UUID, body []byte) {
	clientsID := hs.ser.GetHeyaClientsIDByHeyaID(heyaID)

	for _, clientID := range clientsID {
		hs.sendToClient(clientID, body)
	}
}

func (hs *HeyaStreamer) sendToClient(clientID uuid.UUID, body []byte) {
	hs.heyaClients[clientID].sender <- body
}
