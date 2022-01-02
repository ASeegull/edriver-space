package model

import "errors"

var (
	ErrUserNotFound       = errors.New("user doesn't exists")
	ErrSessionNotFound    = errors.New("session doesn't exists")
	ErrWrongFileType      = errors.New("wrong file type")
	ErrUserWithEmailExist = errors.New("the user with such email is registered")
)
