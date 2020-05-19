package main

import (
	"fmt"
	"os"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/consul"

	"micro.demo/Library/config"
	"micro.demo/Library/log"
	"micro.demo/Library/orm"
	"micro.demo/Proto/pb"
	"micro.demo/Servers/User/Handler"
	"micro.demo/Servers/User/Migrate"
)

var (
	register registry.Registry
	service  micro.Service
	cfg      *config.Config
)

// 注册配置
func InitConfig() {
	cfg = config.Cfg()
}

// 注册consul
func InitConsul() {
	register = consul.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%v",
			cfg.GetString("consul.host", os.Getenv("CONSUL_HOST")),
			cfg.GetString("consul.port", os.Getenv("CONSUL_PORT")),
		)),
		registry.Timeout(5*time.Second),
	)
}

// 注册服务
func InitServer() {
	name := cfg.GetString("server.name", os.Getenv("SERVER_NAME"))                        // 服务名称
	version := cfg.GetString("server.version", os.Getenv("SERVER_VERSION"))               // 服务版本号
	address := fmt.Sprintf(":%s", cfg.GetString("server.port", os.Getenv("SERVER_PORT"))) // 服务地址

	// 注册服务
	service = micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Address(address),
		micro.Registry(register),
	)
}

// 数据迁移
func InitMigrate() {
	/************************************/
	/********** 	 数据迁移      ********/
	/************************************/
	orm.AutoMigrate(Migrate.AutoMigrate())
}

// 注册RPC服务
func InitService() {
	/************************************/
	/********** 	 RPC服务      ********/
	/************************************/
	_ = pb.RegisterUserServiceHandler(
		service.Server(),
		new(Handler.User),
		server.InternalHandler(true),
	)
}

// 事件订阅
func InitSubscriber() {
	/************************************/
	/********** 	 消息队列      ********/
	/************************************/
	// queue.Init("192.168.1.228:4150", "192.168.1.228:4161", 1)
	// Subscriber.Register() // 注册消费者
}

func init() {
	// 注册配置文件
	InitConfig()

	// 注册consul
	InitConsul()

	// 注册服务
	InitServer()
}

func main() {
	// 注册其他
	service.Init(
		micro.Action(
			func(context *cli.Context) {
				InitMigrate()    // 数据库迁移
				InitService()    // 注册rpc服务
				InitSubscriber() // 注册事件订阅
			},
		),
	)

	// 启动service
	if err := service.Run(); err != nil {
		log.Error("user-srv服务启动失败 ...")
	}
}
