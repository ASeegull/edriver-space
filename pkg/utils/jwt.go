package utils

import (
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Login string `json:"login"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWTToken(user *models.User) (string, error) {
	claims := &Claims{
		Login: user.Login,
		ID:    string(rune(user.ID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
