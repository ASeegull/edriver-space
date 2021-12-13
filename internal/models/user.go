package models

import "errors"

type User struct {
	Id       string    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u User) ComparePasswords(password string) error {
	if u.Password != password {
		return errors.New("passwords are not equal")
	}
	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

type UserWithTokens struct {
	User   *User
	Tokens *JWTTokens
}
