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

var (
	errMissingPort           = errors.New("PORT is empty")
	errMissingAddress        = errors.New("ADDRESS is empty")
	errMissingJWTTokenSecret = errors.New("JWT_TOKEN_SECRET is empty")
)

func NewConfig() (Config, error) {
	config := Config{}

	config.Address = os.Getenv("ADDRESS")
	if config.Address == "" {
		return Config{}, errMissingAddress
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		return Config{}, errMissingPort
	}

	config.JWTTokenSecret = os.Getenv("JWT_TOKEN_SECRET")
	if config.JWTTokenSecret == "" {
		return Config{}, errMissingJWTTokenSecret
	}
	return config, nil
}
