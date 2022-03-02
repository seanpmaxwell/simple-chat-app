package routers

// Layer
type ApiRouter struct {
	AuthRouter *AuthRouter
	UserRouter *UserRouter
}

// Wire()
func WireApiRouter(authRouter *AuthRouter, userRouter *UserRouter) *ApiRouter {
	return &ApiRouter{authRouter, userRouter}
}
