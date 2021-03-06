// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"simple-chat-app/server/server/repos"
	"simple-chat-app/server/server/routers"
	"simple-chat-app/server/server/services"
	"simple-chat-app/server/server/shared"
	"simple-chat-app/server/server/util"
)

// Injectors from wire.go:

// Setup dependency injection
func InitializeServer() (*Server, error) {
	envVars := shared.WireEnvVars()
	jwtUtil := util.WireJwtUtil(envVars)
	db := WireDbConn(envVars)
	userRepo := repos.WireUserRepo(db)
	pwdUtil := util.WirePwdUtil()
	authService := services.WireAuthService(userRepo, pwdUtil)
	authRouter := routers.WireAuthRouter(envVars, jwtUtil, authService)
	userService := services.WireUserService(userRepo, pwdUtil)
	userRouter := routers.WireUserRouter(userService)
	apiRouter := routers.WireApiRouter(authRouter, userRouter)
	middlware := routers.WireMiddleware(envVars, jwtUtil)
	server := WireServer(envVars, apiRouter, middlware)
	return server, nil
}
