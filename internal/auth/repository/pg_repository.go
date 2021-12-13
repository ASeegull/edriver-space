package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/models"
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
		&foundUser.Id,
		&foundUser.Login,
		&foundUser.Password,
		&foundUser.Role,
	); err != nil {
		return nil, err
	}
	fmt.Println("found by login", foundUser)
	return foundUser, nil
}

func (r *authRepo) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}

	if err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", userID).Scan(
		&user.Id,
		&user.Login,
		&user.Password,
		&user.Role,
	)
	err != nil{
		return nil, err
	}
	fmt.Println("found by id", user)
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
