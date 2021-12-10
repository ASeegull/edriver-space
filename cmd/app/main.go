package main

import (
	"fmt"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/internal/server"
	_ "github.com/ASeegull/edriver-space/migrations"
	"github.com/ASeegull/edriver-space/pkg/db/postrges"
	"github.com/ASeegull/edriver-space/pkg/db/redis"
	"log"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")

	cfgFile, err := config.LoadConfig("./config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	db, err := postrges.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	//if err := goose.Up(db, "."); err != nil {
	//	logger.LogFatal(err)
	//}

	redisClient := redis.NewRedisClient(cfg)

	s := server.NewServer(db, redisClient, cfg)
	if err = s.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
