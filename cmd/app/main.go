package main

import (
	"fmt"

	"github.com/ASeegull/edriver-space/api/server"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")

	//Initializing Logger
	logger.LogInit()

	// Loading config values
	conf, err := config.LoadConfig("")
	if err != nil {
		fmt.Println(err)
	}

	//Creating and starting server
	s := server.NewServer()
	logger.LogFatal(s.Start(":" + conf.ServerPort))

}
