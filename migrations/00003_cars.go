package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(UpCars, DownCars)
}

func UpCars(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS cars
		(
			id                                     SERIAL PRIMARY KEY,
            mark                                   VARCHAR(55)        NOT NULL,
            type                                   VARCHAR(55)        NOT NULL,
            VIN_code                               VARCHAR(55) UNIQUE NOT NULL,
            maximum_mass                           INT                NOT NULL,
            vehicle_category                       VARCHAR(10)        NOT NULL,
            colour_of_the_vehicle                  VARCHAR(55)        NOT NULL,
            number_of_seats_including_drivers_seat INT                NOT NULL,
            registration_number                    VARCHAR(25) UNIQUE NOT NULL,
            full_name                              VARCHAR(255)       NOT NULL,
            ownership                              VARCHAR(255)       NOT NULL,
            period_of_validity                     DATE,
            date_of_registration                   DATE               NOT NULL
		);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownCars(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS cars;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
