package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MockUser is a hardcoded mock user for demonstration purposes.
var MockUser = gin.H{
	"username": "user1",
	"password": "password123", // In a real-world scenario, this should be securely hashed.
}

// Login function to authenticate user and generate JWT
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Parse JSON request body
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username and password match the mock user
	if loginData.Username != MockUser["username"] || loginData.Password != MockUser["password"] {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ProtectedRoute is an example of a route that requires JWT authentication
func ProtectedRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route!"})
}
