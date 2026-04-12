package initializers

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     RedisURL,
	Password: RedisPass,
	DB:       RedisDBNum,
})

func EnsureRedisConnected() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status := Rdb.Ping(ctx)
	err := status.Err()

	if err != nil {
		return err
	}

	return nil
}
