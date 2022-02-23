package main

import (
	"fmt"
	"simple-chat-app/server/src/models"
	"simple-chat-app/server/src/routers"
	"simple-chat-app/server/src/shared"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	serverStartMsg = "Gin server running on localhost"
	dnsStr         = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

type Server struct {
	EnvVars    *shared.EnvVars
	apiRouter  *routers.ApiRouter
	middleware *routers.Middlware
}

/**
New() and Dependency-Injection.
*/
func NewServer(
	envVars *shared.EnvVars,
	apiRouter *routers.ApiRouter,
	middleware *routers.Middlware,
) *Server {
	return &Server{envVars, apiRouter, middleware}
}

/**
https://github.com/go-gorm/postgres
*/
func getDbConn(envVars *shared.EnvVars) *gorm.DB {
	// Setup connection string
	dbParams := envVars.DbParams
	dsn := fmt.Sprintf(dnsStr, dbParams.Host, dbParams.User, dbParams.Pwd, dbParams.Name,
		dbParams.Port)
	// Open connection
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// Migrate GORM models
	conn.AutoMigrate(&models.User{}, &models.UserCreds{})
	// Init connection
	return conn
}

/**
Start the gin engine.
*/
func (s *Server) Start() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, serverStartMsg)
	})
	s.addRoutes(engine)
	engine.Use()
	engine.Run()
}

/**
Setup all routes
*/
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
