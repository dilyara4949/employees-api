package redis

import (
	"context"
	"encoding/json"
	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type positionCache struct {
	client *redis.Client
	ttl    time.Duration
}

type PositionCache interface {
	Set(ctx context.Context, key string, value *domain.Position) error
	Get(ctx context.Context, key string) (*domain.Position, error)
	Delete(ctx context.Context, key string) error
}

func NewPositionCache(client *redis.Client, ttl time.Duration) PositionCache {
	return &positionCache{client: client, ttl: ttl}
}

func (c *positionCache) Set(ctx context.Context, key string, value *domain.Position) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, c.ttl).Err()
}

func (c *positionCache) Get(ctx context.Context, key string) (*domain.Position, error) {
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value domain.Position
	err = json.Unmarshal([]byte(data), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (c *positionCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
