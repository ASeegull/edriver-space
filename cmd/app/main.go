package main

import (
	"fmt"
	"log"

	"github.com/ASeegull/edriver-space/api/server"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")
	s := server.NewServer()
	s.Logger.Fatal(s.Start(":1323"))

	// Verify if connection is ok
	conn := MustGetConnection()
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected ✓")
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
