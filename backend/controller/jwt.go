package controller

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

func (c *Controller) CreateJWT(email string) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userEmail": email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(c.JWTSecret))
}

func (c *Controller) ValidateJWT(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(c.JWTSecret), nil
	})
}
