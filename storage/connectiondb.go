package storage

import (
	"fmt"
	"os"
	"sync"
	
	"github.com/ASeegull/edriver-space/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db   *sqlx.DB
	once sync.Once
)

// Singleton
// In software engineering, the singleton pattern is a software design pattern that restricts the instantiation of a class to one object.
// This is useful when exactly one object is needed to coordinate actions across the system.
// https://en.wikipedia.org/wiki/Singleton_pattern
// http://marcio.io/2015/07/singleton-pattern-in-go/

// Set database config
// export PGUSER=postgres
// export PGDB=postgres
// export PGHOST=localhost
// export PGPORT=5432
// Run postgresql inside a container
// docker run -d -p 5432:5432 postgres:latest

// MustGetConnection returns database connection
func MustGetConnection() *sqlx.DB {
	once.Do(func() {
		pguser := os.Getenv("PGUSER")
		pgdb := os.Getenv("PGDB")
		pghost := os.Getenv("PGHOST")
		pgport := os.Getenv("PGPORT")
		pgpass := os.Getenv("PGPASS")
		dbURI := fmt.Sprintf("user=%s dbname=%s host=%s port=%v sslmode=disable", pguser, pgdb, pghost, pgport)
		if pgpass != "" {
			dbURI += " password=" + pgpass
		}
		var err error
		db, err = sqlx.Connect("postgres", dbURI)
		if err != nil {
			logger.LogMsg(fmt.Sprintf("Unable to connection to database: %v\n", err),"Panic")
		}
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
	})
	return db
}

func InitConnection() {
	
	conn := MustGetConnection()
	conErr := conn.Ping()
	
	if conErr != nil {
		logger.LogFatal(conErr)
	}
	fmt.Println("Successfully connected ✓")
	
	defer func() {
		if conErr := conn.Close(); conErr != nil {
			fmt.Println("db connection closed.")
		}
	}()
}

