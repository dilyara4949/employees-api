package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	JWTTokenSecret string
	RestPort       string
	GrpcPort       string
	Address        string
	DatabaseType   string
	PostgresConfig
	MongoConfig
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Timeout  int
}

type PostgresConfig struct {
	DB
	MaxConn int
}

type MongoConfig struct {
	DB
	Collections
}

type Collections struct {
	Positions string
	Employees string
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

const (
	positionsCollection = "positions"
	employeesCollection = "employees"
	PostgresDB          = "postgres"
	MongoDB             = "mongo"
	defaultDB           = PostgresDB
	jwtTokenSecretEnv   = "JWT_TOKEN_SECRET"
	restPortEnv         = "REST_PORT"
	grpcPortEnv         = "GRPC_PORT"
	addressEnv          = "ADDRESS"
	databaseTypeEnv     = "DATABASE_TYPE"
	dbHostEnv           = "DB_HOST"
	dbPortEnv           = "DB_PORT"
	dbUserEnv           = "DB_USER"
	dbPasswordEnv       = "DB_PASSWORD"
	dbNameEnv           = "DB_NAME"
	dbTimeoutEnv        = "DB_TIMEOUT"
	dbMaxConnEnv        = "DB_MAX_CONNECTIONS"
	posCollectionEnv    = "POSITIONS_COLLECTION"
	empCollectionEnv    = "EMPLOYEES_COLLECTION"
)

func NewConfig() (Config, error) {

	errs := make([]error, 0)

	jwtTokenSecret := os.Getenv(jwtTokenSecretEnv)
	if jwtTokenSecret == "" {
		errs = append(errs, errMissingJWTTokenSecret)
	}

	restPort := os.Getenv(restPortEnv)
	if restPort == "" {
		errs = append(errs, errMissingRestPort)
	}

	grpcPort := os.Getenv(grpcPortEnv)
	if grpcPort == "" {
		errs = append(errs, errMissingGrpcPort)
	}

	address := os.Getenv(addressEnv)
	if address == "" {
		errs = append(errs, errMissingAddress)
	}

	DbHost := os.Getenv(dbHostEnv)
	if DbHost == "" {
		errs = append(errs, errMissingDbHost)
	}

	DbPort := os.Getenv(dbPortEnv)
	if DbPort == "" {
		errs = append(errs, errMissingDbPort)
	}

	DbUser := os.Getenv(dbUserEnv)
	if DbUser == "" {
		errs = append(errs, errMissingDbUser)
	}

	DbPassword := os.Getenv(dbPasswordEnv)
	if DbPassword == "" {
		errs = append(errs, errMissingDbPassword)
	}

	DbName := os.Getenv(dbNameEnv)
	if DbName == "" {
		errs = append(errs, errMissingDbName)
	}

	DbTimeoutStr := os.Getenv(dbTimeoutEnv)
	if DbTimeoutStr == "" {
		errs = append(errs, errMissingDbTimeout)
	}

	DbTimeout, err := strconv.Atoi(DbTimeoutStr)
	if err != nil {
		errs = append(errs, errDbTimeoutType)
	}

	DbMaxConnStr := os.Getenv(dbMaxConnEnv)
	if DbMaxConnStr == "" {
		errs = append(errs, errMissingDbMaxConn)
	}

	DbMaxconn, err := strconv.Atoi(DbMaxConnStr)
	if err != nil {
		errs = append(errs, errMaxConnType)
	}

	posCollection := os.Getenv(posCollectionEnv)
	if posCollection == "" {
		posCollection = positionsCollection
	}

	empCollection := os.Getenv(empCollectionEnv)
	if empCollection == "" {
		empCollection = employeesCollection
	}

	dbType := strings.ToLower(os.Getenv(databaseTypeEnv))
	if dbType == "" {
		dbType = defaultDB
	}

	if err := errors.Join(errs...); err != nil {
		return Config{}, err
	}

	cfg := Config{
		JWTTokenSecret: jwtTokenSecret,
		RestPort:       restPort,
		GrpcPort:       grpcPort,
		Address:        address,
		DatabaseType:   dbType,
	}

	switch dbType {
	case PostgresDB:
		cfg.PostgresConfig = PostgresConfig{
			DB{
				DbHost,
				DbPort,
				DbUser,
				DbPassword,
				DbName,
				DbTimeout,
			},
			DbMaxconn,
		}
	case MongoDB:
		cfg.MongoConfig = MongoConfig{
			DB{
				DbHost,
				DbPort,
				DbUser,
				DbPassword,
				DbName,
				DbTimeout,
			},
			Collections{
				posCollection,
				empCollection,
			},
		}
	default:
		return Config{}, errors.New("incorrect database")
	}

	return cfg, nil
}
