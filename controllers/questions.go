package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"dblab/questlist/models"
	"dblab/questlist/initializers"
	"net/http"
	"os"
	"time"
)

type CreateQuestionInput struct {
	Title  string `json:"title" binding:"required"`
	UserId string `json:"user_id"`
	Content string `json:"content" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

func CreateQuestion(c *gin.Context) {
	// Validate input
	var input CreateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create question
	question := models.Question{Content: input.Content, UserId: input.UserId, Title: input.Title, Topic: input.Topic}
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

func UpVoteQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	question.UpVotes++

	initializers.DB.Save(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

func DownVoteQuestion(c *gin.Context) {
	var question models.Question

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	question.DownVotes++

	initializers.DB.Save(&question)

	c.JSON(http.StatusOK, gin.H{"data": question})
}

