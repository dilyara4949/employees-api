package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	JWTTokenSecret string
	RestPort       string
	GrpcPort       string
	Address        string
	RedisConfig
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
	errMissingPort           = errors.New("PORT is empty")
	errMissingAddress        = errors.New("ADDRESS is empty")
	errMissingJWTTokenSecret = errors.New("JWT_TOKEN_SECRET is empty")
	errMissingRedisHost      = errors.New("REDIS_HOST is empty")
	errMissingRedisPort      = errors.New("REDIS_PORT is empty")
	errMissingRedisPass      = errors.New("REDIS_PASSWORD is empty")
)

func NewConfig() (Config, error) {
	errs := make([]error, 0)

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

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		errs = append(errs, errMissingRedisHost)
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		errs = append(errs, errMissingRedisPort)
	}

	redisPass := os.Getenv("REDIS_PASSWORD")
	if redisPass == "" {
		errs = append(errs, errMissingRedisPass)
	}

	redisTimeout, err := strconv.Atoi(os.Getenv("REDIS_TIMEOUT"))
	if err != nil {
		redisTimeout = defaultRedisTimeout
	}

	redisTtl, err := strconv.Atoi(os.Getenv("REDIS_TTL"))
	if err != nil {
		redisTtl = defaultRedisTtl
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		redisDB = defaultRedisDB
	}

	redisPoolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
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
		RedisConfig: RedisConfig{
			Host:     redisHost,
			Port:     redisPort,
			Password: redisPass,
			Database: redisDB,
			PoolSize: redisPoolSize,
			Timeout:  time.Duration(redisTimeout) * time.Second,
			Ttl:      time.Duration(redisTtl) * time.Hour,
		},
	}

	return cfg, nil
}
