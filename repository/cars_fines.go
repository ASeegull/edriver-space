package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ASeegull/edriver-space/model"
)

type CarFinesRep struct {
	*sql.DB
}

// NewCarFinesRep returns pointer to new CarFinesRep
func NewCarFinesRep(db *sql.DB) *CarFinesRep {
	return &CarFinesRep{db}
}

// GetCarFines returns all car fines with the given registration number
func (cfr *CarFinesRep) GetCarFines(ctx context.Context, regNum string) ([]model.CarsFine, error) {
	carFines := make([]model.CarsFine, 0) // Store all incoming data

	rows, err := cfr.DB.QueryContext(ctx, "SELECT * FROM cars_fines WHERE vehicle_registration_number = $1", regNum)
	if err != nil {
		err = errors.New("error retrieving rows from the database")
		return nil, err
	}

	for rows.Next() {
		carFine := model.CarsFine{} // Store data of each car fine
		err = rows.Scan(&carFine.Id, &carFine.VehicleRegistrationNumber, &carFine.FineNum, &carFine.DataAndTime, &carFine.Place,
			&carFine.FileLawArticle, &carFine.Price, &carFine.Info, &carFine.ImdUrl)
		if err != nil {
			err = errors.New("error scanning car fine data from the rows")
			return nil, err
		}
		// Add received car fine to car fine slice
		carFines = append(carFines, carFine)
	}

	// Close rows and check errors
	if err = rows.Close(); err != nil {
		err = errors.New("error closing rows")
		return nil, err
	}

	return carFines, nil
}

// GetCarFine returns car fine from the database by car registration number
func (cfr *CarFinesRep) GetCarFine(ctx context.Context, regNum string) (*model.CarsFine, error) {
	carFine := model.CarsFine{} // Store incoming data

	err := cfr.DB.QueryRowContext(ctx, "SELECT * FROM cars_fines WHERE vehicle_registration_number = $1", regNum).Scan(
		&carFine.Id, &carFine.VehicleRegistrationNumber, &carFine.FineNum, &carFine.DataAndTime, &carFine.Place,
		&carFine.FileLawArticle, &carFine.Price, &carFine.Info, &carFine.ImdUrl)
	if err != nil {
		err = errors.New("error retrieving car fine from the database")
		return nil, err
	}
	return &carFine, nil
}

// AddCarFine adds new car fine to the database
func (cfr *CarFinesRep) AddCarFine(ctx context.Context, fine *model.CarsFine) error {
	query := `INSERT INTO cars_fines(
                       vehicle_registration_number,
                       fine_num,
                       date_and_time,
                       place,
                       file_law_article,
                       price,
                       info,
                       img_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := cfr.DB.ExecContext(ctx, query,
		fine.VehicleRegistrationNumber, fine.FineNum, fine.DataAndTime, fine.Place, fine.FileLawArticle, fine.Price, fine.Info, fine.ImdUrl)
	if err != nil {
		err = errors.New("error adding car fine to the database")
		return err
	}
	return nil
}

// DeleteCarFine removes car fine with the given fine number
func (cfr *CarFinesRep) DeleteCarFine(ctx context.Context, fineNum string) error {
	_, err := cfr.DB.ExecContext(ctx, "DELETE FROM cars_fines WHERE fine_num = $1", fineNum)
	if err != nil {
		err = errors.New("error removing car fine from the database")
		return err
	}
	return nil
}
