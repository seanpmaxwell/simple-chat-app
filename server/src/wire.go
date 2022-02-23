//go:build wireinject
// +build wireinject

package main

import (
	"simple-chat-app/server/src/repos"
	"simple-chat-app/server/src/routers"
	"simple-chat-app/server/src/services"
	"simple-chat-app/server/src/shared"
	"simple-chat-app/server/src/util"

	"github.com/google/wire"
)

func InitializeServer() (*Server, error) {
	wire.Build(
		shared.NewEnvVars,
		getDbConn,
		util.NewJwtUtil,
		util.NewPwdUtil,
		repos.NewUserRepo,
		services.NewUserService,
		services.NewAuthService,
		routers.NewMiddleware,
		routers.NewUserRouter,
		routers.NewAuthRouter,
		routers.NewApiRouter,
		NewServer,
	)
	return &Server{}, nil
}
