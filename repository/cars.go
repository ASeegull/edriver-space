package repository

import (
	"context"
	"database/sql"

	"github.com/ASeegull/edriver-space/model"
)

// CarsRepos struct for queries from Car model.
type CarsRepos struct {
	db *sql.DB
}

func NewCarsRepos(db *sql.DB) *CarsRepos {
	return &CarsRepos{
		db: db,
	}
}

func (c *CarsRepos) CreateCar(ctx context.Context, car *model.Car) (*model.Car, error) {
	query := `INSERT INTO cars(
		name, 
		VIN_code, 
		registration_number, 
		vehicle_category, 
		make,
		type, 
		commercial_description,
		maximum_mass, 
		mass_of_the_vehicle_in_service, 
		capacity, 
		colour_of_the_vehicle, 
		number_of_seats_including_drivers_seat, 
		date_of_first_registration, 
		period_of_validity, 
		date_of_registration,
		full_name, 
		address, 
		ownership) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, 
			$12, $13, $14, $15, $16, $17, $18)`
	_, err := c.db.ExecContext(ctx, query,
		car.ID,
		car.Name,
		car.VIN,
		car.RegistrationNum,
		car.VehicleCategory,
		car.Make,
		car.Type,
		car.CommercialDescription,
		car.MaxMass,
		car.ServiceMass,
		car.Capacity,
		car.Colour,
		car.SeatsNum,
		car.FirstRegDate,
		car.ValidityPeriod,
		car.RegistrationDate,
		car.FullName,
		car.Address,
		car.Ownership)
	if err != nil {
		return nil, err
	}
	return car, nil

}

func (c *CarsRepos) GetCar(ctx context.Context, id string) (*model.Car, error) {
	// Define car variable.
	car := model.Car{}

	// Send query to database.
	query := `SELECT * FROM cars WHERE id = $1`

	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&car.ID,
		&car.Name,
		&car.VIN,
		&car.RegistrationNum,
		&car.VehicleCategory,
		&car.Make,
		&car.Type,
		&car.CommercialDescription,
		&car.MaxMass,
		&car.ServiceMass,
		&car.Capacity,
		&car.Colour,
		&car.SeatsNum,
		&car.FirstRegDate,
		&car.ValidityPeriod,
		&car.RegistrationDate,
		&car.FullName,
		&car.Address,
		&car.Ownership)
	if err != nil {
		// Return empty object and error.
		return nil, err
	}
	return &model.Car{}, nil
}

func (c CarsRepos) GetCars(ctx context.Context) (*[]model.Car, error) {
	cars := make([]model.Car, 0) // Store all incoming data

	query := `SELECT * FROM cars`
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		car := model.Car{}
		err = rows.Scan(
			&car.ID,
			&car.Name,
			&car.VIN,
			&car.RegistrationNum,
			&car.VehicleCategory,
			&car.Make,
			&car.Type,
			&car.CommercialDescription,
			&car.MaxMass,
			&car.ServiceMass,
			&car.Capacity,
			&car.Colour,
			&car.SeatsNum,
			&car.FirstRegDate,
			&car.ValidityPeriod,
			&car.RegistrationDate,
			&car.FullName,
			&car.Address,
			&car.Ownership)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &cars, nil
}

func (c CarsRepos) DeleteCar(ctx context.Context, id string) error {
	query := `DELETE FROM cars WHERE id = $1`
	_, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
