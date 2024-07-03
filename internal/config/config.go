package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	JWTTokenSecret string
	RestPort       string
	GrpcPort       string
	Address        string
	DatabaseType   string
	PostgresConfig
	MongoConfig
	RedisConfig
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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Timeout  time.Duration
	PoolSize int
	Database int
	Ttl      time.Duration
}

const (
	defaultRedisTimeout  = 10
	defaultRedisDB       = 0
	defaultRedisPoolSize = 10
	defaultRedisTtl      = 5
)

var (
	errMissingRestPort         = errors.New("REST_PORT is empty")
	errMissingAddress          = errors.New("ADDRESS is empty")
	errMissingJWTTokenSecret   = errors.New("JWT_TOKEN_SECRET is empty")
	errMissingGrpcPort         = errors.New("GRPC_PORT is empty")
	errMissingMongoHost        = errors.New("MONGO_HOST is empty")
	errMissingPostgresHost     = errors.New("POSTGRES_HOST is empty")
	errMissingMongoPort        = errors.New("MONGO_PORT is empty")
	errMissingPostgresPort     = errors.New("POSTGRES_PORT is empty")
	errMissingMongoName        = errors.New("Mongo_USER is empty")
	errMissingPostgresName     = errors.New("POSTGRES_USER is empty")
	errMissingMongoUser        = errors.New("MONGO_USER is empty")
	errMissingPostgresUser     = errors.New("POSTGRES_USER is empty")
	errMissingMongoPassword    = errors.New("MONGO_PASSWORD is empty")
	errMissingPostgresPassword = errors.New("POSTGRES_PASSWORD is empty")
	errMissingMongoTimeout     = errors.New("MONGO_TIMEOUT is empty")
	errMissingPostgresTimeout  = errors.New("POSTGRES_TIMEOUT is empty")
	errMissingPostgresMaxConn  = errors.New("POSTGRES_MAX_CONNECTIONS is empty")
	errPostgresMaxConnType     = errors.New("POSTGRES_MAX_CONNECTIONS must be an integer")
	errMongoTimeoutType        = errors.New("MONGO_TIMEOUT must be an integer")
	errPostgresTimeoutType     = errors.New("POSTGRES_TIMEOUT must be an integer")
	errMissingRedisHost        = errors.New("REDIS_HOST is empty")
	errMissingRedisPort        = errors.New("REDIS_PORT is empty")
	errMissingRedisPass        = errors.New("REDIS_PASSWORD is empty")
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
	mongoHostEnv        = "MONGO_HOST"
	mongoPortEnv        = "MONGO_PORT"
	mongoUserEnv        = "MONGO_USER"
	mongoPasswordEnv    = "MONGO_PASSWORD"
	mongoNameEnv        = "MONGO_NAME"
	mongoTimeoutEnv     = "MONGO_TIMEOUT"
	postgresHostEnv     = "POSTGRES_HOST"
	postgresPortEnv     = "POSTGRES_PORT"
	postgresUserEnv     = "POSTGRES_USER"
	postgresPasswordEnv = "POSTGRES_PASSWORD"
	postgresNameEnv     = "POSTGRES_NAME"
	postgresTimeoutEnv  = "POSTGRES_TIMEOUT"
	postgresMaxConnEnv  = "POSTGRES_MAX_CONNECTIONS"
	posCollectionEnv    = "POSITIONS_COLLECTION"
	empCollectionEnv    = "EMPLOYEES_COLLECTION"
	redisHostEnv        = "REDIS_HOST"
	redisPortEnv        = "REDIS_PORT"
	redisPasswordEnv    = "REDIS_PASSWORD"
	redisTimeoutEnv     = "REDIS_TIMEOUT"
	redisTtlEnv         = "REDIS_TTL"
	redisDbEnv          = "REDIS_DATABASE"
	redisPoolSizeEnv    = "REDIS_POOL_SIZE"
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

	mongoHost := os.Getenv(mongoHostEnv)
	if mongoHost == "" {
		errs = append(errs, errMissingMongoHost)
	}

	postgresHost := os.Getenv(postgresHostEnv)
	if postgresHost == "" {
		errs = append(errs, errMissingPostgresHost)
	}

	mongoPort := os.Getenv(mongoPortEnv)
	if mongoPort == "" {
		errs = append(errs, errMissingMongoPort)
	}

	mongoUser := os.Getenv(mongoUserEnv)
	if mongoUser == "" {
		errs = append(errs, errMissingMongoUser)
	}

	mongoPassword := os.Getenv(mongoPasswordEnv)
	if mongoPassword == "" {
		errs = append(errs, errMissingMongoPassword)
	}

	mongoName := os.Getenv(mongoNameEnv)
	if mongoName == "" {
		errs = append(errs, errMissingMongoName)
	}

	mongoTimeoutStr := os.Getenv(mongoTimeoutEnv)
	if mongoTimeoutStr == "" {
		errs = append(errs, errMissingMongoTimeout)
	}

	mongoTimeout, err := strconv.Atoi(mongoTimeoutStr)
	if err != nil {
		errs = append(errs, errMongoTimeoutType)
	}

	postgresPort := os.Getenv(postgresPortEnv)
	if postgresPort == "" {
		errs = append(errs, errMissingPostgresPort)
	}

	postgresUser := os.Getenv(postgresUserEnv)
	if postgresUser == "" {
		errs = append(errs, errMissingPostgresUser)
	}

	postgresPassword := os.Getenv(postgresPasswordEnv)
	if postgresPassword == "" {
		errs = append(errs, errMissingPostgresPassword)
	}

	postgresName := os.Getenv(postgresNameEnv)
	if postgresName == "" {
		errs = append(errs, errMissingPostgresName)
	}

	postgresTimeoutStr := os.Getenv(postgresTimeoutEnv)
	if postgresTimeoutStr == "" {
		errs = append(errs, errMissingPostgresTimeout)
	}

	postgresTimeout, err := strconv.Atoi(postgresTimeoutStr)
	if err != nil {
		errs = append(errs, errPostgresTimeoutType)
	}

	postgresMaxConnStr := os.Getenv(postgresMaxConnEnv)
	if postgresMaxConnStr == "" {
		errs = append(errs, errMissingPostgresMaxConn)
	}

	postgresMaxconn, err := strconv.Atoi(postgresMaxConnStr)
	if err != nil {
		errs = append(errs, errPostgresMaxConnType)
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

	redisHost := os.Getenv(redisHostEnv)
	if redisHost == "" {
		errs = append(errs, errMissingRedisHost)
	}

	redisPort := os.Getenv(redisPortEnv)
	if redisPort == "" {
		errs = append(errs, errMissingRedisPort)
	}

	redisPass := os.Getenv(redisPasswordEnv)
	if redisPass == "" {
		errs = append(errs, errMissingRedisPass)
	}

	redisTimeout, err := strconv.Atoi(os.Getenv(redisTimeoutEnv))
	if err != nil {
		redisTimeout = defaultRedisTimeout
	}

	redisTtl, err := strconv.Atoi(os.Getenv(redisTtlEnv))
	if err != nil {
		redisTtl = defaultRedisTtl
	}

	redisDB, err := strconv.Atoi(os.Getenv(redisDbEnv))
	if err != nil {
		redisDB = defaultRedisDB
	}

	redisPoolSize, err := strconv.Atoi(os.Getenv(redisPoolSizeEnv))
	if err != nil {
		redisPoolSize = defaultRedisPoolSize
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
		RedisConfig: RedisConfig{
			Host:     redisHost,
			Port:     redisPort,
			Password: redisPass,
			Database: redisDB,
			PoolSize: redisPoolSize,
			Timeout:  time.Duration(redisTimeout) * time.Second,
			Ttl:      time.Duration(redisTtl) * time.Hour,
		},
		PostgresConfig: PostgresConfig{
			DB{
				postgresHost,
				postgresPort,
				postgresUser,
				postgresPassword,
				postgresName,
				postgresTimeout,
			},
			postgresMaxconn,
		},
		MongoConfig: MongoConfig{
			DB{
				mongoHost,
				mongoPort,
				mongoUser,
				mongoPassword,
				mongoName,
				mongoTimeout,
			},
			Collections{
				posCollection,
				empCollection,
			},
		},
	}

	return cfg, nil
}
