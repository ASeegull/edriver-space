package main

import (
	"fmt"

	"github.com/ASeegull/edriver-space/api/server"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/ASeegull/edriver-space/storage"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")
	
	//Initializing Logger
	logger.LogInit()

	// Loading config values
	conf, err := config.LoadConfig("")
	
	// Path to env
	if err := godotenv.Load("env/docker/postgres.env", "env/docker/app.env"); err != nil {
		logger.LogFatal(fmt.Errorf("error loading env variables: %s", err.Error()))
	}

	if err != nil {
		logger.LogErr(err)
	}

	//Verify if connection is ok
	storage.InitConnection()

	//Creating and starting server
	s := server.NewServer()
	logger.LogFatal(s.Start(":" + conf.ServerPort))

}
