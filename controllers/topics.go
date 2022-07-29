package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dblab/questlist/models"
	"dblab/questlist/initializers"
)

type CreateTopicInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func CreateTopic(c *gin.Context) {
	// Validate input
	var input CreateTopicInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create topic
	topic := models.Topic{Title: input.Title, Description: input.Description}
	initializers.DB.Create(&topic)

	c.JSON(http.StatusOK, gin.H{"data": topic})
}

type UpdateTopicInput struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

func UpdateTopic(c *gin.Context) {
	// Get model if exist
	var topic models.Topic

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&topic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Update topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&topic)

	c.JSON(http.StatusOK, gin.H{"data": topic})
}

func GetTopics(c *gin.Context) {
	var topics []models.Topic
	initializers.DB.Find(&topics)

	c.JSON(http.StatusOK, gin.H{"data": topics})
}

func GetTopic(c *gin.Context) {
	var topic models.Topic

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&topic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": topic})
}

func DeleteTopic(c *gin.Context) {
	var topic models.Topic

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&topic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&topic)

	c.JSON(http.StatusOK, gin.H{"data": topic})
}
