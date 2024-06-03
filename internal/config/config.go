package config

import (
	"os"
)

type Config struct {
	JWTTokenSecret string
}

func NewConfig() Config {
	config := Config{}

	config.JWTTokenSecret = os.Getenv("JWTTokenSecret")
	return config
}
