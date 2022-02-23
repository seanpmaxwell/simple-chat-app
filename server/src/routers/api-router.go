package routers

type ApiRouter struct {
	AuthRouter *AuthRouter
	UserRouter *UserRouter
}

/**
New()
*/
func NewApiRouter(authRouter *AuthRouter, userRouter *UserRouter) *ApiRouter {
	return &ApiRouter{authRouter, userRouter}
}
