package heya

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type heyaClient struct {
	id       uuid.UUID
	userID   uuid.UUID
	heyaID   uuid.UUID
	conn     *websocket.Conn
	receiver *chan *heyaCliMessage
	sender   chan []byte
	closer   chan bool
}

type heyaCliMessage struct {
	userID uuid.UUID
	heyaid uuid.UUID
	body   []byte
}

func (hc *heyaClient) listen() {
	for {
		_, message, err := hc.conn.ReadMessage()
		if err != nil {
			hc.closer <- true
			break
		}

		*hc.receiver <- &heyaCliMessage{
			userID: hc.userID,
			heyaid: hc.heyaID,
			body:   message,
		}
	}
}

func (hc *heyaClient) serve() {
	for {
		mes := <-hc.sender

		err := hc.conn.WriteMessage(websocket.BinaryMessage, mes)
		if err != nil {
			hc.closer <- true
			break
		}
	}
}
