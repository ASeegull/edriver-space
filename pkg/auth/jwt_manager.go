package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(userId, userRole string, ttl time.Duration) (string, error)
	Parse(accessToken string) (UserClaims, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

type CustomClaims struct {
	UserId   string
	UserRole string
	jwt.StandardClaims
}

func (m *Manager) NewJWT(userId string, userRole string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserId:   userId,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	})

	return token.SignedString([]byte(m.signingKey))
}

type UserClaims struct {
	Id   string
	Role string
}

func (m *Manager) Parse(accessToken string) (UserClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return UserClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserClaims{}, fmt.Errorf("error get user claims from token")
	}

	return UserClaims{
		Id:   claims["UserId"].(string),
		Role: claims["UserRole"].(string),
	}, nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
