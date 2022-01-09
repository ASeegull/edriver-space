package repository

import (
	"context"
	"database/sql"

	"github.com/ASeegull/edriver-space/model"
)

// DriversRepos struct for queries from Driver model.
type DriversRepos struct {
	db *sql.DB
}

func NewDriversRepos(db *sql.DB) *DriversRepos {
	return &DriversRepos{
		db: db,
	}
}

func (d *DriversRepos) AddDriver(ctx context.Context, driver model.Driver) error {
	query := `INSERT INTO drivers(
		id,
		license_number,
		date_of_issue,
		expire_date,
		individual_tax_number,
		category,
		category_issuing_date,
		category_expire,
		full_name,
		date_of_birth,
		place_of_birth,
		restrictions
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := d.db.ExecContext(ctx, query,
		driver.ID,
		driver.LicenseNumber,
		driver.DateOfIssue,
		driver.ExpireDate,
		driver.IndividualTaxNumber,
		driver.Category,
		driver.CategoryIssuingDate,
		driver.CategoryExpire,
		driver.FullName,
		driver.DateOfBirth,
		driver.PlaceOfBirth,
		driver.Restrictions,
	)
	if err != nil {
		return err
	}
	return nil

}

func (d *DriversRepos) GetDriver(ctx context.Context, id string) (*model.Driver, error) {
	// Define driver variable.
	driver := model.Driver{}

	// Send query to database.
	query := `SELECT * FROM drivers WHERE id = $1`

	err := d.db.QueryRowContext(ctx, query, id).Scan(
		&driver.ID,
		&driver.LicenseNumber,
		&driver.DateOfIssue,
		&driver.ExpireDate,
		&driver.IndividualTaxNumber,
		&driver.Category,
		&driver.CategoryIssuingDate,
		&driver.CategoryExpire,
		&driver.FullName,
		&driver.DateOfBirth,
		&driver.PlaceOfBirth,
		&driver.Restrictions)
	if err != nil {
		// Return empty object and error.
		return nil, err
	}
	return &model.Driver{}, nil
}

func (d DriversRepos) GetDrivers(ctx context.Context) ([]model.Driver, error) {
	drivers := make([]model.Driver, 0) // Store all incoming data

	query := `SELECT * FROM drivers`
	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		driver := model.Driver{}
		err = rows.Scan(
			&driver.ID,
			&driver.LicenseNumber,
			&driver.DateOfIssue,
			&driver.ExpireDate,
			&driver.IndividualTaxNumber,
			&driver.Category,
			&driver.CategoryIssuingDate,
			&driver.CategoryExpire,
			&driver.FullName,
			&driver.DateOfBirth,
			&driver.PlaceOfBirth,
			&driver.Restrictions)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil
}

func (d DriversRepos) DeleteDriver(ctx context.Context, id string) error {
	query := `DELETE FROM drivers WHERE id = $1`
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
