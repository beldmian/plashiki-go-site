package main

import (
	"log"

	"github.com/beldmian/plashiki-go-site/internal/app/server"
)

var (
	configPath string
)

func main() {
	config := server.NewConfig()
	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
