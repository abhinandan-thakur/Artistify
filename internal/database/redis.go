package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var Ctx = context.Background()

func ConnectRedis() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	return rdb
}
