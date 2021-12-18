package main

import (
	"fmt"
	"github.com/ASeegull/edriver-space/api/server"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	_ "github.com/ASeegull/edriver-space/migrations"
	"github.com/ASeegull/edriver-space/storage"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"os"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")

	// Initializing Logger
	logger.LogInit()

	// Loading config values
	conf, err := config.LoadConfig("")

	if err != nil {
		logger.LogErr(err)
	}

	// Path to env
	if err = godotenv.Load("env/docker/postgres.env", "env/docker/app.env"); err != nil {
		logger.LogFatal(fmt.Errorf("error loading env variables: %s", err.Error()))
	}

	// Open db connection
	conn := storage.MustGetConnection()

	if conErr := conn.Ping(); conErr != nil {
		logger.LogFatal(conErr)
	}
	fmt.Println("Successfully connected ✓")

	defer func() {
		if conErr := conn.Close(); conErr != nil {
			logger.LogFatal(err)
		}
	}()

	if err = goose.SetDialect(os.Getenv("DB_DRIVER")); err != nil {
		logger.LogFatal(err)
	}
	// Up new migrations
	if err = goose.Up(conn, "."); err != nil {
		logger.LogFatal(err)
	}
	// Creating and starting server
	s := server.NewServer()
	logger.LogFatal(s.Start(":" + conf.ServerPort))
}
