package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisClient{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (r *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	value, err := r.Client.Get(r.Ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}

func (r *RedisClient) SetJSON(key string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.Client.Set(r.Ctx, key, jsonData, expiration).Err()
}

func (r *RedisClient) GetJSON(key string, dest interface{}) error {
	value, err := r.Client.Get(r.Ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), dest)
}

func (r *RedisClient) Invalidate(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *RedisClient) InvalidPrefix(prefix string) error {
	var cursor uint64
	for {
		keys, nextCursor, err := r.Client.Scan(r.Ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			return fmt.Errorf("error scanning keys: %w", err)
		}

		if len(keys) > 0 {
			if err := r.Client.Del(r.Ctx, keys...).Err(); err != nil {
				return fmt.Errorf("error deleting keys: %w", err)
			}
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
	return nil
}
