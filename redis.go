package main

import (
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		log.Fatal("REDIS_ADDRESS environment variable must be set")
	}
	redisDatabaseStr := os.Getenv("REDIS_DATABASE")
	if redisDatabaseStr == "" {
		log.Fatal("REDIS_DATABASE environment variable must be set")
	}
	redisDatabase, err := strconv.Atoi(redisDatabaseStr)
	if err != nil {
		log.Fatalf("Error converting REDIS_DATABASE to int: %v", err)
	}
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     redisAddress,
			Password: redisPassword,
			DB:       redisDatabase,
		})
	return rdb, nil
}
