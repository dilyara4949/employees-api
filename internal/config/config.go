package config

import (
	"errors"
	"os"
	"strconv"
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
	Timeout  int
	MaxConn  int
}

var (
	errMissingRestPort       = errors.New("REST_PORT is empty")
	errMissingAddress        = errors.New("ADDRESS is empty")
	errMissingJWTTokenSecret = errors.New("JWT_TOKEN_SECRET is empty")
	errMissingGrpcPort       = errors.New("GRPC_PORT is empty")
	errMissingDbHost         = errors.New("DB_HOST is empty")
	errMissingDbPort         = errors.New("DB_PORT is empty")
	errMissingDbName         = errors.New("DB_USER is empty")
	errMissingDbUser         = errors.New("DB_USER is empty")
	errMissingDbPassword     = errors.New("DB_PASSWORD is empty")
	errMissingDbTimeout      = errors.New("DB_TIMEOUT is empty")
	errMissingDbMaxConn      = errors.New("DB_MAX_CONNECTIONS is empty")
	errMaxConnType           = errors.New("DB_MAX_CONNECTIONS must be an integer")
	errDbTimeoutType         = errors.New("DB_TIMEOUT must be an integer")
)

func NewConfig() (Config, error) {

	errs := make([]error, 0)

	jwtTokenSecret := os.Getenv("JWT_TOKEN_SECRET")
	if jwtTokenSecret == "" {
		errs = append(errs, errMissingJWTTokenSecret)
	}

	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		errs = append(errs, errMissingRestPort)
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		errs = append(errs, errMissingGrpcPort)
	}

	address := os.Getenv("ADDRESS")
	if address == "" {
		errs = append(errs, errMissingAddress)
	}

	DbHost := os.Getenv("DB_HOST")
	if DbHost == "" {
		errs = append(errs, errMissingDbHost)
	}

	DbPort := os.Getenv("DB_PORT")
	if DbPort == "" {
		errs = append(errs, errMissingDbPort)
	}

	DbUser := os.Getenv("DB_USER")
	if DbUser == "" {
		errs = append(errs, errMissingDbUser)
	}

	DbPassword := os.Getenv("DB_PASSWORD")
	if DbPassword == "" {
		errs = append(errs, errMissingDbPassword)
	}

	DbName := os.Getenv("DB_NAME")
	if DbName == "" {
		errs = append(errs, errMissingDbName)
	}

	DbTimeoutStr := os.Getenv("DB_TIMEOUT")
	if DbTimeoutStr == "" {
		errs = append(errs, errMissingDbTimeout)
	}

	DbTimeout, err := strconv.Atoi(DbTimeoutStr)
	if err != nil {
		errs = append(errs, errDbTimeoutType)
	}

	DbMaxConnStr := os.Getenv("DB_MAX_CONNECTIONS")
	if DbMaxConnStr == "" {
		errs = append(errs, errMissingDbMaxConn)
	}

	DbMaxconn, err := strconv.Atoi(DbMaxConnStr)
	if err != nil {
		errs = append(errs, errMaxConnType)
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
			DbTimeout,
			DbMaxconn,
		}}, nil
}
