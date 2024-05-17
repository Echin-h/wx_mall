package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"wxapp/config"
)

var Redis *redis.Client

func CacheLoading() {
	redisConf := config.GetConfig().Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Address,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect ping failed, err:", err)
		return
	}
	Redis = rdb
	fmt.Println("redis connect success...")
}
