package simplechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// handleUserDeleteRequest handles API request to deletes a user
func handleUserDeleteRequest(c *gin.Context) {
	// Convert query string to uint
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Delete the user
	err = DeleteUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

// handleUserUpdateRequest handles API request to update user's name and email
func handleUserUpdateRequest(c *gin.Context) {
	// Extract user from json
	var req UpdateUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	err = UpdateUserByID(req.ID, req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

// handleUserAddRequest handles API request to add a new user.
func handleUserAddRequest(c *gin.Context) {
	var user *User
	// Extract user from json
	var req AddUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// Query db
	user, err = AddUser(req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"user":   user, // TODO This may need conversion
	})
}

// handleGetAllUsersRequest handles API request to fetch all users.
func handleGetAllUsersRequest(c *gin.Context) {
	users, _, err := FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
