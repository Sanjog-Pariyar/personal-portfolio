package models

import (
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt int64     `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password,omitempty"`
}

type LoginRequest struct {
	Name     string
	Email    string
	Password string
}

type UserResponse struct {
	Email string    `json:"email"`
	Token string    `json:"token"`
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
}

func HashPwd(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePwd(hash string, pwd []byte) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	return err == nil
}
