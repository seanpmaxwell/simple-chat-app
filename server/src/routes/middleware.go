package routes

import (
	"net/http"
	"simple-chat-app/server/src/shared"
	"simple-chat-app/server/src/util"

	"github.com/gin-gonic/gin"
)

/**
Check the jwt-cookie is present.
*/
func apiMiddleware(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(shared.GetCookieParams().Name)
	if err != nil || jwtstr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Pase the string and get the claims
	data, err := util.ParseJwt(jwtstr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	// Set Session Data
	session := parseJwtData(data)
	c.Set(shared.SessionDataKey(), session)
	// Return
	c.Next()
}

/**
Check the jwt-cookie is present.
*/
func authMiddleware(c *gin.Context) {
	// Get the jwt string from the cookie
	jwtstr, err := c.Cookie(shared.GetCookieParams().Name)
	if jwtstr == "" || err != nil {
		c.Next()
		return
	}
	// Pase the string and get the claims
	data, err := util.ParseJwt(jwtstr)
	if data == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	// Set Session Data
	session := parseJwtData(data)
	c.Set(shared.SessionDataKey(), session)
	// Return
	c.Next()
}

/**
The API middleware needs this too
*/
func parseJwtData(data *map[string]interface{}) *SessionData {
	return &SessionData{
		ID:    uint((*data)["id"].(float64)),
		Email: (*data)["email"].(string),
		Name:  (*data)["name"].(string),
	}
}
