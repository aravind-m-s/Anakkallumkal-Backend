package main

import (
	"log"

	"github.com/aravind-m-s/anakallumkal-backend/config"
	"github.com/aravind-m-s/anakallumkal-backend/di"
)

func main() {
	config := config.InitConfig()

	server, err := di.InitServer(config)

	if err != nil {
		log.Fatal("Uanble to connect to db", err)
	} else {
		server.Start(config)
	}

}
