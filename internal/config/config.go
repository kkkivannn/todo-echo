package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     int    `env:"PORT"`
	Host     string `env:"HOST"`
	DBString string `env:"DB_CONNECTION_STRING"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config := ReadEnv()

	return &config
}

func ReadEnv() Config {
	var config Config
	if err := cleanenv.ReadEnv(&config); err != nil {
		panic(err)
	}

	return config
}
