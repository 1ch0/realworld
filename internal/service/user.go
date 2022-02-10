package service

import (
	"context"
	v1 "realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (rep *v1.UserResponse, err error) {
	return &v1.UserResponse{
		User: &v1.UserResponse_User{
			Username: "test",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (rep *v1.UserResponse, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserResponse{
		User: &v1.UserResponse_User{
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (rep *v1.UserResponse, err error) {
	return &v1.UserResponse{
		User: &v1.UserResponse_User{
			Username: "boom",
		},
	}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (rep *v1.UserResponse, err error) {
	return &v1.UserResponse{
		User: &v1.UserResponse_User{
			Username: "boom",
		},
	}, nil
}
