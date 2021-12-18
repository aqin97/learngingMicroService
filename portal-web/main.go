package main

import (
	"context"
	"learningMicroService/proto/sum"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/v2/web"
)

var (
	sumServiceClient sum.SumService
)

func main() {
	//新建一个web服务
	service := web.NewService(
		web.Name("sum"),
		web.Address(":8080"),
		web.StaticDir("html"),
	)

	//初始化
	service.Init()

	sumServiceClient = sum.NewSumService("go.micro.learning.srv.sum", service.Options().Service.Options().Client)

	service.Handle("/", http.FileServer(http.Dir("../html")))
	service.HandleFunc("/sum", Sum)

	//
	if err := service.Run(); err != nil {
		panic(err)
	}

}

func Sum(w http.ResponseWriter, r *http.Request) {
	inputString := r.URL.Query().Get("input")
	input, err := strconv.ParseInt(inputString, 10, 10)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	req := &sum.SumRequest{
		Input: input,
	}

	//客户端
	rsp, err := sumServiceClient.GetSum(context.Background(), req)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(strconv.Itoa(int(rsp.Output))))
}
