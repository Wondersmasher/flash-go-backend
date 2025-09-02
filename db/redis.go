package db

import (
	"context"
	"log"
	"time"

	"github.com/flash-backend/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	log.Println("Initializing Redis...")
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDRESS,
		Password: config.REDIS_PASSWORD,
		Username: config.REDIS_USERNAME,
		DB:       0,
		Protocol: 2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Set(ctx, "foo2", "bar", 0).Err()
	if err != nil {
		log.Fatalf("Redis connection error: %v", err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatalf("Redis GET error: %v", err)
	}
	log.Println("foo", val)

}
