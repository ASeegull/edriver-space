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
			id                    	SERIAL 			PRIMARY KEY,
			email            		VARCHAR(255) 	NOT NULL,
			password         		VARCHAR(72)   	NOT NULL,
			role 					VARCHAR(30) 	DEFAULT 'user',
			driver_licence_number   VARCHAR(55) 	DEFAULT ''
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
