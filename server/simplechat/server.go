package simplechat

import (
	"github.com/gin-gonic/gin"
)

const serverStartMsg = "Gin server running on localhost"

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	s := &Server{
		Engine: gin.Default(),
	}
	// Add routers (Groups) to the gin-engine
	addRoutes(s.Engine)

	return s
}

// Start starts the Gin server.
func (s *Server) Start() error {

	// Start the server
	return s.Engine.Run()
}
