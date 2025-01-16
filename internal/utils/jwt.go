package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/run-bigpig/conduit/internal/constants"
	"time"
)

func CreateJwt(claims map[string]interface{}, exp int) (string, error) {
	claims["exp"] = time.Now().Add(time.Duration(exp) * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString([]byte(constants.JwtSecret))
}
