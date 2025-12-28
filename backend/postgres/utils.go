package postgres

import (
	"errors"

	"github.com/sanjog-pariyar/user-service/errorhandler"
	"github.com/sanjog-pariyar/user-service/models"
)

func checkUsernameAndPassword(user models.User) error {
	if user.Email == "" {
		return &errorhandler.UserServiceError{
			ErrorType:     errorhandler.Invalid,
			ClientMessage: "email is required",
			Err:           errors.New("email is required"),
		}
	}

	if user.Password == "" {
		return &errorhandler.UserServiceError{
			ErrorType:     errorhandler.Invalid,
			ClientMessage: "password is required",
			Err:           errors.New("password is required"),
		}
	}

	return nil
}