package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "9.135.222.178:6380", "ZMxpN*3726Zw", []byte("bj38"))
	route.Use(sessions.Sessions("mysession", store))
	route.GET("/test", func(context *gin.Context) {
		// 设置cookie
		//context.SetCookie("mytest","djshfufhvosof",3600,"","",true,true)
		//context.Writer.WriteString("测试 Cookie ...")

		// 使用Session
		session := sessions.Default(context)
		// 设置session
		session.Set("itcast","sakdfokfoas")
		session.Save()

		context.Writer.WriteString("测试 Session")
		v := session.Get("itcast")
		fmt.Println("获取 Session:", v.(string))

	})
	route.Run(":9999")
}
