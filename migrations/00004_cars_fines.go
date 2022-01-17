package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(UpCarsFines, DownCarsFines)
}

func UpCarsFines(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS cars_fines
		(
		    id							 SERIAL,
			vehicle_registration_number  VARCHAR(25)  NOT NULL,
            date_and_time                DATE         NOT NULL,
            place                        VARCHAR(255) NOT NULL,
            file_law_article             VARCHAR(30)  NOT NULL,
            price                        INT          NOT NULL,
            info                         VARCHAR(255),
            img_url                      VARCHAR(255),
            FOREIGN KEY (vehicle_registration_number) REFERENCES cars (registration_number)
		);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownCarsFines(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS cars_fines;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
