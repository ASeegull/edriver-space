package main

import (
	"fmt"
	"github.com/ASeegull/edriver-space/api/server"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")
	s := server.NewServer()
	s.Logger.Fatal(s.Start(":1323"))
}
