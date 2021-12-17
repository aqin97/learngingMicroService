package main

import (
	"learningMicroService/proto/sum"

	"learningMicroService/sum-srv/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	//创建服务
	srv := micro.NewService(micro.Name("go.micro.learning.srv.sum"))

	//服务初始化
	srv.Init(micro.BeforeStart(func() error {
		log.Log("启动前的日志")
		return nil
	}), micro.AfterStart(func() error {
		log.Log("启动后的日志")
		return nil
	}))

	//挂载接口
	_ = sum.RegisterSumHandler(srv.Server(), handler.GetHandler())

	//运行
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
