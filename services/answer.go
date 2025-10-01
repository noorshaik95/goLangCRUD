package services

import (
	"goLandCRUD/logger"
	"goLandCRUD/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAnswer(c *gin.Context) {
	var answer models.Answer
	err := c.ShouldBindJSON(&answer)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	answer.UserID = c.GetInt64("userId")
	logger.Info("Create answer with userId: ", answer)
	if answer.Body == "" || answer.QuestionID == 0 || answer.UserID == 0 {
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}
	err = answer.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create answer"})
		return
	}
	c.JSON(201, gin.H{
		"message": "Answer created successfully",
		"answer":  answer,
	})
}

func GetAnswersByQuestionId(c *gin.Context) {
	var answers []models.AnswerWithUser
	questionId, err := strconv.ParseInt(c.Param("questionId"), 10, 64)
	if err != nil {
		logger.Error("Error Parsing the questionId from Context", err)
		c.JSON(400, gin.H{"error": "Invalid questionId"})
		return
	}
	answers, err = models.GetAllAnswersByQuestionId(questionId)
	if err != nil {
		logger.Error("Error fetching the list of answers", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve answers"})
		return
	}
	c.JSON(200, gin.H{
		"answers": answers,
	})
}
