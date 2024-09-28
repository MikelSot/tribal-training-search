package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedis(ctx context.Context, client *redis.Client) Redis {
	return Redis{
		ctx:    ctx,
		client: client,
	}
}

func (r Redis) Set(key string, value interface{}, exp time.Duration) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error json.Marshal(): %w", err)
	}

	value = valueBytes
	if err := r.client.Set(r.ctx, key, value, exp).Err(); err != nil {
		return fmt.Errorf("error redis.client.Set(): %w", err)
	}

	return nil
}

func (r Redis) Get(key string) (string, error) {
	value, err := r.client.Get(r.ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		log.Print("error redis.client.Get(): key not found")

		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("error redis.client.Get(): %w", err)
	}

	return value, err
}
