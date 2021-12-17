package handler

import (
	"context"
	"learningMicroService/prime-srv/service"
	"learningMicroService/proto/prime"
)

type handler struct{}

func (h handler) GetPrime(ctx context.Context, req *prime.PrimeRequest, rsp *prime.PrimeResponse) error {
	input := req.Input
	rsp.Output = service.GetPrime(input)

	return nil
}

func GetHandler() prime.PrimeHandler {
	return handler{}
}
