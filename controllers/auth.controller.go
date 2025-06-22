package controllers

import (
	"fmt"
	"go-jwt-auth/models"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)

var users []models.User

func Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debug: Print current users
	fmt.Println("Current users before registration:", users)

	// Check if username already exists
	for _, user := range users {
		if user.Username == input.Username {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
	}

	user := models.User{
		ID:       uint(len(users) + 1),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	users = append(users, user)

	// Debug: Print users after registration
	fmt.Println("Users after registration:", users)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by usernamee
	var foundUser models.User
	for _, user := range users {
		if user.Username == input.Username {
			foundUser = user
			break
		}
	}

	// Check if user exists and password matchess
	if foundUser.ID == 0 || foundUser.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// TODO: Generate JWT token here
	// For now, just return success
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       foundUser.ID,
			"username": foundUser.Username,
			"email":    foundUser.Email,
		},
	})
}

func Profile(c *gin.Context) {
	// Get the username from the JWT claims (set in middleware)
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Find user by username
	var foundUser models.User
	for _, user := range users {
		if user.Username == username {
			foundUser = user
			break
		}
	}

	if foundUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       foundUser.ID,
			"username": foundUser.Username,
			"email":    foundUser.Email,
		},
	})
}

