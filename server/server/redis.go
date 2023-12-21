package server

import (
	"fmt"
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/go-redis/redis"
	"github.com/gookit/config/v2"
	"time"
)

// Redis redis主连接池
var Redis struct {
	Main *redis.Client
}

// InitRedisClient 初始化redis客户端
func initRedisClient(redisConf *conf.RedisConfigStruct, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port), // Redis地址
		Password:    redisConf.Password,                                   // Redis密码
		DB:          db,                                                   // Redis库
		PoolSize:    16,                                                   // Redis连接池大小
		MaxRetries:  3,                                                    // 最大重试次数
		IdleTimeout: 10 * time.Second,                                     // 空闲链接超时时间
	})
	_, err := client.Ping().Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("[Redis] connection failed")
	} else if err != nil {
		return nil, fmt.Errorf("[Redis] connection failed: %w", err)
	}
	fmt.Printf("[Redis] Redis connected. (redis://***@%s:%d/%d)\n", redisConf.Host, redisConf.Port, db)
	return client, nil
}

// InitRedisDB 初始化Redis客户端
func InitRedisDB() error {
	redisConf := &conf.RedisConfigStruct{}
	var err error
	err = config.BindStruct("redis", redisConf)
	if err != nil {
		return err
	}
	// 连接0号库
	Redis.Main, err = initRedisClient(redisConf, redisConf.MainDB)
	if err != nil {
		return err
	}
	return nil
}
