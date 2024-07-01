package redis

import (
	"context"
	"fmt"
	"github.com/dilyara4949/employees-api/internal/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg config.RedisConfig) (*redis.Client, error) {
	ctx := context.Background()
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    cfg.Password,
		DB:          cfg.Database,
		PoolSize:    cfg.PoolSize,
		PoolTimeout: cfg.Timeout,
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}
	return client, nil
}
