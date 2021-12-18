package main

import (
	"context"
	"learningMicroService/proto/sum"
	"log"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/v2/web"
)

var (
	srvClient sum.SumService
)

func main() {
	//创建web服务
	service := web.NewService(
		web.Name("go.micro.learning.web.portal"),
		web.Address(":8080"),
		web.StaticDir("/html"),
	)

	service.Init()

	srvClient = sum.NewSumService("go.micro.learning.srv.sum", service.Options().Service.Client())
	service.Handle("/", http.FileServer(http.Dir("html")))
	service.HandleFunc("/sum", Sum)

	if err := service.Run(); err != nil {
		//no err
		panic(err)
	}
}

func Sum(w http.ResponseWriter, r *http.Request) {
	inputString := r.URL.Query().Get("input")
	input, err := strconv.ParseInt(inputString, 10, 10)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	req := &sum.SumRequest{
		Input: input,
	}
	//客户端
	rsp, err := srvClient.GetSum(context.Background(), req)
	if err != nil {
		log.Printf("srvClient err")
	}
	w.Write([]byte(strconv.Itoa(int(rsp.Output))))
}
