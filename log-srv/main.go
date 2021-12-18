//一个日志服务器，监听Broker的异步消息
package main

import (
	"context"

	proto "learningMicroService/proto/log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

type Sub struct{}

func (s Sub) Process(ctx context.Context, evt *proto.LogEvent) error {
	//业务逻辑，这个服务是监听其他服务推送过来的日志
	log.Logf("[sub] 收到日志: %s", evt.Msg)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.srv.log"),
	)

	service.Init()

	//监听的实现
	micro.RegisterSubscriber("go.micro.learning.topic.log", service.Server(), &Sub{})

	service.Run()
}
