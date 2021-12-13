package utils
//
//import (
//	"errors"
//	"github.com/ASeegull/edriver-space/internal/models"
//	"github.com/golang-jwt/jwt"
//	"time"
//)
//
//type Claims struct {
//	Login string `json:"login"`
//	ID    int    `json:"id"`
//	jwt.StandardClaims
//}
//
//func GenerateJWTToken(user *models.User) (string, error) {
//	claims := &Claims{
//		Login: user.Login,
//		ID:    user.Id,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	tokenString, err := token.SignedString([]byte("secret_key"))
//	if err != nil {
//		return "", err
//	}
//
//	return tokenString, nil
//}
//
//func ExtractClaimsFromJWT(tokenString string) (map[string]interface{}, error) {
//	claims := jwt.MapClaims{}
//
//	token, err := jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (interface{}, error) {
//		return []byte("secret_key"), nil
//	})
//
//	if err != nil {
//		if errors.Is(err, jwt.ErrSignatureInvalid) {
//			return nil, errors.New("invalid token signature")
//		}
//		return nil, err
//	}
//
//	if !token.Valid {
//		return nil, errors.New("invalid token ")
//	}
//
//	return claims, nil
//}
