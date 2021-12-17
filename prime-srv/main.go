package main

import (
	"learningMicroService/prime-srv/handler"
	"learningMicroService/proto/prime"

	"github.com/micro/go-micro/v2"
)

func main() {
	//创建服务
	srv := micro.NewService(micro.Name("go.micro.learning.srv.prime"))
	//初始化服务
	srv.Init()
	//挂载接口
	_ = prime.RegisterPrimeHandler(srv.Server(), handler.GetHandler())
	//运行
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
