package service

import (
	"context"

	v1 "github.com/xiaoyan648/project-layout/api/helloworld/v1"
	"github.com/xiaoyan648/project-layout/internal/model"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *model.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *model.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &model.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
