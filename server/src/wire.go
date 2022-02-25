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

// Setup dependency injection
func InitializeServer() (*Server, error) {
	wire.Build(
		shared.WireEnvVars,
		WireDbConn,
		util.WireJwtUtil,
		util.WirePwdUtil,
		repos.WireUserRepo,
		services.WireUserService,
		services.WireAuthService,
		routers.WireMiddleware,
		routers.WireUserRouter,
		routers.WireAuthRouter,
		routers.WireApiRouter,
		WireServer,
	)
	return &Server{}, nil
}
