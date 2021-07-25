package model

import "github.com/gomodule/redigo/redis"

var RedisPool redis.Pool

// 创建全局redis 连接池 句柄

func InitRedis() {
	// 创建函数, 初始化Redis连接池
	RedisPool = redis.Pool{
		MaxIdle:         20,     // 最大空闲数 == 初始化连接数
		MaxActive:       50,     // 最大存活数 > MaxIdle
		MaxConnLifetime: 60 * 5, // 最大生命周期。
		IdleTimeout:     60,     // 空闲超时时间。
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "9.135.222.178:6380",
				redis.DialPassword("ZMxpN*3726Zw"))
		},
	}
}
