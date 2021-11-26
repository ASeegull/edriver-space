package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(UpCarsOwners, DownCarsOwners)
}

func UpCarsOwners(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS cars_owners
		(
			user_id INT,
			car_id  INT,
			FOREIGN KEY (user_id) REFERENCES users (id),
			FOREIGN KEY (car_id)  REFERENCES cars (id)
		);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownCarsOwners(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS cars_owners;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
