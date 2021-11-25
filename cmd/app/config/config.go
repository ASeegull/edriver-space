package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Config struct stores all configuration values for project using Viper
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

//LoadConfig reads configuration from .env file
func LoadConfig(path string) (config Config, err error) {

	//Setting default path for config file
	if path == "" {
		path = "./config"
	}

	//Declareting path and type for config file
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	//Parsing config vals from file (first step)
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	//Parsing config vals from file (second step)
	err = viper.Unmarshal(&config)

	return

}

// GetConfigString returns specific value from config file
func GetConfigString(ValName string) (val string, err error) {

	//Loading config
	config, _ := LoadConfig("./config")

	switch ValName {
	case "DBDriver":
		val = config.DBDriver
	case "DBSource":
		val = config.DBSource
	case "ServerAddress":
		val = config.ServerAddress
	default:
		err = fmt.Errorf("Cannot find value " + ValName)
	}

	return

}
