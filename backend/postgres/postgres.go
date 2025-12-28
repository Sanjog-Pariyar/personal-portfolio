package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/sanjog-pariyar/user-service/errorhandler"
	"github.com/sanjog-pariyar/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore interface {
	CreateUser(user models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User)
	GetUserById(id string) (*models.User, error)
	Signup(user models.User) (*models.User, error)
	Login(user models.User) (*models.User, error)
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
	if user.Email == "" {
		return nil, &errorhandler.UserServiceError{
			ErrorType:     errorhandler.Invalid,
			ClientMessage: "email is required",
			Err:           errors.New("email is required"),
		}
	}

	if user.Password == "" {
		return nil, &errorhandler.UserServiceError{
			ErrorType:     errorhandler.Invalid,
			ClientMessage: "password is required",
			Err:           errors.New("password is required"),
		}
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
	db, err := gorm.Open(postgres.Open(config.PostgresDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		return nil
	}

	pg := &Postgres{
		db: db,
	}

	autoMigrate(db)
	return pg
}

func autoMigrate(database *gorm.DB) {
	database.AutoMigrate(&models.User{})
	fmt.Println("Automigrate complete")
}
