package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/ws"
	"google.golang.org/protobuf/proto"
)

type HeyaHandler struct {
}

func (hs *HeyaStreamer) heyaWSHandler(mes *heyaCliMessage) error {
	var WsHeyaData *ws.WsHeyaData

	err := proto.Unmarshal(mes.body, WsHeyaData)
	if err != nil {
		return err
	}

	switch WsHeyaData.GetPayload().(type) {
	case *ws.WsHeyaData_CreateHiqidashi:
		return hs.createHiqidashiHandler(mes.heyaid, WsHeyaData.GetCreateHiqidashi())
	default:
		return nil
	}
}
