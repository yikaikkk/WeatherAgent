package dbService

import (
	"WeatherAgent/src/global"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	ctx = context.Background()
)

type RedisClient struct {
	*redis.Client
}

var redisHelper *RedisClient

var redisOnce sync.Once

func GetRedisHelper() *RedisClient {
	return redisHelper
}

func InitRedisClient(config *global.Config) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         config.Redis.Host + ":" + config.Redis.Port,
		Password:     config.Redis.Password, // no password set
		DB:           config.Redis.DB,       // use default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     config.Redis.PoolSize,
		PoolTimeout:  30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisClient)
		rdh.Client = rdb
		redisHelper = rdh
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("failed to connect to Redis: " + err.Error())
	} else {
		println("Redis connected successfully")
	}

}
