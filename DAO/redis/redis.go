package redis

import (
	"Gin_Scaffold/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedis(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return err
}

func CloseRedis() {
	_ = rdb.Close()
}
