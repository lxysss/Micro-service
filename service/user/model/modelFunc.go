package model

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func CheckSmsCode(phone, code string) error {
	// 校验短信验证码
	// 链接redis
	conn := RedisPool.Get()

	// 从 redis 中, 根据 key 获取 Value --- 短信验证码  码值
	smsCode, err := redis.String(conn.Do("get", phone+"_code"))
	if err != nil {
		fmt.Println("redis get phone_code err:", err)
		return err
	}
	// 验证码匹配  失败
	if smsCode != code {
		return errors.New("验证码匹配失败!")
	}
	// 匹配成功!
	return nil
}

func RegisterUser(mobile, pwd string) error {
	// 注册用户信息,写 MySQL 数据库.
	var user User
	user.Name = mobile // 默认使用手机号作为用户名

	// 使用 md5 对 pwd 加密
	m5 := md5.New()                             // 初始md5对象
	m5.Write([]byte(pwd))                       // 将 pwd 写入缓冲区
	pwd_hash := hex.EncodeToString(m5.Sum(nil)) // 不使用额外的秘钥

	user.Password_hash = pwd_hash
	user.Mobile = mobile

	// 插入数据到MySQL
	return GlobalConn.Create(&user).Error
}
