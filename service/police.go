package service

import (
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/repository"
)

type PoliceService struct {
	CarFinesRepos    repository.CarFines
	DriverFinesRepos repository.DriverFines
	cfg              *config.Config
}

// NewPoliceService returns pointer to new PoliceService
func NewPoliceService(repos *repository.Repositories, cfg *config.Config) *PoliceService {
	return &PoliceService{
		CarFinesRepos:    repos.CarFines,
		DriverFinesRepos: repos.DriverFines,
		cfg:              cfg}
}

// GetFinesDriverLicense provides logic for getting all driver fines
func (ps *PoliceService) GetFinesDriverLicense(ctx context.Context, licence string) ([]model.DriversFine, error) {
	fines, err := ps.DriverFinesRepos.GetDriverFines(ctx, licence)
	if err != nil {
		return nil, err
	}
	return fines, nil
}

// GetFinesCarRegNum provides logic for getting all car fines
func (ps *PoliceService) GetFinesCarRegNum(ctx context.Context, regNum string) ([]model.CarsFine, error) {
	fines, err := ps.CarFinesRepos.GetCarFines(ctx, regNum)
	if err != nil {
		return nil, err
	}
	return fines, nil
}

// GetDriverFine provides logic for getting driver fine by its fine number
func (ps *PoliceService) GetDriverFine(ctx context.Context, fineNum string) (*model.DriversFine, error) {
	fine, err := ps.DriverFinesRepos.GetDriverFine(ctx, fineNum)
	if err != nil {
		logger.LogErr(err)
		return nil, err
	}
	return fine, nil
}

// GetCarFine provides logic for getting car fine by its fine number
func (ps *PoliceService) GetCarFine(ctx context.Context, fineNum string) (*model.CarsFine, error) {
	fine, err := ps.CarFinesRepos.GetCarFine(ctx, fineNum)
	if err != nil {
		logger.LogErr(err)
		return nil, err
	}
	return fine, nil
}

// RemoveDriverFine provides logic for removing a driver fine
func (ps *PoliceService) RemoveDriverFine(ctx context.Context, fineNum string) error {
	err := ps.DriverFinesRepos.DeleteDriverFine(ctx, fineNum)
	if err != nil {
		return err
	}
	return nil
}

// RemoveCarFine provides logic for removing a car fine
func (ps *PoliceService) RemoveCarFine(ctx context.Context, fineNum string) error {
	err := ps.CarFinesRepos.DeleteCarFine(ctx, fineNum)
	if err != nil {
		return err
	}
	return nil
}
