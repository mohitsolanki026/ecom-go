package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohitsolanki026/econ-go/config"
	"github.com/mohitsolanki026/econ-go/types"
	"github.com/mohitsolanki026/econ-go/utils"
)

type contextKey string 

const UserKey contextKey = "userID"

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

func WithJwtAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := getTokenFromRequest(r)
		token, err := validateToken(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID,_ := strconv.Atoi(str)

		u,err := store.GetUserById(userID)
		if  err != nil {
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey,u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w,r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["algo"])
		}
		return []byte(config.Envs.JWTSecretKey), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userID
}