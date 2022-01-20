package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ASeegull/edriver-space/model"
)

type DriverFinesRep struct {
	*sql.DB
}

// NewDriverFinesRep returns pointer to new DriverFinesRep
func NewDriverFinesRep(db *sql.DB) *DriverFinesRep {
	return &DriverFinesRep{db}
}

// GetDriverFines returns all driver fines with the given licence
func (dfr *DriverFinesRep) GetDriverFines(ctx context.Context, licence string) ([]model.DriversFine, error) {
	driverFines := make([]model.DriversFine, 0) // Store all incoming data

	rows, err := dfr.DB.QueryContext(ctx, "SELECT * FROM drivers_fines WHERE licence_number = $1", licence)
	if err != nil {
		err = errors.New("error retrieving rows from the database")
		return nil, err
	}

	for rows.Next() {
		driverFine := model.DriversFine{} // Store data of each driver fine
		err = rows.Scan(&driverFine.Id, &driverFine.LicenceNumber, &driverFine.FineNum, &driverFine.DataAndTime,
			&driverFine.Place, &driverFine.FileLawArticle, &driverFine.Price, &driverFine.VehicleRegistrationNumber)
		if err != nil {
			err = errors.New("error retrieving driver fine from the database")
			return nil, err
		}
		// Add received driver fine to car fine slice
		driverFines = append(driverFines, driverFine)
	}

	// Close rows and check errors
	if err = rows.Close(); err != nil {
		err = errors.New("error closing rows")
		return nil, err
	}

	return driverFines, nil
}

// GetDriverFine returns driver fine from the database by driver licence
func (dfr *DriverFinesRep) GetDriverFine(ctx context.Context, licence string) (*model.DriversFine, error) {
	driverFine := model.DriversFine{} // Store incoming data

	err := dfr.DB.QueryRowContext(ctx, "SELECT * FROM drivers_fines WHERE licence_number = $1", licence).Scan(
		&driverFine.Id, &driverFine.LicenceNumber, &driverFine.FineNum, &driverFine.DataAndTime,
		&driverFine.Place, &driverFine.FileLawArticle, &driverFine.Price, &driverFine.VehicleRegistrationNumber)
	if err != nil {
		err = errors.New("error retrieving driver fine from the database")
		return nil, err
	}
	return &driverFine, nil
}

// DeleteDriverFine removes driver fine with given fine number
func (dfr *DriverFinesRep) DeleteDriverFine(ctx context.Context, fineNum string) error {
	_, err := dfr.DB.ExecContext(ctx, "DELETE FROM drivers_fines WHERE fine_num = $1", fineNum)
	if err != nil {
		err = errors.New("error removing driver fine from the database")
		return err
	}
	return nil
}
