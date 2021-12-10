package heya

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/ws"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
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

func (hs *HeyaStreamer) editHiqidashiHandler(userID, heyaID uuid.UUID, body *ws.WsEditHiqidashi) error {
	uuid, err := uuid.FromString(body.GetId())
	if err != nil {
		return err
	}

	hiqidashi := &model.NullHiqidashi{}
	hiqidashi.ID = uuid
	if body.GetTitle() != nil {
		hiqidashi.Title = utils.NullStringFrom(body.GetTitle().Value)
	}
	hiqidashi.Description = utils.NullString()
	if body.GetDrawing() != nil {
		hiqidashi.Drawing = utils.NullStringFrom(body.GetDrawing().Value)
	}
	if body.GetColorCode() != nil {
		hiqidashi.ColorCode = utils.NullStringFrom(body.GetColorCode().Value)
	}
	hiqidashi.LastEditorID = userID
	hiqidashi.UpdatedAt = time.Now()

	err = hs.ser.UpdateHiqidashiByID(context.Background(), hiqidashi)
	if err != nil {
		return err
	}

	err = hs.sendHeyaMes(heyaID,
		&ws.WsHeyaData{
			Payload: &ws.WsHeyaData_EditHiqidashi{
				EditHiqidashi: body,
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func (hs *HeyaStreamer) deleteHiqidashiHandler(heyaID uuid.UUID, body *ws.WsDeleteHiqidashi) error {
	uuid, err := uuid.FromString(body.GetId())
	if err != nil {
		return err
	}

	err = hs.ser.DeleteHiqidashiByID(context.Background(), uuid)
	if err != nil {
		return err
	}

	err = hs.sendHeyaMes(heyaID,
		&ws.WsHeyaData{
			Payload: &ws.WsHeyaData_DeleteHiqidashi{
				DeleteHiqidashi: body,
			},
		})
	if err != nil {
		return err
	}

	return nil
}
