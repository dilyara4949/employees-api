package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dilyara4949/employees-api/internal/domain"
	"github.com/redis/go-redis/v9"
)

type employeeCache struct {
	client *redis.Client
	ttl    time.Duration
}

type EmployeeCache interface {
	Set(ctx context.Context, key string, value *domain.Employee) error
	Get(ctx context.Context, key string) (*domain.Employee, error)
	Delete(ctx context.Context, key string) error
}

func NewEmployeeCache(client *redis.Client, ttl time.Duration) EmployeeCache {
	return &employeeCache{client: client, ttl: ttl}
}

func (c *employeeCache) Set(ctx context.Context, key string, value *domain.Employee) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, c.ttl).Err()
}

func (c *employeeCache) Get(ctx context.Context, key string) (*domain.Employee, error) {
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value domain.Employee

	err = json.Unmarshal([]byte(data), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}
func (c *employeeCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
