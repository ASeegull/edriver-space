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

	mp := make(map[string]string)
	mp["Info"] = "First Info"
	mp["Warn2"] = "First Warning"
	mp2 := make(map[string]string)

	err := logger.LogMsgWithFields("Firts Message", "Warn", mp)
	fmt.Println(err)

	err = logger.LogMsgWithFields("Firts Message", "Trace", mp)
	fmt.Println(err)

	err = logger.LogMsgWithFields("Firts Message", "", mp)
	fmt.Println(err)

	err = logger.LogMsgWithFields("Firts Message", "fd", mp2)
	fmt.Println(err)

	err = logger.LogMsg("2 Message", "Warn")
	fmt.Println(err)

	err = logger.LogMsg("3 Message", "fd")
	fmt.Println(err)

	// Loading config values
	conf, _ := config.LoadConfig("")

	//Creating and starting server
	s := server.NewServer()
	s.Logger.Fatal(s.Start(":" + conf.ServerPort))

}
