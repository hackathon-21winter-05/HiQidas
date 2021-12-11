package parser

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/parser"

	"github.com/labstack/echo/v4"
)

func (ps *ParserStreamer) ConnectParserWS(c echo.Context) error {
	if len(ps.client) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Already connected")
	}

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer conn.Close()

	cli := &client{
		conn:     conn,
		receiver: &ps.receiveBuffer,
		sender:   make(chan []byte),
		closer:   make(chan bool),
	}

	ps.client = append(ps.client, cli)

	go cli.serve()
	go cli.listen()

	err = ps.sendHiqidashiDescriptions()
	if err != nil {
		ps.client = ps.client[:len(ps.client)-1]
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	<-cli.closer

	ps.client = ps.client[:len(ps.client)-1]

	return nil
}

func (ps *ParserStreamer) sendHiqidashiDescriptions() error {
	hiqidashis, err := ps.ser.GetHiqidashis(context.Background())
	if err != nil {
		return err
	}

	descriptions := make([]*parser.Description, 0)
	for _, hiqidashi := range hiqidashis {
		descriptions = append(descriptions, &parser.Description{
			HiqidashiId: hiqidashi.ID.String(),
			Content:     hiqidashi.Description,
		})
	}

	sendData := &parser.ParserSendData{
		Payload: &parser.ParserSendData_ParserDescriptions{
			ParserDescriptions: &parser.ParserDescriptions{
				Descriptions: descriptions,
			},
		}}

	err = ps.sendParserMes(sendData)
	if err != nil {
		return err
	}

	return nil
}
