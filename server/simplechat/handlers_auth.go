package simplechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// handleLoginRequest handles API requests to log in
// Log user in by setting JWT in cookie
func handleLoginRequest(c *gin.Context) {
	// Set req data
	var loginReq LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"handleLoginRequest": err.Error()})
		return
	}
	// Verify and fetch the user
	user, err := VerifyAndFetchUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"handleLoginRequest": err.Error()})
		return
	}
	// Get a jwt string if the user passed authentication
	sessionData := Session{user.ID, user.Email, user.Name}
	jwtstr, err := SignJwt(&sessionData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"handleLoginRequest": err.Error()})
		return
	}
	// Set the cookie
	setCookie(c, jwtstr)
	// Return json
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// handleLoginRequest handles API requests to log out
// Logout user by setting cookies maxAge = 0 and removing jwtstr
func handleLogoutRequest(c *gin.Context) {
	setCookie(c, "")
	// Return
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// handleGetSessionRequest returns current session data URL=/api/auth/session
func handleGetSessionRequest(c *gin.Context) {
	// Check if the user is not logged in, if not that's okay,
	// there just won't be any session data
	session, exists := c.Get(GetSessionDataKey())
	if !exists {
		c.JSON(http.StatusOK, gin.H{"logged-in": false})
		return
	}
	// Return the data if it's there
	c.JSON(http.StatusOK, gin.H{"data": session})
}

func setCookie(c *gin.Context, value string) {
	// Set the cookie
	cp := GetCookieParams()
	c.SetCookie(cp.Name, value, 0, cp.Path, cp.Domain, cp.Secure, true)
}
