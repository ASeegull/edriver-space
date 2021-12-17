package service

import (
	"context"
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/models"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/repository"
	"time"
)

type UserSignUpInput struct {
	Email    string
	Password string
}

type UserSignInInput struct {
	Login    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	authRepos    repository.Auth
	sessionRepos repository.Sessions
	tokenManager auth.TokenManager
	cfg          *config.Config
}

func NewAuthService(repos *repository.Repositories, tokenManager auth.TokenManager, cfg *config.Config) *AuthService {
	return &AuthService{
		authRepos:    repos.Auth,
		sessionRepos: repos.Sessions,
		tokenManager: tokenManager,
		cfg:          cfg,
	}
}

func (a *AuthService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	user, err := a.authRepos.GetUserByCredentials(ctx, input.Login, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return Tokens{}, models.ErrUserNotFound
		}

		return Tokens{}, err
	}
	return a.createSession(ctx, user.Id)
}

func (a AuthService) RefreshTokens(ctx context.Context, sessionId string) (Tokens, error) {
	userId, err := a.sessionRepos.GetSessionById(ctx, sessionId)
	if err != nil {
		return Tokens{}, err
	}

	if err := a.sessionRepos.DeleteSession(ctx, sessionId); err != nil {
		return Tokens{}, err
	}
	return a.createSession(ctx, *userId)
}

func (a *AuthService) DeleteSession(ctx context.Context, sessionId string) error {

	return a.sessionRepos.DeleteSession(ctx, sessionId)
}

func (a *AuthService) GetUserById(ctx context.Context, userId string) (*models.User, error) {

	return a.authRepos.GetUserById(ctx, userId)
}

func (a *AuthService) createSession(ctx context.Context, userId string) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	accessTokenTTL := time.Duration(a.cfg.TokensTTL.Access) * time.Minute

	res.AccessToken, err = a.tokenManager.NewJWT(userId, "user", accessTokenTTL)
	if err != nil {
		return Tokens{}, nil
	}

	res.RefreshToken, err = a.tokenManager.NewRefreshToken()
	if err != nil {
		return Tokens{}, nil
	}

	refreshTokenTTL := time.Duration(a.cfg.TokensTTL.Refresh) * time.Minute

	err = a.sessionRepos.SetSession(ctx, res.RefreshToken, userId, refreshTokenTTL)

	return res, err
}
