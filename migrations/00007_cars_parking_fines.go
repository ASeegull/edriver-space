package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(UpCarsParkingFines, DownCarsParkingFines)
}

func UpCarsParkingFines(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS cars_parking_fines
		(
			id 			SERIAL PRIMARY KEY,
			fine_num 	VARCHAR(55) UNIQUE NOT NULL,
    		issue_time 	VARCHAR(55) NOT NULL,
    		car_VIN 	VARCHAR(55) UNIQUE NOT NULL,
    		cost 		INT NOT NULL,
    		photo_url 	VARCHAR(255), 
    		FOREIGN KEY (car_VIN) REFERENCES cars (VIN_code)
		);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownCarsParkingFines(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS cars_parking_fines;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
