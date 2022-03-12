package routers

/**** Types ****/

// ApiRouter layer
type ApiRouter struct {
	AuthRouter *AuthRouter
	UserRouter *UserRouter
}


/**** Functions ****/

// Wire()
func WireApiRouter(authRouter *AuthRouter, userRouter *UserRouter) *ApiRouter {
	return &ApiRouter{authRouter, userRouter}
}
