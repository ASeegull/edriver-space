package model

import "errors"

var (
	ErrUserNotFound    = errors.New("user doesn't exists")
	ErrSessionNotFound = errors.New("session doesn't exists")
)
