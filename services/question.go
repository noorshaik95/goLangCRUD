package services

import (
	"goLandCRUD/logger"
	"goLandCRUD/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQuestionByUserId(c *gin.Context) {
	userId := c.GetInt64("userId")
	var questions []models.QuestionWithUser
	questions, err := models.GetUserQuestionsList(userId)
	if err != nil {
		logger.Error("Error fetching the list of questions", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve questions"})
		return
	}
	c.JSON(200, gin.H{
		"questions": questions,
	})
}

func GetQuestionById(c *gin.Context) {
	var question models.QuestionWithUser
	questionId := c.Param("questionId")
	var err error
	question.Question.ID, err = strconv.ParseInt(questionId, 10, 64)
	if err != nil {
		logger.Error("Error Parsing the questionId from Context", err)
		c.JSON(400, gin.H{"error": "Invalid questionId"})
		return
	}
	err = question.GetQuestionDetails()
	if err != nil {
		logger.Error("Error fetching the question details", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve question details"})
		return
	}
	c.JSON(200, gin.H{
		"question": question,
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

func DeleteQuestion(c *gin.Context) {
	var question models.Question
	questionId := c.Param("questionId")
	var err error
	question.ID, err = strconv.ParseInt(questionId, 10, 64)
	if err != nil {
		logger.Error("Error Parsing the questionId from Context", err)
		c.JSON(400, gin.H{"error": "Invalid questionId"})
		return
	}
	err = question.GetQuestionById()
	if err != nil {
		logger.Error("Error fetching the question details", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve question details"})
		return
	}
	err = question.DeleteQuestion()
	if err != nil {
		logger.Error("Error deleting the question", err)
		c.JSON(500, gin.H{"error": "Failed to delete question"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Question deleted successfully",
	})
}

func UpdateQuestion(c *gin.Context) {
	var question models.Question
	questionId := c.Param("questionId")
	var err error
	question.ID, err = strconv.ParseInt(questionId, 10, 64)
	if err != nil {
		logger.Error("Error Parsing the questionId from Context", err)
		c.JSON(400, gin.H{"error": "Invalid questionId"})
		return
	}
	err = question.GetQuestionById()
	if err != nil {
		logger.Error("Error fetching the question details", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve question details"})
		return
	}
	if err := c.ShouldBindJSON(&question); err != nil {
		logger.Error("Error Parsing the body", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if question.Title == "" || question.Body == "" {
		logger.Error("Missing fields", question)
		c.JSON(400, gin.H{"error": "Missing required fields"})
		return
	}
	err = question.UpdateQuestion()
	if err != nil {
		logger.Error("Error updating the question", err)
		c.JSON(500, gin.H{"error": "Failed to update question"})
		return
	}
	err = question.GetQuestionById()
	if err != nil {
		logger.Error("Question Updated but failed to retrieve", err)
		c.JSON(500, gin.H{"error": "Question Updated but failed to retrieve"})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Question updated successfully",
		"question": question,
	})
}
