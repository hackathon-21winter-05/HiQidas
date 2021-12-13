package yjs

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type client struct {
	id          uuid.UUID
	hiqidashiID uuid.UUID
	userID      uuid.UUID
	conn        *websocket.Conn
	receiver    *chan *cliMessage
	sender      chan []byte
	closer      chan bool
}

type cliMessage struct {
	clientID    uuid.UUID
	hiqidashiID uuid.UUID
	userID    uuid.UUID
	body        []byte
}

func (c *client) listen() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.closer <- true
			break
		}

		*c.receiver <- &cliMessage{
			clientID:    c.id,
			hiqidashiID: c.hiqidashiID,
			userID:    c.userID,
			body:        message,
		}
	}
}

func (c *client) serve() {
	for {
		mes := <-c.sender

		err := c.conn.WriteMessage(websocket.BinaryMessage, mes)
		if err != nil {
			c.closer <- true
			break
		}
	}
}
