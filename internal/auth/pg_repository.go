package auth

import (
	"context"
	"github.com/ASeegull/edriver-space/internal/models"
)

type Repository interface {
	FindByLogin(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
}
