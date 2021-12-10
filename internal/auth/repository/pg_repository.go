package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type authRepo struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) FindByLogin(ctx context.Context, user *models.User) (*models.User, error) {
	foundUser := &models.User{}
	if err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE login = $1;", user.Login).Scan(
		&foundUser.ID,
		&foundUser.Login,
		&foundUser.Password,
	); err != nil {
		log.Warn(err.Error())
		return nil, err
	}
	return foundUser, nil
}

func (r *authRepo) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	user := &models.User{}

	if err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", userID).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
