package main

import (
	"bj38web/service/getCaptcha/handler"
	"bj38web/service/getCaptcha/model"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"

	getCaptcha "bj38web/service/getCaptcha/proto/getCaptcha"
)

func main() {
	couslReg := consul.NewRegistry()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.getCaptcha"),
		micro.Registry(couslReg),
		micro.Version("latest"),
		micro.Address("192.168.255.10:10000"),
	)
	// 初始化Redis链接池
	model.InitRedis()

	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
