package routes

import (
	"goLandCRUD/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
	g.GET("/questions/user/:userId", func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid userId"})
			return
		}
		var questions []models.Question
		questions, err = models.GetUserQuestionsList(userId)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve questions"})
			return
		}
		c.JSON(200, gin.H{
			"questions": questions,
		})
	})
}
