package service

import (
	"context"
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/hash"
	"github.com/ASeegull/edriver-space/repository"
	"golang.org/x/sync/errgroup"
	"time"
)

type UsersService struct {
	usersRepos   repository.Users
	sessionRepos repository.Sessions
	tokenManager auth.TokenManager
	hasher       hash.PasswordHasher
	cfg          *config.Config
}

func NewUsersService(userRepos repository.Users, sessionRepos repository.Sessions,
	tokenManager auth.TokenManager, hasher hash.PasswordHasher, cfg *config.Config) *UsersService {
	return &UsersService{
		usersRepos:   userRepos,
		sessionRepos: sessionRepos,
		tokenManager: tokenManager,
		hasher:       hasher,
		cfg:          cfg,
	}
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type UserSignUpInput struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (a *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	hashPassword, err := a.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := a.usersRepos.GetUserByCredentials(ctx, input.Email, hashPassword)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return Tokens{}, model.ErrUserNotFound
		}

		return Tokens{}, err
	}
	return a.createSession(ctx, user.Id)
}

func (a UsersService) RefreshTokens(ctx context.Context, sessionId string) (Tokens, error) {
	userId, err := a.sessionRepos.GetSessionById(ctx, sessionId)
	if err != nil {
		return Tokens{}, err
	}

	if err := a.sessionRepos.DeleteSession(ctx, sessionId); err != nil {
		return Tokens{}, err
	}
	return a.createSession(ctx, *userId)
}

func (a *UsersService) DeleteSession(ctx context.Context, sessionId string) error {

	return a.sessionRepos.DeleteSession(ctx, sessionId)
}

func (a *UsersService) GetUserById(ctx context.Context, userId string) (*model.User, error) {

	return a.usersRepos.GetUserById(ctx, userId)
}

func (a *UsersService) SignUp(ctx context.Context, user UserSignUpInput) (Tokens, error) {
	passwordHash, err := a.hasher.Hash(user.Password)
	if err != nil {
		return Tokens{}, err
	}

	newUser := model.User{
		Firstname: &user.Firstname,
		Lastname:  &user.Lastname,
		Email:     &user.Email,
		Password:  &passwordHash,
	}

	userId, err := a.usersRepos.CreateUser(ctx, newUser)
	if err != nil {
		return Tokens{}, err
	}
	return a.createSession(ctx, userId)
}

func (a *UsersService) createSession(ctx context.Context, userId string) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	g := new(errgroup.Group)

	g.Go(func() error {
		accessTokenTTL := time.Duration(a.cfg.TokensTTL.Access) * time.Minute

		res.AccessToken, err = a.tokenManager.NewJWT(userId, "user", accessTokenTTL)

		return err
	})

	g.Go(func() error {
		res.RefreshToken, err = a.tokenManager.NewRefreshToken()

		return err
	})

	if err := g.Wait(); err != nil {
		return Tokens{}, err
	}

	refreshTokenTTL := time.Duration(a.cfg.TokensTTL.Refresh) * time.Minute

	err = a.sessionRepos.SetSession(ctx, res.RefreshToken, userId, refreshTokenTTL)

	return res, err
}
