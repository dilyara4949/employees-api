package config

import (
	"os"
)

type Config struct {
	JWTTokenSecret string
	Port           string
}

func NewConfig() Config {
	config := Config{}

	config.JWTTokenSecret = os.Getenv("JWTTokenSecret")
	config.Port = os.Getenv("Port")
	return config
}
