package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "realworld/api/realworld/v1"
	"realworld/internal/biz"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.GreeterUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements realworld.GreeterServer
func (s *RealWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
