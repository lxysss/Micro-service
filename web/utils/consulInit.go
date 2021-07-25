package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

var GloabMicro micro.Service

func InitMicro() {
	// 初始化客户端
	consulReg := consul.NewRegistry()
	GloabMicro = micro.NewService(
		micro.Registry(consulReg),
	)
}
