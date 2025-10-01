package routes

import (
	"goLandCRUD/logger"
	"goLandCRUD/middlewares"
	"goLandCRUD/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
	g.POST("/register", services.RegisterUser)
	g.POST("/login", services.LoginUser)
	authenticated := g.Group("/").Use(middlewares.Authenticate)
	authenticated.GET("/questions/user/:userId", services.GetQuestionByUserId)
	authenticated.POST("/question", services.CreateQuestion)
	logger.Info("Routes created successfully")
}
