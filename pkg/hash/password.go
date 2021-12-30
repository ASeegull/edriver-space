package hash

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) (*SHA256Hasher, error) {
	if salt == "" {
		return nil, errors.New("empty salt")
	}

	return &SHA256Hasher{salt: salt}, nil
}

func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
