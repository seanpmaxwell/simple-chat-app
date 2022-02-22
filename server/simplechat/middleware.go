package simplechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO It is not clear to me why you have two almost identical middlewares?

// APIMiddleware checks that the jwt-cookie is present.
func APIMiddleware(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(GetCookieName())
	if err != nil || jwtstr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Pase the string and get the claims
	data, err := ParseJwt(jwtstr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Set Session Data
	session := parseJwtData(data)
	c.Set(GetSessionDataKey(), session)
	// Return
	c.Next()
}

// AuthMiddleware checks that the jwt-cookie is present.
func AuthMiddleware(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(GetCookieName())
	if jwtstr == "" || err != nil {
		c.Next()
		return
	}
	// Parse the string and get the claims
	data, err := ParseJwt(jwtstr)
	if data == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// Set Session Data
	session := parseJwtData(data)
	c.Set(GetSessionDataKey(), session)
	// Return
	c.Next()
}

// parseJwtData converts JWT into Session object
func parseJwtData(data *map[string]interface{}) *Session {
	return &Session{
		ID:    uint((*data)["id"].(float64)),
		Email: (*data)["email"].(string),
		Name:  (*data)["name"].(string),
	}
}
