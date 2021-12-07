package streamer

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type client struct {
	id       uuid.UUID
	userID   uuid.UUID
	roomID   uuid.UUID
	conn     *websocket.Conn
	receiver *chan *cliMessage
	sender   chan []byte
	closer   chan bool
}

type cliMessage struct {
	userID uuid.UUID
	roomid uuid.UUID
	body   []byte
}

func (cli *client) listen() {
	for {
		_, message, err := cli.conn.ReadMessage()
		if err != nil {
			cli.closer <- true
			break
		}

		*cli.receiver <- &cliMessage{
			userID: cli.userID,
			roomid: cli.roomID,
			body:   message,
		}
	}
}

func (cli *client) serve() {
	for {
		mes := <-cli.sender

		err := cli.conn.WriteMessage(websocket.BinaryMessage, mes)
		if err != nil {
			cli.closer <- true
			break
		}
	}
}
