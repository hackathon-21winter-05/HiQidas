package ws

import (
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
)

type WSHandlerGroup struct {
	s *streamer.Streamer
}

func NewWSHandlerGroup(s *streamer.Streamer) *WSHandlerGroup {
	return &WSHandlerGroup{
		s: s,
	}
}

