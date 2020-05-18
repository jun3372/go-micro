package Subscriber

import (
	"micro.demo/Library/log"
	"micro.demo/Library/queue"
)

// 注册消费者
func Register() {
	// 可监听多个主题
	err := queue.Subscribe("go.micro.srv.service", "sayHello", new(SayHello))
	if err != nil {
		log.Fatal(err)
		return
	}
}
