package routers

import (
	"net/http"
	"simple-chat-app/server/server/shared"
	"simple-chat-app/server/server/shared/constants"
	"simple-chat-app/server/server/util"
	"strings"

	"github.com/gin-gonic/gin"
)

/**** Types ****/

// Things stored in a login session
type SessionData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Middleware Layer
type Middlware struct {
	EnvVars *shared.EnvVars
	JwtUtil *util.JwtUtil
}


/**** Functions ****/

// Wire Middleware
func WireMiddleware(envVars *shared.EnvVars, jwtUtil *util.JwtUtil) *Middlware {
	return &Middlware{envVars, jwtUtil}
}

// Check the jwt-cookie is present.
func (m *Middlware) SessionMw(c *gin.Context) {
	// Get the jwt string from the cookie. For the session data route if the jwt is not present
	// just continue and assument the
	jwtstr, err := c.Cookie(m.EnvVars.CookieParams.Name)
	if err != nil || jwtstr == "" {
		if strings.HasSuffix(c.Request.URL.Path, "/api/auth/session-data") {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
		}
		return
	}
	// Pase the string and get the claims
	data, err := m.JwtUtil.Parse(jwtstr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Parse JWT data
	session := &SessionData{
		ID:    uint((*data)["id"].(float64)),
		Email: (*data)["email"].(string),
		Name:  (*data)["name"].(string),
	}
	// Set session Data
	c.Set(constants.SessionDataKey(), session)
	c.Next()
}
