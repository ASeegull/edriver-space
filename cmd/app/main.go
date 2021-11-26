package main

import (
	"fmt"
	"log"

	"github.com/ASeegull/edriver-space/api/server"
	"github.com/joho/godotenv"
)

func main() {

	// Path to env
	if err := godotenv.Load("env/docker/postgres.env", "env/docker/app.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	fmt.Println("Hello Lv-644.Go!")
	s := server.NewServer()
<<<<<<< Updated upstream
	s.Logger.Fatal(s.Start(":1323"))
=======
	s.Logger.Fatal(s.Start(":1323")) // Move to env .. hardcode port repot
>>>>>>> Stashed changes

	// Verify if connection is ok
	conn := MustGetConnection()
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected ✓")
<<<<<<< Updated upstream
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
=======
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("db connection closed.")
		}
	}()
>>>>>>> Stashed changes
}
