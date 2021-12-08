package main

import (
	"log"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/server"
	"github.com/hackathon-21winter-05/HiQidas/service"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	repo, err := repository.NewGormRepository(c)
	if err != nil {
		log.Panic(err)
	}

	ser := service.NewService(repo)

	s := server.NewServer(c, ser)

	s.Run()
}
