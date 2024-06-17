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
	DB
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var (
	errMissingPort           = errors.New("PORT is empty")
	errMissingAddress        = errors.New("ADDRESS is empty")
	errMissingJWTTokenSecret = errors.New("JWT_TOKEN_SECRET is empty")
)

func NewConfig() (Config, error) {

	errs := make([]error, 0)

	jwtTokenSecret := os.Getenv("JWT_TOKEN_SECRET")
	if jwtTokenSecret == "" {
		errs = append(errs, errors.New("JWT_TOKEN_SECRET is empty"))
	}

	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		errs = append(errs, errors.New("REST_PORT is empty"))
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		errs = append(errs, errors.New("GRPC_PORT is empty"))
	}

	address := os.Getenv("ADDRESS")
	if address == "" {
		errs = append(errs, errors.New("ADDRESS is empty"))
	}

	DbHost := os.Getenv("DB_HOST")
	if DbHost == "" {
		errs = append(errs, errors.New("DB_HOST is empty"))
	}

	DbPort := os.Getenv("DB_PORT")
	if DbPort == "" {
		errs = append(errs, errors.New("DB_PORT is empty"))
	}

	DbUser := os.Getenv("DB_USER")
	if DbUser == "" {
		errs = append(errs, errors.New("DB_USER is empty"))
	}

	DbPassword := os.Getenv("DB_PASSWORD")
	if DbPassword == "" {
		errs = append(errs, errors.New("DB_PASSWORD is empty"))
	}

	DbName := os.Getenv("DB_NAME")
	if DbName == "" {
		errs = append(errs, errors.New("DB_NAME is empty"))
	}

	if err := errors.Join(errs...); err != nil {
		return Config{}, err
	}

	return Config{
		jwtTokenSecret,
		restPort,
		grpcPort,
		address,
		DB{
			DbHost,
			DbPort,
			DbUser,
			DbPassword,
			DbName,
		}}, nil
}
