package server

import (
	"simple-chat-app/server/server/routers"
	"simple-chat-app/server/server/shared"

	"github.com/gin-gonic/gin"
)

// **** Vals **** //

const (
	serverStartMsg = "Gin server running on localhost"
)


/**** Types ****/

// Server
type Server struct {
	EnvVars    *shared.EnvVars
	apiRouter  *routers.ApiRouter
	middleware *routers.Middlware
}


/**** Functions ****/

// Wire Server
func WireServer(
	envVars *shared.EnvVars,
	apiRouter *routers.ApiRouter,
	middleware *routers.Middlware,
) *Server {
	return &Server{envVars, apiRouter, middleware}
}

// Start the gin engine.
func (s *Server) Start() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, serverStartMsg)
	})
	s.addRoutes(engine)
	engine.Use()
	engine.Run()
}

// Setup all routes
func (s *Server) addRoutes(engine *gin.Engine) {
	//// Setup API routes
	apiGroup := engine.Group("/api")
	// Setup auth routes
	authGroup := apiGroup.Group("/auth")
	ar := s.apiRouter.AuthRouter
	authGroup.PUT("/login", ar.Login)
	authGroup.GET("/logout", ar.Logout)
	authGroup.Use(s.middleware.SessionMw)
	authGroup.GET("/session-data", ar.SessionData)
	// Setup user routes
	apiGroup.Use(s.middleware.SessionMw)
	userGroup := apiGroup.Group("/users")
	ur := s.apiRouter.UserRouter
	userGroup.GET("/", ur.FetchAll)
	userGroup.POST("/", ur.Add)
	userGroup.PUT("/", ur.Update)
	userGroup.DELETE("/:id", ur.Delete)
	//// Setup Static routes
	// TODO
}
