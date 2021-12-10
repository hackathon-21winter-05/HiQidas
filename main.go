package main

import (
	"log"
	"time"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server"
)

func init() {
	// タイムゾーンの設定
	const location = "Asia/Tokyo"

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	s := server.NewServer(c)

	s.Run()
}
