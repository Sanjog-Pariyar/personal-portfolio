package controller

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"golang.org/x/oauth2"
)

type Controller struct {
	pg                Postgres
	GoogleOauthConfig *oauth2.Config
	JWTSecret         string
	cld               *cloudinary.Cloudinary
}

type Config interface {
	JwtSecret() string
	NewCloudinary() *cloudinary.Cloudinary
	GoogleOauthConfig() *oauth2.Config
}

func SetController(newPg Postgres, config Config) {
	controller.pg = newPg
	controller.JWTSecret = config.JwtSecret()
	controller.cld = config.NewCloudinary()
	controller.GoogleOauthConfig = config.GoogleOauthConfig()
}


var controller = &Controller{}

func Instance() *Controller {
	return controller
}
