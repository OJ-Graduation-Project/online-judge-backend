package redis_pool

import (
	"fmt"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"
	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	redis_uri := fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port)
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 120000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_uri)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
