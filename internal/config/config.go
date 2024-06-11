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
	jwtTokenSecret := os.Getenv("JWT_TOKEN_SECRET")
	if jwtTokenSecret == "" {
		return Config{}, errors.New("JWT_TOKEN_SECRET is empty")
	}

	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		return Config{}, errors.New("REST_PORT is empty")
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		return Config{}, errors.New("GRPC_PORT is empty")
	}

	address := os.Getenv("ADDRESS")
	if address == "" {
		return Config{}, errors.New("ADDRESS is empty")
	}

	return Config{jwtTokenSecret, restPort, grpcPort, address}, nil
}
