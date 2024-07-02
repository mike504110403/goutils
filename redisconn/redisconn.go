package redisconn

import (
	"github.com/gomodule/redigo/redis"
)

func New(Cfg RedisConfig) {
	pool := &redis.Pool{
		MaxIdle:     Cfg.MaxIdle,
		MaxActive:   Cfg.MaxActive,
		IdleTimeout: Cfg.IdleTimeout,
		Dial:        Cfg.Dial,
	}
	redisPool = pool
}

func Get() redis.Conn {
	return redisPool.Get()
}
