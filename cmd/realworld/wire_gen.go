// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/biz"
	"realworld/internal/conf"
	"realworld/internal/data"
	"realworld/internal/server"
	"realworld/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	profileRepo := data.NewProfileRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, profileRepo, logger)
	realWorldService := service.NewRealWorldService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, realWorldService, logger)
	grpcServer := server.NewGRPCServer(confServer, realWorldService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
