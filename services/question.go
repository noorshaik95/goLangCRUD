package services

import (
	"goLandCRUD/logger"
	"goLandCRUD/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQuestionByUserId(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		logger.Error("Error Parsing the userId from Context", err)
		c.JSON(400, gin.H{"error": "Invalid userId"})
		return
	}
	var questions []models.Question
	questions, err = models.GetUserQuestionsList(userId)
	if err != nil {
		logger.Error("Error fetching the list of questions", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve questions"})
		return
	}
	c.JSON(200, gin.H{
		"questions": questions,
	})
}

func CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		logger.Error("Error Parsing the body", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	question.UserID = c.GetInt64("userId")
	if question.Title == "" || question.Body == "" || question.UserID == 0 {
		logger.Error("Missing fields", question)
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}
	err := question.CreateQuestion()
	if err != nil {
		logger.Error("Creation of question failed", err)
		c.JSON(500, gin.H{"error": "Failed to create question"})
		return
	}
	err = question.GetQuestionById()
	if err != nil {
		logger.Error("Question Created but failed to retrieve", err)
		c.JSON(500, gin.H{"error": "Question Created but failed to retrieve"})
		return
	}
	c.JSON(201, gin.H{
		"message":  "Question created successfully",
		"question": question,
	})
}
