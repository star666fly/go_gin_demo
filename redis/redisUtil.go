package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go_demo/config"
	"time"
)

var MyRedis *redis.Client

func Setup() {
	var c config.Conf
	conf := c.GetConf()
	MyRedis = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Password, // no password set
		DB:       conf.Redis.DB,       // use default DB
	})

	_, err := MyRedis.Ping().Result()
	if err != nil {
		fmt.Println("redis failed err:{}", err)
	} else {
		fmt.Println("redis connect success")
	}
}

func SetRedis(key string, value string, t int) bool {
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Set(key, value, expire).Err(); err != nil {
		return false
	}
	return true
}

func GetRedis(key string) string {
	result, err := MyRedis.Get(key).Result()
	if err != nil {
		return ""
	}
	return result
}

func DelRedis(key string) bool {
	_, err := MyRedis.Del(key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ExpireRedis(key string, t int) bool {
	// 延长过期时间
	expire := time.Duration(t) * time.Second
	if err := MyRedis.Expire(key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
