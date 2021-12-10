package yjs

import (
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type YjsStreamer struct {
	clients       map[uuid.UUID]*client
	receiveBuffer chan *cliMessage
	ser           *service.Service
}

func NewHeyaStreamer(ser *service.Service) *YjsStreamer {
	s := &YjsStreamer{
		clients:       map[uuid.UUID]*client{},
		receiveBuffer: make(chan *cliMessage),
		ser:           ser,
	}

	return s
}

func (hs *YjsStreamer) Listen() {
	for {
		msg := <-hs.receiveBuffer

		hs.sendToHiqidashi(msg.clientID, msg.hiqidashiID, msg.body)
	}
}

func (hs *YjsStreamer) sendToHiqidashi(selfID, hiqidashiID uuid.UUID, body []byte) {
	clientsID := hs.ser.GetEditorClientsIDByHiqidashiID(hiqidashiID)

	for _, clientID := range clientsID {
		if clientID != selfID {
			hs.sendToClient(clientID, body)
		}
	}
}

func (hs *YjsStreamer) sendToClient(clientID uuid.UUID, body []byte) {
	hs.clients[clientID].sender <- body
}
