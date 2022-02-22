package simplechat

import "github.com/gin-gonic/gin"

// addRoutes sets up the API URL routes
// NOTE I put this here to be part of the reusable server, but
// it could easily be moved to app/main.go and only be part of
// the app.
func addRoutes(engine *gin.Engine) {

	engine.GET("/", func(c *gin.Context) {
		c.String(200, serverStartMsg)
	})

	api := engine.Group("/api")

	// et up the routes for users
	users := api.Group("/users")
	users.GET("/", handleGetAllUsersRequest)
	users.POST("/", handleUserAddRequest)
	users.PUT("/", handleUserUpdateRequest)
	users.DELETE("/:id", handleUserDeleteRequest)
	users.Use(APIMiddleware)

	// Add the auth-router (Group) to the gin-engine.
	auth := api.Group("/auth")
	auth.PUT("/handleLoginRequest", handleLoginRequest)
	auth.GET("/handleLogoutRequest", handleLogoutRequest)
	auth.GET("/session-data", handleGetSessionRequest)
	auth.Use(AuthMiddleware)

}
