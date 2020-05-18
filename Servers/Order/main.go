package main

import (
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/consul"

	"micro.demo/Library/log"
	"micro.demo/Library/orm"
	"micro.demo/Proto/pb"
	"micro.demo/Servers/Order/Migrate"
	"micro.demo/Servers/User/Handler"
)

var (
	register registry.Registry
	service  micro.Service
)

func init() {
	// 注册consul
	register = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
		registry.Timeout(5*time.Second),
	)

	service = micro.NewService(
		micro.Name("micro.demo.wallet"),
		micro.Version("v1"),
		micro.Address(":803"),
		micro.Registry(register),
	)
}

func main() {
	service.Init(
		micro.Action(func(context *cli.Context) {
			/************************************/
			/********** 	 数据迁移      ********/
			/************************************/
			orm.AutoMigrate(Migrate.AutoMigrate())

			/************************************/
			/********** 	 RPC服务      ********/
			/************************************/
			_ = pb.RegisterUserServiceHandler(
				service.Server(),
				new(Handler.UserHandler),
				server.InternalHandler(true),
			)

			/************************************/
			/********** 	 消息队列      ********/
			/************************************/
			// queue.Init("192.168.1.228:4150", "192.168.1.228:4161", 1)
			// Subscriber.Register() // 注册消费者
		}),
	)

	// 启动service
	if err := service.Run(); err != nil {
		log.Error("wallet-srv服务启动失败 ...")
	}
}
