package handler

import (
	"context"
	"learningMicroService/proto/sum"
	"learningMicroService/sum-srv/service"
)

type handler struct{}

func (h handler) GetSum(ctx context.Context, req *sum.SumRequest, rsp *sum.SumResponse) error {
	input := req.Input
	rsp.Output = service.GetSum(input)
	return nil
}

func GetHandler() sum.SumHandler {
	return handler{}
}
