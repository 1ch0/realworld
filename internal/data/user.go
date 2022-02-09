package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, g *biz.User) error {
	//r.data.db.Create()
	return nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
