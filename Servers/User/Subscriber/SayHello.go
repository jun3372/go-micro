package Subscriber

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"

	"micro.demo/Library/log"
	"micro.demo/Proto/pb"
)

type SayHello struct {
}

func (s SayHello) HandleMessage(m *nsq.Message) error {
	log.Info("读取队列")
	// return errors.New("错误")
	msg := pb.Message{}
	err := json.Unmarshal(m.Body, &msg)
	if err != nil {
		log.Warn(err)
		return err
	}

	log.Infof("header: %v   body: %v ", msg.Header, string(msg.Body))
	return nil
}
