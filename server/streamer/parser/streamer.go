package parser

import (
	"context"
	"database/sql"
	"log"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/parser"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"google.golang.org/protobuf/proto"
)

type ParserStreamer struct {
	client        []*client
	receiveBuffer chan []byte
	ser           *service.Service
}

func NewParserStreamer(ser *service.Service) *ParserStreamer {
	s := &ParserStreamer{
		client:        []*client{},
		receiveBuffer: make(chan []byte),
		ser:           ser,
	}

	return s
}

func (ps *ParserStreamer) Listen() {
	for {
		msg := <-ps.receiveBuffer

		editDescription := &parser.ParserEditDescription{}
		err := proto.Unmarshal(msg, editDescription)
		if err != nil {
			log.Println(err)
			continue
		}

		hiqidashiID, err := uuid.FromString(editDescription.Description.HiqidashiId)
		if err != nil {
			log.Println(err)
			continue
		}

		err = ps.ser.UpdateHiqidashiByID(context.Background(), &model.NullHiqidashi{
			ID: hiqidashiID,
			Description: sql.NullString{
				String: editDescription.Description.Content,
				Valid:  true,
			},
		})
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (hs *ParserStreamer) sendParserMes(msg *parser.ParserSendData) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	for _, client := range hs.client {
		client.sender <- body
	}

	return nil
}
