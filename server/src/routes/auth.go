package routes

import (
	"net/http"
	"simple-chat-app/server/src/services"
	"simple-chat-app/server/src/shared"
	"simple-chat-app/server/src/util"

	"github.com/gin-gonic/gin"
)

type SessionData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/**
Add the auth-router (Group) to the gin-engine.
*/
func initAuthRoutes(router *gin.RouterGroup) {
	group := router.Group("/auth")
	group.PUT("/login", login)
	group.GET("/logout", logout)
	group.Use(authMiddleware)
	group.GET("/session-data", getSessionData)
}

/**
URL: "/api/auth/login"
*/
func login(c *gin.Context) {
	// Set req data
	var loginReq LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Verify and fetch the user
	user, err := services.VerifyAndFetchUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Get a jwt string if the user passed authentication
	sessionData := SessionData{user.ID, user.Email, user.Name}
	jwtstr, err := util.SignJwt(&sessionData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Set the cookie
	cookieParams := shared.GetCookieParams()
	c.SetCookie(cookieParams.Name, jwtstr, cookieParams.Exp, cookieParams.Path,
		cookieParams.Domain, cookieParams.Secure, true)
	// Return json
	c.JSON(http.StatusOK, gin.H{"success": true})
}

/**
- URL: "/api/auth/logout"
- Logout user by setting cookies maxAge = 0 and removing jwtstr
*/
func logout(c *gin.Context) {
	// Set the cookie
	cookieParams := shared.GetCookieParams()
	c.SetCookie(cookieParams.Name, "", 0, cookieParams.Path, cookieParams.Domain,
		cookieParams.Secure, true)
	// Return
	c.JSON(http.StatusOK, gin.H{"success": true})
}

/**
URL: "/api/auth/session"
*/
func getSessionData(c *gin.Context) {
	// Check if the user is not logged in, if not that's okay,
	// there just won't be any session data
	session, exists := c.Get(shared.SessionDataKey())
	if !exists {
		c.JSON(http.StatusOK, gin.H{"logged-in": false})
		return
	}
	// Return the data if it's there
	c.JSON(http.StatusOK, gin.H{"data": session})
}
