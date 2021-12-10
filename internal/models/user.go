package models

import "errors"

type User struct {
	ID    int    `json:"id,omitempty"`
	Login string `json:"login"`
	Password string `json:"password"`
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

type UserWithToken struct {
	User  *User
	Token string
}
