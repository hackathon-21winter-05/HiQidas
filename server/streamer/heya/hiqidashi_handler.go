package heya

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/ws"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (hs *HeyaStreamer) createHiqidashiHandler(userID, heyaID uuid.UUID, body *ws.WsCreateHiqidashi) error {
	parentUUID, err := uuid.FromString(body.GetParentId())
	if err != nil {
		return err
	}

	created, err := hs.ser.CreateHiqidashi(context.Background(), userID, heyaID, parentUUID)
	if err != nil {
		return err
	}

	var drawing *wrapperspb.StringValue = nil
	if created.Drawing.Valid {
		drawing = &wrapperspb.StringValue{Value: created.Drawing.String}
	}

	res := &ws.WsSendHiqidashi{
		Hiqidashi: &ws.Hiqidashi{
			Id:          created.ID.String(),
			CreatorId:   created.CreatorID.String(),
			ParentId:    created.ParentID.UUID.String(),
			Title:       created.Title,
			Description: created.Description,
			Drawing:     drawing,
			ColorCode:   created.ColorCode,
		},
	}

	err = hs.sendHeyaMes(heyaID,
		&ws.WsHeyaData{
			Payload: &ws.WsHeyaData_SendHiqidashi{
				SendHiqidashi: res,
			},
		})
	if err != nil {
		return err
	}

	return nil
}
