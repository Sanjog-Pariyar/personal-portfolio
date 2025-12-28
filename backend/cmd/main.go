package main

import (
	"log"

	"github.com/sanjog-pariyar/user-service/api"
	"github.com/sanjog-pariyar/user-service/config"
	"github.com/sanjog-pariyar/user-service/controller"
	"github.com/sanjog-pariyar/user-service/postgres"
)

func main() {

	newConfig := config.NewConfig()

	newPg := postgres.NewPostgres(newConfig)

	controller.SetController(newPg, newConfig)
	
	server := api.Instance()
	server.Start(newConfig.Addr())


	log.Println("Server exiting")
}
