package utils

import (
	"context"
	"rantr/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	url := config.GetEnv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)

}

func SetRedisValue(rdb *redis.Client, key string, value string) error {
	return rdb.Set(ctx, key, value, 10*time.Hour).Err()
}

func GetRedisValue(rdb *redis.Client, key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func DeleteRedisValue(rdb *redis.Client, key string) error {
	return rdb.Del(ctx, key).Err()
}
