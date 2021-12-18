package main

import (
	"context"
	"encoding/json"
	"learningMicroService/proto/sum"

	proto "learningMicroService/proto/log"
	"learningMicroService/sum-srv/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	//创建服务
	srv := micro.NewService(micro.Name("go.micro.learning.srv.sum"))

	//结合Wrapper和Broker
	//服务初始化
	srv.Init(
		//micro.WrapHandler(reqLogger(srv.Client())),
		micro.BeforeStart(func() error {
			log.Log("启动前的日志")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Log("启动后的日志")
			return nil
		}),
	)

	//挂载接口
	_ = sum.RegisterSumHandler(srv.Server(), handler.GetHandler())

	//运行
	if err := srv.Run(); err != nil {
		panic(err)
	}
}

//日志Wrapper， 通过Broker异步消息将日志推送到log-srv
func reqLogger(cli client.Client) server.HandlerWrapper {
	//初始化动作
	pub := micro.NewPublisher("go.micro.learning.topic.log", cli)

	return func(hf server.HandlerFunc) server.HandlerFunc {
		//中间动作
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Info("请求:准备发送日志")
			evt := proto.LogEvent{
				Msg: "hello",
			}
			body, _ := json.Marshal(evt.Msg)
			err := pub.Publish(ctx, &broker.Message{
				Header: map[string]string{
					"serveName": "sum",
				},
				Body: body,
			})
			if err != nil {
				return err
			}

			return hf(ctx, req, rsp)
		}
	}
}
