package main

import (
	"log"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"

	"micro.demo/Api/routers"
)

func main() {
	/************************************/
	/********** 服务发现  cousul   ********/
	/************************************/
	reg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	// create new api service
	service := web.NewService(
		web.Name("micro.demo.api"),
		web.Registry(reg),
		web.RegisterTTL(time.Second*15),      // 重新注册时间
		web.RegisterInterval(time.Second*10), // 注册过期时间
		web.Address(":8082"),
	)

	service.Init(func(o *web.Options) {
		// 注册路由
		service.Handle("/", routers.Init())
	})

	// run service
	if err := service.Run(); err != nil {
		log.Panic(err.Error())
	}
}
