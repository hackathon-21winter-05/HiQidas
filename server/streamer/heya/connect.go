package heya

import (
	"context"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/ws"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (hs *HeyaStreamer) ConnectHeyaWS(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	userID := sess.Values["user_id"].(uuid.UUID)

	heyaIDString := c.Param("heyaid")
	heyaID, err := uuid.FromString(heyaIDString)
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

	if err := hs.sendHiqidashis(clientID, heyaID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	hs.ser.AddHeyaClient(heyaID, clientID)
	hs.heyaClients[clientID] = cli

	<-cli.closer

	_ = hs.ser.DeleteHeyaClient(heyaID, clientID)

	delete(hs.heyaClients, clientID)

	return nil
}

func (hs *HeyaStreamer) sendHiqidashis(clientID, heyaID uuid.UUID) error {
	hiqidashis, err := hs.ser.GetHiqidashisByHeyaID(context.Background(), heyaID)
	if err != nil {
		return err
	}

	res := []*ws.Hiqidashi{}
	for _, hiqidashi := range hiqidashis {
		resHiqidashi := &ws.Hiqidashi{
			Id:          hiqidashi.ID.String(),
			CreatorId:   hiqidashi.CreatorID.String(),
			Title:       hiqidashi.Title,
			Description: hiqidashi.Description,
			ColorCode:   hiqidashi.ColorCode,
		}
		if hiqidashi.ParentID.Valid {
			resHiqidashi.ParentId = &wrapperspb.StringValue{Value: hiqidashi.ParentID.UUID.String()}
		}
		if hiqidashi.Drawing.Valid {
			resHiqidashi.Drawing = &wrapperspb.StringValue{Value: hiqidashi.Drawing.String}
		}

		res = append(res, resHiqidashi)
	}

	buffer, err := proto.Marshal(&ws.WsHeyaData{
		Payload: &ws.WsHeyaData_SendHiqidashis{
			SendHiqidashis: &ws.WsSendHiqidashis{
				Hiqidashis: res,
			}}})
	if err != nil {
		return err
	}

	hs.sendToClient(clientID, buffer)

	return nil
}
