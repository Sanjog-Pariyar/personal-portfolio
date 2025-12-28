package main

import (
	"log"

	"github.com/sanjog-pariyar/user-service/api"
	"github.com/sanjog-pariyar/user-service/cloudinary"
	"github.com/sanjog-pariyar/user-service/config"
	"github.com/sanjog-pariyar/user-service/controller"
	"github.com/sanjog-pariyar/user-service/internal"
	"github.com/sanjog-pariyar/user-service/postgres"
)

func main() {
	newConfig := config.NewConfig()
	newPg := postgres.NewPostgres(newConfig.Postgres)
	newCloudinary := cloudinary.NewCloudinary(newConfig.Cloudinary.CloudinaryUrl())
	newServer := api.NewServer(newConfig.Api)
	controller.SetController(newPg, newConfig.Jwt_secret, newCloudinary, newConfig.Google_Handler.GoogleOauthConfig())

	internal.Waiting(newServer)

	log.Println("Server exiting")
}
