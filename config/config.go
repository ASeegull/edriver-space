package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	Postgres  PostgresConfig
	Redis     RedisConfig
	Cookie    Cookie
	TokensTTL TokensTTL
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  string
	PostgresqlDriver   string
}

type RedisConfig struct {
	RedisHost string
	Password  string
	DB        int
}

type Cookie struct {
	Name     string
	MaxAge   int
	Path     string
	Secure   bool
	HTTPOnly bool
}

type TokensTTL struct {
	Access  int
	Refresh int
}

//LoadConfig - Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

//ParseConfig - Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	// get from env
	c.Redis.RedisHost = v.Get("REDIS_HOST").(string)
	c.Postgres.PostgresqlHost = v.Get("POSTGRES_HOST").(string)

	return &c, nil
}
