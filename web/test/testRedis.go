package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main(){
	// 连接数据库
	conn, err := redis.Dial("tcp", "9.135.222.178:6380",
		redis.DialPassword("ZMxpN*3726Zw"),
	)
	if err != nil {
		fmt.Println("redis.Dial err",err)
		return
	}
	defer conn.Close()

	// 操作数据库
	reply, err := conn.Do("set", "hello", "potterliu")
	if err != nil {
		fmt.Println("conn.Do err",err)
		return
	}

	// 回复助手函数
	r, e := redis.String(reply,err)
	fmt.Println(r,e)


}
