package mongo

import (
	"context"
	"fmt"
	"github.com/dilyara4949/employees-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMongo(cfg config.DB) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s/?authSource=admin", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.Name), nil
}
