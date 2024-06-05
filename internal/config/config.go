package config

import (
	"errors"
	"os"
)

type Config struct {
	JWTTokenSecret string
	Port           string
	Address        string
}

func NewConfig() (Config, error) {
	config := Config{}

	config.Address = os.Getenv("ADDRESS")
	if config.Address == "" {
		return Config{}, errors.New("ADDRESS is empty")
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		return Config{}, errors.New("PORT is empty")
	}

	config.JWTTokenSecret = os.Getenv("JWT_TOKEN_SECRET")
	if config.JWTTokenSecret == "" {
		return Config{}, errors.New("JWT_TOKEN_SECRET is empty")
	}
	return config, nil
}
