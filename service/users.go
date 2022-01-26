package service

import (
	"context"
	"errors"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/hash"
	"github.com/ASeegull/edriver-space/pkg/payment"
	"github.com/ASeegull/edriver-space/repository"
	"golang.org/x/sync/errgroup"
	"time"
)

type UsersService struct {
	usersRepos       repository.Users
	sessionRepos     repository.Sessions
	carFinesRepos    repository.CarFines
	driverFinesRepos repository.DriverFines
	tokenManager     auth.TokenManager
	hasher           hash.PasswordHasher
	cfg              *config.Config
}

func NewUsersService(userRepos repository.Users, sessionRepos repository.Sessions, carFinesRepos repository.CarFines,
	driverFinesRepos repository.DriverFines, tokenManager auth.TokenManager, hasher hash.PasswordHasher, cfg *config.Config) *UsersService {
	return &UsersService{
		usersRepos:       userRepos,
		sessionRepos:     sessionRepos,
		carFinesRepos:    carFinesRepos,
		driverFinesRepos: driverFinesRepos,
		tokenManager:     tokenManager,
		hasher:           hasher,
		cfg:              cfg,
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

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	hashPassword, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := s.usersRepos.GetUserByCredentials(ctx, input.Email, hashPassword)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return Tokens{}, model.ErrUserNotFound
		}

		return Tokens{}, err
	}
	return s.createSession(ctx, user.Id)
}

func (s UsersService) RefreshTokens(ctx context.Context, sessionId string) (Tokens, error) {
	userId, err := s.sessionRepos.GetSessionById(ctx, sessionId)
	if err != nil {
		return Tokens{}, err
	}

	if err := s.sessionRepos.DeleteSession(ctx, sessionId); err != nil {
		return Tokens{}, err
	}
	return s.createSession(ctx, *userId)
}

func (s *UsersService) DeleteSession(ctx context.Context, sessionId string) error {

	return s.sessionRepos.DeleteSession(ctx, sessionId)
}

func (s *UsersService) GetUserById(ctx context.Context, userId string) (*model.User, error) {

	return s.usersRepos.GetUserById(ctx, userId)
}

func (s *UsersService) SignUp(ctx context.Context, user UserSignUpInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(user.Password)
	if err != nil {
		return Tokens{}, err
	}

	newUser := model.User{
		Firstname: &user.Firstname,
		Lastname:  &user.Lastname,
		Email:     &user.Email,
		Password:  &passwordHash,
	}

	userId, err := s.usersRepos.CreateUser(ctx, newUser)
	if err != nil {
		return Tokens{}, err
	}
	return s.createSession(ctx, userId)
}

func (s *UsersService) createSession(ctx context.Context, userId string) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	g := new(errgroup.Group)

	g.Go(func() error {
		accessTokenTTL := time.Duration(s.cfg.Token.AccessTTL) * time.Minute

		res.AccessToken, err = s.tokenManager.NewJWT(userId, "user", accessTokenTTL)

		return err
	})

	g.Go(func() error {
		res.RefreshToken, err = s.tokenManager.NewRefreshToken()

		return err
	})

	if err := g.Wait(); err != nil {
		return Tokens{}, err
	}

	refreshTokenTTL := time.Duration(s.cfg.Token.RefreshTTL) * time.Minute

	err = s.sessionRepos.SetSession(ctx, res.RefreshToken, userId, refreshTokenTTL)

	return res, err
}

type AddDriverLicenceInput struct {
	IndividualTaxNumber string
}

func (s *UsersService) AddDriverLicence(ctx context.Context, input AddDriverLicenceInput, userId string) error {
	licenceNumber, err := s.usersRepos.GetDriverLicence(ctx, input.IndividualTaxNumber)
	if err != nil {
		return err
	}

	return s.usersRepos.UpdateUserDriverLicence(ctx, userId, licenceNumber)
}

func (s *UsersService) GetFines(ctx context.Context, userId string) (model.Fines, error) {
	var (
		err error

		carsFines    []model.CarsFine
		driversFines []model.DriversFine
	)

	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		carsFines, err = s.usersRepos.GetCarsFines(groupCtx, userId)

		return err
	})

	g.Go(func() error {
		driversFines, err = s.usersRepos.GetDriversFines(groupCtx, userId)

		return err
	})

	if err := g.Wait(); err != nil {
		return model.Fines{}, err
	}

	return model.Fines{CarsFines: carsFines, DriversFines: driversFines}, nil
}

// PayFines provides logic for paying fines sent as an argument
func (s *UsersService) PayFines(ctx context.Context, fines model.Fines) error {
	var totalCost int

	// Go through all fines and collect their total price
	for _, driverFine := range fines.DriversFines {
		totalCost += driverFine.Price
	}
	for _, carFines := range fines.CarsFines {
		totalCost += carFines.Price
	}

	// Call payment service
	inputFunds := totalCost
	err := payment.DoPayment(inputFunds, totalCost)
	if err != nil {
		return err
	}

	// Payment successful, remove fines from the database
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error { // Removing driver fines
		for _, driverFine := range fines.DriversFines {
			err := s.driverFinesRepos.DeleteDriverFine(groupCtx, driverFine.FineNum)
			if err != nil {
				return err
			}
		}
		return nil
	})

	g.Go(func() error { // Removing car fines
		for _, carFine := range fines.CarsFines {
			err := s.carFinesRepos.DeleteCarFine(groupCtx, carFine.FineNum)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		return err
	}
	return nil
}

// PayFine provides logic for paying one fine of current user with given fine number
func (s *UsersService) PayFine(ctx context.Context, fines model.Fines, fineNum string) error {
	// Find needed fine in driver fines
	for _, driverFine := range fines.DriversFines {
		if driverFine.FineNum == fineNum { // If needed fine was found
			// Pay the fine
			inputFunds := driverFine.Price
			err := payment.DoPayment(inputFunds, driverFine.Price)
			if err != nil {
				return err
			}
			// Remove fine from the database
			err = s.driverFinesRepos.DeleteDriverFine(ctx, driverFine.FineNum)
			if err != nil {
				return err
			}
			return nil // If all correct
		}
	}
	// Find needed fine in car fines
	for _, carFine := range fines.CarsFines {
		if carFine.FineNum == fineNum { // If needed fine was found
			// Pay the fine
			inputFunds := carFine.Price
			err := payment.DoPayment(inputFunds, carFine.Price)
			if err != nil {
				return err
			}
			// Remove fine from the database
			err = s.carFinesRepos.DeleteCarFine(ctx, carFine.FineNum)
			if err != nil {
				return err
			}
			return nil // If all correct
		}
	}
	// In case fine was not found
	err := errors.New("there is no fine with provided fine number assigned to you")
	return err
}
