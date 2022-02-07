package service

import (
	"context"
	v1 "realworld/api/realworld/v1"
)

func (r *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (rep *v1.UserResponse, err error) {
	return &v1.UserResponse{
		User: &v1.UserResponse_User{
			Username: "test",
		},
	}, nil
}
