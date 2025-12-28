package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sanjog-pariyar/user-service/models"
	"github.com/sanjog-pariyar/user-service/utils"
)

type Postgres interface {
	CreateUser(user models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User)
	GetUserById(id string) (*models.User, error)
	Signup(user models.User) (*models.User, error)
	Login(user models.User) (*models.User, error)
}

func (c *Controller) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("errr ", err)
		utils.RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	newUser, err := c.pg.Signup(user)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, newUser)
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("errr ", err)
		utils.RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	loggedInUser, err := c.pg.Login(user)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, loggedInUser)
}