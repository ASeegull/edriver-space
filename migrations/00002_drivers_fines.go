package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(UpUsersFines, DownUsersFines)
}

func UpUsersFines(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS drivers_fines
		(
		    id							 SERIAL,
			licence_number               VARCHAR(55)  NOT NULL,
    		fine_num 			         VARCHAR(55) UNIQUE NOT NULL,
			date_and_time                DATE         NOT NULL,
			place                        VARCHAR(255) NOT NULL,
			file_law_article             VARCHAR(30)  NOT NULL,
			price                        INT          NOT NULL,
			vehicle_registration_number  VARCHAR(25)  NOT NULL,
			FOREIGN KEY (licence_number) REFERENCES drivers (licence_number)
		);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownUsersFines(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS drivers_fines;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
