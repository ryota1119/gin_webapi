package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

// NewRedis　はRedisの接続を初期化する
func NewRedis() error {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	redisClient = rdb
	log.Println("Redis connection established successfully.")
	return nil
}

// GetRedisClient はredis.Clientを返す
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		log.Fatal("Database is not initialized. Call NewDB() first.")
	}
	return redisClient
}
