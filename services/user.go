package services

import (
	"goLandCRUD/logger"
	"goLandCRUD/models"
	"goLandCRUD/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error("Error Parsing the body", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := user.Save()
	if err != nil {
		logger.Error("Error saving the user", err)
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(201, gin.H{
		"message": "User registered successfully",
	})
}

func LoginUser(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		logger.Error("Error Parsing the body", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := credentials.ValidateCredentials()
	if err != nil {
		logger.Error("Authentication failed", err)
		c.JSON(401, gin.H{"error": "Authentication failed"})
		return
	}
	token, err := utils.GenerateToken(credentials.Email, credentials.ID)
	if err != nil {
		logger.Error("Error generating token", err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
