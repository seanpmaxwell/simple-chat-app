package routers

import (
	"net/http"
	"simple-chat-app/server/server/services"

	"strconv"

	"github.com/gin-gonic/gin"
)

/**** Types ****/

// UserRouter layer
type UserRouter struct {
	UserService *services.UserService
}

// Add user request
type AddUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Update user request
type UpdateUserReq struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}


/**** Functions ****/

// Wire UserRouter
func WireUserRouter(userService *services.UserService) *UserRouter {
	return &UserRouter{userService}
}

// Fetch all users.
func (u *UserRouter) FetchAll(c *gin.Context) {
	users, err := u.UserService.FetchAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Add a new user.
func (u *UserRouter) Add(c *gin.Context) {
	// Extract user from json
	var req AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	err = u.UserService.Add(req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

// Update user's email and name.
func (u *UserRouter) Update(c *gin.Context) {
	// Extract user from json
	var req UpdateUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	err = u.UserService.Update(req.ID, req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

// Delete one user.
func (u *UserRouter) Delete(c *gin.Context) {
	// Convert query string to unint
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Delete the user
	err = u.UserService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}
