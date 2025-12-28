package controller

import (
	"net/http"

	"golang.org/x/oauth2"
	"github.com/sanjog-pariyar/user-service/utils"
)

type Controller struct {
	pg                Postgres
	GoogleOauthConfig *oauth2.Config
	JWTSecret         string
	cld               Cloudinary
}

func (c *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	c.SignUpHandler(w, r)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	c.LoginHandler(w, r)
}

func (c *Controller) ImageTransform(w http.ResponseWriter, r *http.Request) {
	res := c.cld.GetAssetInfo()
	utils.RespondWithJSON(w, http.StatusFound, res)
}

func SetController(newPg Postgres, Jwt_secret string, cloudinary Cloudinary, oAuthConfig *oauth2.Config) {
	controller.pg = newPg
	controller.JWTSecret = Jwt_secret
	controller.cld = cloudinary
	controller.GoogleOauthConfig = oAuthConfig
}

var controller = &Controller{}

func Instance() *Controller {
	return controller
}
