package main

import (
	"context"
	"fmt"
	"learningMicroService/proto/sum"

	"github.com/micro/go-micro/v2"
)

func main() {
	//创建新的服务
	service := micro.NewService(micro.Name("go.micro.learning.srv.sum"))

	//初始化微服务
	service.Init()

	//创建一个新的求和服务器
	sumService := sum.NewSumService("go.micro.learning.srv.sum", service.Client())

	rsp, err := sumService.GetSum(context.Background(), &sum.SumRequest{Input: 10})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Output)
}
