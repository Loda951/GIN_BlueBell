package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func InitRedis(config *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	_, err = client.Ping().Result()
	return err
}

func CloseRedis() {
	_ = client.Close()
}
