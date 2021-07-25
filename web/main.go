package main

import (
	"bj38web/web/controller"
	"bj38web/web/model"
	"bj38web/web/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化 Session 对象
		s := sessions.Default(ctx)
		userName := s.Get("userName")
		if userName == nil {
			ctx.Abort() // 从这里返回, 不必继续执行了
		} else {
			ctx.Next() // 继续向下
		}
	}
}

func main() {
	// 初始化路由
	router := gin.Default()
	// 初始化容器
	store, _ := redis.NewStore(10, "tcp", "9.135.222.178:6380", "ZMxpN*3726Zw", []byte("bj38"))
	router.Use(sessions.Sessions("mysession", store))
	// 路由匹配
	router.Static("/home", "view")
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSeesion)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		r1.GET("/areas", controller.GetArea)
		r1.POST("/sessions", controller.PostLogin)

		//r1.Use(LoginFilter()) //以后的路由,都不需要再校验 Session 了. 直接获取数据即可!
		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)
		r1.POST("/user/avatar", controller.PostAvatar)
	}
	utils.InitMicro()
	model.InitRedis()
	model.InitDb()
	// 启动运行
	err := router.Run(":8080")	if err != nil {
		fmt.Println("router.Run", err)
		return
	}
}
