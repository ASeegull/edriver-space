package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Cookie   CookieConfig
	Token    TokenConfig
}

type ServerConfig struct {
	Port         string `envconfig:"port"`
	DBDriver     string `envconfig:"db_driver"`
	JWTSecretKey string `envconfig:"jwt_secret_key"`
	HashSalt     string `envconfig:"hash_salt"`
}

type PostgresConfig struct {
	Host     string `envconfig:"host"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	DB       string `envconfig:"db"`
	SSLMode  string `envconfig:"sslmode"`
	Driver   string `envconfig:"driver"`
}

type RedisConfig struct {
	Host     string `envconfig:"host"`
	Password string `envconfig:"password"`
	DB       int    `envconfig:"db"`
}

type CookieConfig struct {
	Name     string `envconfig:"name"`
	MaxAge   int    `envconfig:"max_age"`
	Path     string `envconfig:"path"`
	Secure   bool   `envconfig:"secure"`
	HTTPOnly bool   `envconfig:"http_only"`
}

type TokenConfig struct {
	AccessTTL  int `envconfig:"access_ttl"`
	RefreshTTL int `envconfig:"refresh_ttl"`
}

const (
	appGroup      = "server"
	cookieGroup   = "cookie"
	postgresGroup = "postgres"
	redisGroup    = "redis"
	tokenGroup    = "token"
)

func CreateConfig() (*Config, error) {
	config := new(Config)

	if err := envconfig.Process(appGroup, &config.Server); err != nil {
		return nil, err
	}
	if err := envconfig.Process(cookieGroup, &config.Cookie); err != nil {
		return nil, err
	}
	if err := envconfig.Process(postgresGroup, &config.Postgres); err != nil {
		return nil, err
	}
	if err := envconfig.Process(redisGroup, &config.Redis); err != nil {
		return nil, err
	}
	if err := envconfig.Process(tokenGroup, &config.Token); err != nil {
		return nil, err
	}
	return config, nil
}
