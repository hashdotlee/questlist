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
	Topics string `json:"topic"`
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
	question := models.Question{Content: input.Content, Image: input.Image, UserID: user.ID, Title: input.Title, Topics: topicsDB, Type: input.Type }
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

	// check if vote exists and update it
	if err := initializers.DB.Where("question_id = ? AND user_id = ?", question.ID, c.MustGet("user").(models.User).ID).First(&vote).Error; err == nil {
		vote.Type = input.Type
		initializers.DB.Save(&vote)
	} else {
		initializers.DB.Create(&vote)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully voted!"})
}

// create struct to return user 
type NestedUserReturn struct {
	ID uint `json:"id"`
	Username string `json:"username"`
}

func GetQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).Preload("Topics").First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// create struct for return data
	type Data struct {
		Question models.Question `json:"question"`
		User NestedUserReturn `json:"user"`
	}

	if(question.Type == models.QuestionTypePrivate) {
		c.JSON(http.StatusOK, gin.H{"data": Data{Question: question}})
		return
	}

	// Find user who created question
	var user models.User
	initializers.DB.Where("id = ?", question.UserID).First(&user)

	var userReturn NestedUserReturn
	userReturn.ID = user.ID
	userReturn.Username = user.Username

	// Set user to question
	var data Data
	data.Question = question
	data.User = userReturn 

	c.JSON(http.StatusOK, gin.H{"data": data})
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

func GetQuestionVotes(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var votes []models.VoteQuestion
	initializers.DB.Where("question_id = ?", question.ID).Find(&votes)

	c.JSON(http.StatusOK, gin.H{"data": votes})
}

func DeleteVoteQuestion(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	var vote models.VoteQuestion

	if err := initializers.DB.Where("question_id = ? AND user_id = ?", c.Param("id"), user.ID).First(&vote).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// check if user is creator of vote question
	if vote.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the creator of this vote question!"})
		return
	}

	initializers.DB.Delete(&vote)

	c.JSON(http.StatusOK, gin.H{"data": vote})
}
