package service

import (
	v1 "realworld/api/realworld/v1"
	"realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.UserUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}
