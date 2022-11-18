package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
)

func NewClient(addr, username, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()); err != nil {
		return nil, err.Err()
	}

	return rdb, nil
}
