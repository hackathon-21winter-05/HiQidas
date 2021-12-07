package streamer

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (s *Streamer) ConnectWs(c echo.Context) error {
	roomIDString := c.Param("roomid")
	roomID, err := uuid.FromString(roomIDString)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid room ID")
	}

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close()

	clientID, err := uuid.NewV4()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	cli := &client{
		id:       clientID,
		roomID:   roomID,
		conn:     conn,
		receiver: &s.receiveBuffer,
		sender:   make(chan []byte),
		closer:   make(chan bool),
	}

	go cli.serve()
	go cli.listen()

	s.clients[clientID] = cli

	<-cli.closer

	delete(s.clients, clientID)

	return nil
}
