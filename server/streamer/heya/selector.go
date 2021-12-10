package heya

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/ws"
	"google.golang.org/protobuf/proto"
)

func (hs *HeyaStreamer) handlerSelector(mes *cliMessage) error {
	var WsHeyaData *ws.WsHeyaData

	err := proto.Unmarshal(mes.body, WsHeyaData)
	if err != nil {
		_ = hs.sendErrorMes(mes.clientID, err.Error())
		return err
	}

	switch WsHeyaData.GetPayload().(type) {
	case *ws.WsHeyaData_CreateHiqidashi:
		err := hs.createHiqidashiHandler(mes.userID, mes.heyaid, WsHeyaData.GetCreateHiqidashi())
		if err != nil {
			_ = hs.sendErrorMes(mes.clientID, err.Error())
			return err
		}
		return nil

	case *ws.WsHeyaData_EditHiqidashi:
		err := hs.editHiqidashiHandler(mes.userID, mes.heyaid, WsHeyaData.GetEditHiqidashi())
		if err != nil {
			_ = hs.sendErrorMes(mes.clientID, err.Error())
			return err
		}
		return nil

	case *ws.WsHeyaData_DeleteHiqidashi:
		err := hs.deleteHiqidashiHandler(mes.heyaid, WsHeyaData.GetDeleteHiqidashi())
		if err != nil {
			_ = hs.sendErrorMes(mes.clientID, err.Error())
			return err
		}
		return nil

	default:
		_ = hs.sendErrorMes(mes.clientID, "unknown payload")
		return errors.New("unknown payload")
	}
}

func (hs *HeyaStreamer) sendHeyaMes(heyaID uuid.UUID, mes *ws.WsHeyaData) error {
	buffer, err := proto.Marshal(mes)
	if err != nil {
		return err
	}

	hs.sendToHeya(heyaID, buffer)
	return nil
}

func (hs *HeyaStreamer) sendErrorMes(clientID uuid.UUID, message string) error {
	mes, err := proto.Marshal(
		&ws.WsHeyaData{
			Payload: &ws.WsHeyaData_Error{
				Error: &ws.WsError{Message: message},
			},
		})
	if err != nil {
		return err
	}

	hs.sendToClient(clientID, mes)

	return nil
}
