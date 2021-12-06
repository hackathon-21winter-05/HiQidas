package main

import (
	"log"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	s := server.NewServer(c)
	s.Run()
}
