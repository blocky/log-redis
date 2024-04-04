package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
)

func main() {
	rdb, err := NewRedisClient()
	if err != nil {
		log.Fatalf("Error creating Redis client: %v", err)
	}

	ctx := context.Background()
	// Configure Redis to send notifications for keyspace events.
	_, err = rdb.Do(
		ctx, "CONFIG", "SET", "notify-keyspace-events", "KEA",
	).Result()
	if err != nil {
		log.Fatalf("Error configuring Redis: %v", err)
	}
	// subscribe to the key event channel for "set" events.
	keyEvent := rdb.PSubscribe(ctx, "__keyevent@0__:set").Channel()
	// log the key and value for each set event.
	for {
		key := (<-keyEvent).Payload
		value := rdb.Get(ctx, key).Val()
		msg := fmt.Sprintf("Key '%s' has been set to value '%s'", key, value)
		slog.Info(msg)
	}
}
