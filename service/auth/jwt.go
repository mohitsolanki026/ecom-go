package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohitsolanki026/econ-go/config"
)

func CreateJwtToken(UserId int) (string, error) {
	// Create JWT token
	expirationTime := time.Second * time.Duration(config.Envs.JWTExpirationInSecond)
	secret := []byte(config.Envs.JWTSecretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    strconv.Itoa(UserId),
		"expiredAT": time.Now().Add(expirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
