package heya

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (hs *HeyaStreamer) ConnectHeyaWS(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	userID := sess.Values["user_id"].(uuid.UUID)

	heyaIDString := c.Param("roomid")
	heyaID, err := uuid.FromString(heyaIDString)
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

	hs.ser.AddHeyaClient(heyaID, clientID)

	cli := &heyaClient{
		id:       clientID,
		userID:   userID,
		heyaID:   heyaID,
		conn:     conn,
		receiver: &hs.receiveBuffer,
		sender:   make(chan []byte),
		closer:   make(chan bool),
	}

	go cli.serve()
	go cli.listen()

	hs.heyaClients[clientID] = cli

	<-cli.closer

	_ = hs.ser.DeleteHeyaClient(heyaID, clientID)

	delete(hs.heyaClients, clientID)

	return nil
}
