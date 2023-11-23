package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Config struct {
	AppConfig AppConfig
	DBConfig  DBConfig
	AuthConfig AuthConfig
}

func GetConfig() (Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Println(err)
		return Config{}, err
	}
	return cfg, nil
}

