package redisconn

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

type RedisConfig struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Dial        func() (redis.Conn, error)
}
