package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"
	"log"
)

var Ctx = context.Background()

func ConnectRedis(config *config.Config) *redis.Client {
	log.Println("connecting to redis...")
	redisHost := config.RedisHost
	redisPort := config.RedisPort
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})
	log.Println("Connected to Redis")
	return rdb
}
