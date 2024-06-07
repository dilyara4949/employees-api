package config

import (
	"errors"
	"os"
)

type Config struct {
	JWTTokenSecret string
	RestPort       string
	GrpcPort       string
	Address        string
}

func NewConfig() (Config, error) {
	config := Config{}

	config.JWTTokenSecret = os.Getenv("JWT_TOKEN_SECRET")
	if config.JWTTokenSecret == "" {
		return Config{}, errors.New("JWT_TOKEN_SECRET is empty")
	}

	config.RestPort = os.Getenv("REST_PORT")
	if config.RestPort == "" {
		return Config{}, errors.New("REST_PORT is empty")
	}

	config.GrpcPort = os.Getenv("GRPC_PORT")
	if config.GrpcPort == "" {
		return Config{}, errors.New("GRPC_PORT is empty")
	}

	config.Address = os.Getenv("ADDRESS")
	if config.Address == "" {
		return Config{}, errors.New("ADDRESS is empty")
	}
	return config, nil
}
