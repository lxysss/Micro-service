package main

import (
	"bj38web/service/user/handler"
	"bj38web/service/user/model"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"

	user "bj38web/service/user/proto/user"
)

func main() {
	couslReg := consul.NewRegistry()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Registry(couslReg),
		micro.Version("latest"),
		micro.Address("192.168.255.10:10001"),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))
	// 初始化Mysql
	model.InitDb()
	// 初始化Redis
	model.InitRedis()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
