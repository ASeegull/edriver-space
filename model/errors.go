package model

import "errors"

var (
	ErrUserNotFound          = errors.New("user doesn't exists")
	ErrSessionNotFound       = errors.New("session doesn't exists")
	ErrDriverLicenceNotFound = errors.New("driver licence doesn't exists")
	ErrWrongFileType         = errors.New("wrong file type")
	ErrUserWithEmailExist    = errors.New("the user with such email is registered")
	ErrJWTEmpty              = errors.New("JWT token has empty body")
)
