package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func SaveImageCode(code string, uuid string) error {
	// 存储图片到redis
	// 连接数据库
	conn := RedisPool.Get()
	defer conn.Close()
	// 操作数据库 插入Redis设置有效时间
	_, err := conn.Do("setex", uuid, 60*5, code)
	if err != nil {
		fmt.Println("conn.Do err", err)
		return err
	}
	return err
}

func SaveSmsCode(phone string, code string) error {
	// 存储短信验证码 到redis
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", phone+"_code", 60*3, code)
	if err != nil {
		return err
	}
	return nil
}

func CheckImgCode(uuid, imgCode string) bool {
	// 链接 redis
	conn := RedisPool.Get()
	defer conn.Close()

	// 查询 redis 数据
	code, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		fmt.Println("查询错误 err:", err)
		return false
	}

	// 返回校验结果
	return code == imgCode
}