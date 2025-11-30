package client

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

func (rc *RedisClient) Connect() {
	rc.Client = redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: "password",
		DB:       0,
	})
}

var redisClient *RedisClient

func GetRedisClient() *RedisClient {
	if redisClient == nil {
		redisClient = &RedisClient{}
	}
	return redisClient
}
