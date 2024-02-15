package redis

import (
	"context"
	"fmt"
	"iTask/config"

	"github.com/go-redis/redis/v8"
)

var (
	defaultRedisMaxActive = 0 // unlimited max active connection
	defaultRedisMaxIdle   = 10
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password, // no set password
		DB:           cfg.Redis.DB,       // use default db
		PoolSize:     defaultRedisMaxActive,
		MinIdleConns: defaultRedisMaxIdle,
	})

	// ping to test connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
