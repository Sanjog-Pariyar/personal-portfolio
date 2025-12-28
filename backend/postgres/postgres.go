package postgres

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/sanjog-pariyar/user-service/errorhandler"
	"github.com/sanjog-pariyar/user-service/models"
	"gorm.io/gorm"
)

type Config interface {
	NewPostgres() *gorm.DB
}

type Postgres struct {
	db *gorm.DB
}

func (pg *Postgres) CreateUser(user models.User) (*models.User, error) {

	user.ID, _ = uuid.NewV4()
	hPass, e := models.HashPwd([]byte(user.Password))

	if e != nil {
		return nil, &errorhandler.UserServiceError{
			ErrorType:     errorhandler.Internal,
			ClientMessage: "Internal error, Unable to hash password",
			Err:           fmt.Errorf("Unable to hash password %w", e),
		}
	}

	user.Password = hPass

	err := pg.db.Create(&user).Error

	if err != nil {
		return nil, &errorhandler.UserServiceError{
			ErrorType:     errorhandler.Internal,
			ClientMessage: "Internal error, Unable to create new user",
			Err:           fmt.Errorf("Unable to create new user %w", err),
		}
	}

	return &user, nil
}

func (pg *Postgres) GetUserById(id string) (*models.User, error) {
	var user models.User

	pg.db.First(&user, "id = ?", id)

	return &user, nil
}

func (pg *Postgres) GetUserByEmail(email string) *models.User {

	var user models.User

	pg.db.First(&user, "email = ?", email)

	return &user
}

func (pg *Postgres) Signup(user models.User) (*models.User, error) {
	err := checkUsernameAndPassword(user)

	if err != nil {
		return  nil, err
	}

	existedUser := pg.GetUserByEmail(user.Email)

	if user.Email == existedUser.Email {
		return nil, &errorhandler.UserServiceError{
			ErrorType:     errorhandler.AlreadyExist,
			ClientMessage: "User with email already exist in the database",
			Err:           errors.New("email already exists"),
		}
	}

	newUser, err := pg.CreateUser(user)
	if err != nil {
		return nil, err
	}
	newUser.Password = ""
	return newUser, nil
}

func (pg *Postgres) Login(user models.User) (*models.User, error) {
	err := checkUsernameAndPassword(user)

	if err != nil {
		return  nil, err
	}

	existedUser := pg.GetUserByEmail(user.Email)

	if len(existedUser.Email) == 0 {
		return nil, &errorhandler.UserServiceError{
			ErrorType:     errorhandler.NotFound,
			ClientMessage: "invalid email",
			Err:           errors.New("invalid email"),
		}
	}

	if models.ComparePwd(existedUser.Password, []byte(user.Password)) {
		existedUser.Password = ""
		return existedUser, nil
	}
	return nil, &errorhandler.UserServiceError{
		ErrorType:     errorhandler.Invalid,
		ClientMessage: "invalid password",
		Err:           errors.New("invalid password"),
	}
}

func NewPostgres(config Config) *Postgres {

	pg := &Postgres{
		db: config.NewPostgres(),
	}
	return pg
}
