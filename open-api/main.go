package main

import (
	"context"
	"encoding/json"
	"fmt"
	"learningMicroService/proto/prime"
	"learningMicroService/proto/sum"
	"strconv"

	"github.com/micro/go-micro/v2"
	api "github.com/micro/go-micro/v2/api/proto"
)

var (
	sumServiceClient   sum.SumService
	primeServiceClient prime.PrimeService
)

type OpenAPI struct{}

func (o OpenAPI) Fetch(ctx context.Context, request *api.Request, response *api.Response) error {
	//我们目前有两个服务，求和和求素数
	sumInputStr := request.Get["sum"].Values[0]
	primeInputStr := request.Get["prime"].Values[0]

	sumInput, err := strconv.ParseInt(sumInputStr, 10, 10)
	if err != nil {
		fmt.Println(err)
	}
	primeInput, err := strconv.ParseInt(primeInputStr, 10, 10)
	if err != nil {
		fmt.Println(err)
	}

	sumReq := &sum.SumRequest{
		Input: sumInput,
	}
	primeReq := &prime.PrimeRequest{
		Input: primeInput,
	}

	//调用客户端
	sumRsp, err := sumServiceClient.GetSum(ctx, sumReq)
	if err != nil {
		fmt.Println(err)
	}
	primeRsp, err := primeServiceClient.GetPrime(ctx, primeReq)
	if err != nil {
		fmt.Println(err)
	}

	ret, err := json.Marshal(map[string]interface{}{
		"sum":   sumRsp,
		"prime": primeRsp,
	})
	if err != nil {
		fmt.Println(err)
	}

	response.Body = string(ret)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.api.openapi"),
	)

	service.Init()

	sumServiceClient = sum.NewSumService("sum", service.Client())
	primeServiceClient = prime.NewPrimeService("prime", service.Client())

	//暴露接口

	service.Run()
}
