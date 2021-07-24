package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWT_SECRET = "ACP-10"

func GenerateJWT(customerId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["customerId"] = customerId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWT_SECRET))
}
