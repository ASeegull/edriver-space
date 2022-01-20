package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ASeegull/edriver-space/model"
)

type ParkingFinesRepos struct {
	*sql.DB
}

// NewParkingFinesRepos returns pointer to new UploadRepos
func NewParkingFinesRepos(db *sql.DB) *ParkingFinesRepos {
	return &ParkingFinesRepos{db}
}

// AddParkingFine adds new parking fine to the database
func (u *ParkingFinesRepos) AddParkingFine(ctx context.Context, fine model.ParkingFine) error {
	_, err := u.DB.ExecContext(ctx, "INSERT INTO cars_parking_fines(fine_num, issue_time, car_VIN, cost, photo_url) VALUES ($1, $2, $3, $4, $5)",
		fine.FineNum, fine.IssueTime, fine.CarVIN, fine.Cost, fine.PhotoURL)
	if err != nil {
		err = errors.New("error adding parking fine to the database")
		return err
	}
	return nil
}

// GetParkingFine returns parking fine from the database by its id
func (u *ParkingFinesRepos) GetParkingFine(ctx context.Context, id string) (*model.ParkingFine, error) {
	parkingFine := model.ParkingFine{} // Store incoming data

	err := u.DB.QueryRowContext(ctx, "SELECT * FROM cars_parking_fines WHERE id = $1", id).Scan(
		&parkingFine.ID, &parkingFine.FineNum, &parkingFine.IssueTime, &parkingFine.CarVIN, &parkingFine.Cost, &parkingFine.PhotoURL)
	if err != nil {
		err = errors.New("error retrieving parking fine from the database")
		return nil, err
	}
	return &parkingFine, nil
}

// GetParkingFines returns all existing parking fines from the database
func (u *ParkingFinesRepos) GetParkingFines(ctx context.Context) ([]model.ParkingFine, error) {
	parkingFines := make([]model.ParkingFine, 0) // Store all incoming data

	rows, err := u.DB.QueryContext(ctx, "SELECT * FROM cars_parking_fines")
	if err != nil {
		err = errors.New("error retrieving rows from the database")
		return nil, err
	}

	for rows.Next() {
		parkingFine := model.ParkingFine{} // Store data of each parking fine
		err = rows.Scan(&parkingFine.ID, &parkingFine.FineNum, &parkingFine.IssueTime, &parkingFine.CarVIN, &parkingFine.Cost, &parkingFine.PhotoURL)
		if err != nil {
			err = errors.New("error scanning the parking fine data from rows")
			return nil, err
		}
		// Add received parking fine to parking fine slice
		parkingFines = append(parkingFines, parkingFine)
	}

	// Close rows and check errors
	if err = rows.Close(); err != nil {
		err = errors.New("error closing rows")
		return nil, err
	}

	return parkingFines, nil
}

// DeleteParkingFine removes parking fine from the database by its id
func (u *ParkingFinesRepos) DeleteParkingFine(ctx context.Context, id string) error {
	_, err := u.DB.ExecContext(ctx, "DELETE FROM cars_parking_fines WHERE id = $1", id)
	if err != nil {
		err = errors.New("error removing parking fine from the database")
		return err
	}
	return nil
}
