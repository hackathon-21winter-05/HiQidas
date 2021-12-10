package yjs

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (hs *YjsStreamer) ConnectHiqidashiWS(c echo.Context) error {
	hiqidashiIDString := c.Param("hiqidashiid")
	hiqidashiID, err := uuid.FromString(hiqidashiIDString)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid heyaID")
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
		id:          clientID,
		hiqidashiID: hiqidashiID,
		conn:        conn,
		receiver:    &hs.receiveBuffer,
		sender:      make(chan []byte),
		closer:      make(chan bool),
	}

	go cli.serve()
	go cli.listen()

	hs.ser.AddEditorClient(hiqidashiID, clientID)
	hs.clients[clientID] = cli

	<-cli.closer

	_ = hs.ser.DeleteEditorClient(hiqidashiID, clientID)

	delete(hs.clients, clientID)

	return nil
}
