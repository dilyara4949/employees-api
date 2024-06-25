package redis

import (
	"fmt"
	"github.com/dilyara4949/employees-api/internal/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    cfg.Password,
		DB:          cfg.Database,
		PoolSize:    cfg.PoolSize,
		PoolTimeout: cfg.Timeout,
	})

	return client
}
