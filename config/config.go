package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	ServiceName  string `env:"SERVICE_NAME"`
	ServicePort  string `env:"SERVICE_PORT"`
	ServiceHost  string `env:"SERVICE_HOST"`
	FileLocation string `env:"FILE_LOCATION"`
	FileServer   string `env:"FILE_SERVER_URL"`
}

var Config ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Warn().Msg("Cannot find .env file. OS Environments will be used")
	}
	err = env.Parse(&Config)

	return err
}
