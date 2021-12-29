package storage

import (
	"github.com/ASeegull/edriver-space/config"
)

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s",
		cfg.Postgres.PostgresqlDriver,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlSSLMode)

	db, err := goose.OpenDBWithDriver(cfg.Postgres.PostgresqlDriver, dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
