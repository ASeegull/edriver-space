package repository

import (
	"context"
	"database/sql"
	"github.com/ASeegull/edriver-space/model"
)

type UploadRepos struct {
	*sql.DB
}

// NewUploadRepos returns pointer to new UploadRepos
func NewUploadRepos(db *sql.DB) *UploadRepos {
	return &UploadRepos{db}
}

func (u *UploadRepos) AddFine(ctx context.Context, fine model.ParkingFine) error {
	//TODO Add fine to the database
	panic("implement me")
}

func (u *UploadRepos) GetFine(ctx context.Context, id string) (*model.ParkingFine, error) {
	//TODO Get Parking fine from the database
	panic("implement me")
}

func (u *UploadRepos) GetFines(ctx context.Context) ([]model.ParkingFine, error) {
	//TODO Get all Parking fines from the database
	panic("implement me")
}

func (u *UploadRepos) DeleteFine(ctx context.Context, id string) error {
	//TODO delete parking fine from the database
	panic("implement me")
}
