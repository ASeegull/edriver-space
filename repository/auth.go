package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/model"
)

type AuthRepos struct {
	db *sql.DB
}

func NewAuthRepos(db *sql.DB) *AuthRepos {
	return &AuthRepos{
		db: db,
	}
}

func (a *AuthRepos) GetUserByCredentials(ctx context.Context, login, password string) (*model.User, error) {

	return a.getUser(ctx, "SELECT * FROM accounts WHERE email = $1 AND password = $2", login, password)
}

func (a *AuthRepos) GetUserById(ctx context.Context, userId string) (*model.User, error) {

	return a.getUser(ctx, "SELECT * FROM accounts WHERE id = $1", userId)
}

func (a *AuthRepos) getUser(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	user := &model.User{}

	if err := a.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Email, &user.Password, &user.Role,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
