package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ASeegull/edriver-space/model"
	"github.com/lib/pq"
)

type UsersRepos struct {
	db *sql.DB
}

func NewUsersRepos(db *sql.DB) *UsersRepos {
	return &UsersRepos{
		db: db,
	}
}

func (r *UsersRepos) GetUserByCredentials(ctx context.Context, email, password string) (*model.User, error) {

	return r.getUser(ctx, "SELECT * FROM users WHERE email = $1 AND password = $2", email, password)
}

func (r *UsersRepos) GetUserById(ctx context.Context, userId string) (*model.User, error) {

	return r.getUser(ctx, "SELECT * FROM users WHERE id = $1", userId)
}

func (r *UsersRepos) CreateUser(ctx context.Context, newUser model.User) (string, error) {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(firstname, lastname, email, password) VALUES ($1, $2, $3, $4)",
		newUser.Firstname, newUser.Lastname, newUser.Email, newUser.Password)
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

	if err := r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", newUser.Email).Scan(&userId); err != nil {
		return "", err
	}
	return userId, nil
}

func (r *UsersRepos) getUser(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	user := &model.User{}

	if err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Role, &user.DriverLicenseNumber,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

func (r *UsersRepos) GetDriverLicence(ctx context.Context, itn string) (string, error) {
	var licenceNumber string
	if err := r.db.QueryRowContext(ctx, "SELECT licence_number FROM drivers WHERE individual_tax_number = $1", itn).Scan(&licenceNumber); err != nil {
		if err == sql.ErrNoRows {
			return "", model.ErrDriverLicenceNotFound
		}
		return "", err
	}
	return licenceNumber, nil
}

func (r UsersRepos) UpdateUserDriverLicence(ctx context.Context, userId, driverLicenceNumber string) error {
	result, err := r.db.ExecContext(ctx, "UPDATE users SET driver_licence_number = $1 WHERE id = $2", driverLicenceNumber, userId)
	if err != nil {
		return err
	}
	if num, _ := result.RowsAffected(); num != 1 {
		return errors.New("noting to update")
	}
	return nil
}
