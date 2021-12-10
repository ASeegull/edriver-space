package usecase

import (
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/models"
	"github.com/ASeegull/edriver-space/pkg/httpErrors"
	"github.com/ASeegull/edriver-space/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type authUC struct {
	authRepo auth.Repository
	cfg      *config.Config
}

func NewAuthUseCase(authRepo auth.Repository, cfg *config.Config) auth.UseCase {
	return &authUC{
		authRepo: authRepo,
		cfg:      cfg,
	}
}

func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	foundUser, err := u.authRepo.FindByLogin(ctx, user)
	if err != nil {
		log.Warn(err)
		return nil, httpErrors.NewRestError(http.StatusInternalServerError, err.Error())
	}

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		log.Warn(err)
		return nil, httpErrors.NewRestError(http.StatusUnauthorized, err.Error())
	}

	foundUser.SanitizePassword()

	token, err := utils.GenerateJWTToken(foundUser)
	if err != nil {
		log.Warn(err)
		return nil, httpErrors.NewRestError(http.StatusInternalServerError, err.Error())
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}

func (u *authUC) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	return u.authRepo.GetUserByID(ctx, userID)
}
