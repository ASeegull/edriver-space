package auth

import (
	"context"
	"github.com/ASeegull/edriver-space/internal/models"
)

type UseCase interface {
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
}
