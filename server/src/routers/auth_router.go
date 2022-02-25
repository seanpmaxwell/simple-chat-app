package routers

import (
	"net/http"
	"simple-chat-app/server/src/services"
	"simple-chat-app/server/src/shared"
	"simple-chat-app/server/src/shared/constants"
	"simple-chat-app/server/src/util"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRouter struct {
	EnvVars     *shared.EnvVars
	JwtUtil     *util.JwtUtil
	AuthService *services.AuthService
}

// New()
func NewAuthRouter(
	envVars *shared.EnvVars,
	jwtUtil *util.JwtUtil,
	authService *services.AuthService,
) *AuthRouter {
	return &AuthRouter{envVars, jwtUtil, authService}
}

// URL: "/api/auth/login"
func (a *AuthRouter) Login(c *gin.Context) {
	// Set req data
	var loginReq LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Verify and fetch the user
	user, err := a.AuthService.VerifyAndFetchUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Get a jwt string if the user passed authentication
	sessionData := SessionData{user.ID, user.Email, user.Name}
	jwtstr, err := a.JwtUtil.Sign(&sessionData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"login": err.Error()})
		return
	}
	// Set the cookie
	cookieParams := a.EnvVars.CookieParams
	c.SetCookie(cookieParams.Name, jwtstr, cookieParams.Exp, cookieParams.Path,
		cookieParams.Domain, cookieParams.Secure, true)
	// Return json
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// URL: "/api/auth/logout"
// Logout user by setting cookies maxAge = 0 and removing jwtstr
func (a *AuthRouter) Logout(c *gin.Context) {
	// Set the cookie
	cookieParams := a.EnvVars.CookieParams
	c.SetCookie(cookieParams.Name, "", 0, cookieParams.Path, cookieParams.Domain,
		cookieParams.Secure, true)
	// Return
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// URL: "/api/auth/session"
func (a *AuthRouter) SessionData(c *gin.Context) {
	// Check if the user is not logged in, if not that's okay,
	// there just won't be any session data
	session, exists := c.Get(constants.SessionDataKey())
	if !exists {
		c.JSON(http.StatusOK, gin.H{"logged-in": false})
		return
	}
	// Return the data if it's there
	c.JSON(http.StatusOK, gin.H{"data": session})
}
