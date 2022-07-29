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
	Image string `json:"image"`
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

	c.JSON(http.StatusCreated, gin.H{"data": question})
}

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	initializers.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

type VoteQuestionInput struct {
	Type models.VoteQuestionType `json:"type" binding:"required"`
}


func VoteQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input VoteQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vote models.VoteQuestion
	vote = models.VoteQuestion{QuestionID: question.ID, Type: input.Type, UserID: c.MustGet("user").(models.User).ID}

	initializers.DB.Create(&vote)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully voted!"})
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
	user := c.MustGet("user").(models.User)
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	// check if user is creator of question
	if question.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the creator of this question!"})
		return
	}

	initializers.DB.Delete(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

type UpdateQuestionInput struct {
	Content string `json:"content"`
}

func UpdateQuestion(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// check if user is creator of question
	if question.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the creator of this question!"})
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
