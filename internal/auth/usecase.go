package auth

import (
	"context"
	"github.com/ASeegull/edriver-space/internal/models"
)

type UseCase interface {
	Login(ctx context.Context, user *models.User) (*models.UserWithTokens, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}
