package controllers

import (
	"github.com/gin-gonic/gin"
	"dblab/questlist/models"
	"dblab/questlist/initializers"
	"net/http"
	"strings"
)

type CreateQuestionInput struct {
	Title  string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Topics string `json:"topic" binding:"required"`
	Type models.QuestionType `json:"type"`
}

func CreateQuestion(c *gin.Context) {
	var user models.User
	user = c.MustGet("user").(models.User)
	// Validate input
	var input CreateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// split topics
	topics := strings.Split(input.Topics, ",")
	for i := 0; i < len(topics); i++ {
		topics[i] = strings.TrimSpace(topics[i])
	}

	// get topics from db
	var topicsDB []models.Topic
	initializers.DB.Where("name IN (?)", topics).Find(&topicsDB)


	// Create question
	question := models.Question{Content: input.Content, UserID: user.ID, Title: input.Title, Topics: topicsDB, }
	initializers.DB.Create(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	initializers.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func GetQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": question})
}

func DeleteQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

type UpdateQuestionInput struct {
	Content string `json:"content"`
}

func UpdateQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.Content = input.Content

	initializers.DB.Save(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

func UpvoteQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	question.Upvote++

	initializers.DB.Save(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

func DownvoteQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	question.Downvote++

	initializers.DB.Save(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

