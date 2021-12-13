package usecase

import (
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/auth"
	"github.com/ASeegull/edriver-space/internal/models"
	jwt "github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/httpErrors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type authUC struct {
	authRepo   auth.Repository
	cfg        *config.Config
	jwtManager *jwt.JWTManager
}

func NewAuthUseCase(authRepo auth.Repository, cfg *config.Config, jwtManager *jwt.JWTManager) auth.UseCase {
	return &authUC{
		authRepo:   authRepo,
		cfg:        cfg,
		jwtManager: jwtManager,
	}
}

func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithTokens, error) {
	foundUser, err := u.authRepo.FindByLogin(ctx, user)
	if err != nil {
		log.Warn(err)
		return nil, httpErrors.NewRestError(http.StatusInternalServerError, err.Error())
	}

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		log.Warn(err)
		return nil, httpErrors.NewRestError(http.StatusUnauthorized, err.Error())
	}

	ttl := time.Duration(u.cfg.Tokens.AccessTokenTTL) * time.Minute

	accessToken, err := u.jwtManager.NewJWT(foundUser.Id, ttl)
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.jwtManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &models.UserWithTokens{
		User: foundUser,
		Tokens: &models.JWTTokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (u *authUC) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	return u.authRepo.GetUserByID(ctx, userID)
}
