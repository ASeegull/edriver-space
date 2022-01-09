package service

import (
	"context"

	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/repository"
)

type driverService struct {
	cfg        *config.Config
	driverRepo repository.Drivers
}

func NewDriverService(cfg *config.Config, driverRepo repository.Drivers) *driverService {
	return &driverService{cfg: cfg, driverRepo: driverRepo}
}

func (d *driverService) CreateDriver(ctx context.Context, driver *model.Driver) (*model.Driver, error) {
	createdDriver, err := d.driverRepo.CreateDriver(ctx, driver)
	if err != nil {
		return nil, err
	}

	return createdDriver, nil
}

func (d *driverService) GetDriver(ctx context.Context, id string) (*model.Driver, error) {
	getDriver, err := d.driverRepo.GetDriver(ctx, id)
	if err != nil {
		return nil, err
	}

	return getDriver, nil
}

func (d *driverService) GetDrivers(ctx context.Context) (*[]model.Driver, error) {
	drivers, err := d.driverRepo.GetDrivers(ctx)
	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func (d *driverService) DeleteDriver(ctx context.Context, id string) error {
	err := d.driverRepo.DeleteDriver(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
