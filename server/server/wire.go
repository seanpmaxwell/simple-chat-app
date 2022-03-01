//go:build wireinject
// +build wireinject

package server

import (
	"simple-chat-app/server/server/repos"
	"simple-chat-app/server/server/routers"
	"simple-chat-app/server/server/services"
	"simple-chat-app/server/server/shared"
	"simple-chat-app/server/server/util"

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
