package di

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(ctx context.Context, cfg *Config) (*mongo.Database, error) {
	if cfg.DB == "" || cfg.DB_APP == "" {
		return nil, fmt.Errorf("invalid config: DB and DB_APP must not be empty")
	}

	opts := options.Client().
		ApplyURI(cfg.DB).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return client.Database(cfg.DB_APP), nil
}
