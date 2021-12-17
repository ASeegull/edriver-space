package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/models"
)

type AuthRepos struct {
	db *sql.DB
}

func NewAuthRepos(db *sql.DB) *AuthRepos {
	return &AuthRepos{
		db: db,
	}
}

func (a *AuthRepos) GetUserByCredentials(ctx context.Context, login, password string) (*models.User, error) {

	return a.getUser(ctx, "SELECT * FROM accounts WHERE email = $1 AND password = $2", login, password)
}

func (a *AuthRepos) GetUserById(ctx context.Context, userId string) (*models.User, error) {

	return a.getUser(ctx, "SELECT * FROM accounts WHERE id = $1", userId)
}

func (a *AuthRepos) getUser(ctx context.Context, query string, args ...interface{}) (*models.User, error) {
	user := &models.User{}

	if err := a.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Email, &user.Password, &user.Role,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
