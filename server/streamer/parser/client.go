package parser

import (
	"github.com/gorilla/websocket"
)

type client struct {
	conn     *websocket.Conn
	receiver *chan []byte
	sender   chan []byte
	closer   chan bool
}

func (c *client) listen() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.closer <- true
			break
		}

		*c.receiver <- message
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
