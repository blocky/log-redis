package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(
	address string,
	password string,
	database int,
) *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
		})
	return rdb
}

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	rdb := NewRedisClient(config.RedisAddress, config.RedisPassword, config.RedisDatabase)

	ctx := context.Background()
	// Configure Redis to send notifications for key events.
	_, err = rdb.Do(
		ctx, "CONFIG", "SET", "notify-keyspace-events", "KEA",
	).Result()
	if err != nil {
		log.Fatalf("Error configuring Redis: %v", err)
	}
	// subscribe to the key event channel for set events.
	keyEvent := rdb.PSubscribe(ctx, "__keyevent@0__:set").Channel()
	// log the key and value for each set event.
	for {
		key := (<-keyEvent).Payload
		value := rdb.Get(ctx, key).Val()
		msg := fmt.Sprintf("Key '%s' has been set to value '%s'", key, value)
		slog.Info(msg)
	}
}
