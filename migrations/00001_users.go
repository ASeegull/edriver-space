package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(UpUsers, DownUsers)
}

func UpUsers(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS users
		(
			id                    SERIAL PRIMARY KEY,
			full_name             VARCHAR(255)       NOT NULL,
			date_of_birth         DATE               NOT NULL,
			place_of_birth        VARCHAR(255)       NOT NULL,
			date_of_issue         DATE               NOT NULL,
			expire_date           DATE               NOT NULL,
			licence_number        VARCHAR(55) UNIQUE NOT NULL,
			category              VARCHAR(25)        NOT NULL,
			category_issuing_date DATE               NOT NULL,
			category_expire       DATE               NOT NULL,
			restrictions          VARCHAR(255) DEFAULT '-',
			individual_tax_number VARCHAR(55) UNIQUE NOT NULL
		);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DownUsers(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS users;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
