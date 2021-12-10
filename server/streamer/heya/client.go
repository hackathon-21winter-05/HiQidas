package heya

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type client struct {
	id       uuid.UUID
	userID   uuid.UUID
	heyaID   uuid.UUID
	conn     *websocket.Conn
	receiver *chan *cliMessage
	sender   chan []byte
	closer   chan bool
}

type cliMessage struct {
	clientID uuid.UUID
	userID   uuid.UUID
	heyaid   uuid.UUID
	body     []byte
}

func (hc *client) listen() {
	for {
		_, message, err := hc.conn.ReadMessage()
		if err != nil {
			hc.closer <- true
			break
		}

		*hc.receiver <- &cliMessage{
			clientID: hc.id,
			userID:   hc.userID,
			heyaid:   hc.heyaID,
			body:     message,
		}
	}
}

func (hc *client) serve() {
	for {
		mes := <-hc.sender

		err := hc.conn.WriteMessage(websocket.BinaryMessage, mes)
		if err != nil {
			hc.closer <- true
			break
		}
	}
}
