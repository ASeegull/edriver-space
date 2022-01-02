package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/model"
	"github.com/lib/pq"
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

	return a.getUser(ctx, "SELECT * FROM users WHERE email = $1 AND password = $2", login, password)
}

func (a *AuthRepos) GetUserById(ctx context.Context, userId string) (*model.User, error) {

	return a.getUser(ctx, "SELECT * FROM users WHERE id = $1", userId)
}

func (a *AuthRepos) CreateUser(ctx context.Context, email, password string) (string, error) {
	_, err := a.db.ExecContext(ctx, "INSERT INTO users(email, password) VALUES ($1, $2)", email, password)
	if err != nil {
		// convert to postgres error
		if err, ok := err.(*pq.Error); ok {
			// unique_violation
			if err.Code == "23505" {
				return "", model.ErrUserWithEmailExist
			}
		}
		return "", err
	}

	var userId string

	if err := a.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", email).Scan(&userId); err != nil {
		return "", err
	}
	return userId, nil
}

func (a *AuthRepos) getUser(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	user := &model.User{}

	if err := a.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Email, &user.Password, &user.Role, &user.DriverLicenseNumber,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
