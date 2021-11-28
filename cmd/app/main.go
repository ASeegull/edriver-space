package main

import (
	"fmt"

	"github.com/ASeegull/edriver-space/api/server"
	"github.com/ASeegull/edriver-space/config"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")

	// Loading config values
	conf, _ := config.LoadConfig("")

	//Creating and starting server
	s := server.NewServer()
	s.Logger.Fatal(s.Start(":" + conf.ServerPort))
}
