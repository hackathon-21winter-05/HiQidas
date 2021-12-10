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

func (cli *heyaClient) listen() {
	for {
		_, message, err := cli.conn.ReadMessage()
		if err != nil {
			cli.closer <- true
			break
		}

		*cli.receiver <- &heyaCliMessage{
			userID: cli.userID,
			heyaid: cli.heyaID,
			body:   message,
		}
	}
}

func (cli *heyaClient) serve() {
	for {
		mes := <-cli.sender

		err := cli.conn.WriteMessage(websocket.BinaryMessage, mes)
		if err != nil {
			cli.closer <- true
			break
		}
	}
}
