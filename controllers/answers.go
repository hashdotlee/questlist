package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dblab/questlist/models"
	"dblab/questlist/initializers"
)

type CreateAnswerInput struct {
	Content string `json:"content" binding:"required"`
	QuestionId uint `json:"question_id" binding:"required"`
	UserID uint `json:"user_id"`
}

func CreateAnswer(c *gin.Context) {
	// Validate input
	var input CreateAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create answer
	answer := models.Answer{Content: input.Content, QuestionID: input.QuestionId, UserID: input.UserID}
	initializers.DB.Create(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func GetAnswers(c *gin.Context) {
	var answers []models.Answer
	initializers.DB.Find(&answers)

	c.JSON(http.StatusOK, gin.H{"data": answers})
}

func GetAnswer(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func DeleteAnswer(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

type UpdateAnswerInput struct {
	Content string `json:"content" binding:"required"`
	UserID uint `json:"user_id"`	
}

func UpdateAnswer(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.Content = input.Content
	answer.UserID = input.UserID

	initializers.DB.Save(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func GetAnswerByQuestion(c *gin.Context) {
	var answers []models.Answer
	initializers.DB.Where("question_id = ?", c.Param("id")).Find(&answers)

	c.JSON(http.StatusOK, gin.H{"data": answers})
}

func GetAnswerByUser(c *gin.Context) {
	var answers []models.Answer
	initializers.DB.Where("user_id = ?", c.Param("id")).Find(&answers)

	c.JSON(http.StatusOK, gin.H{"data": answers})
}

func UpvoteAnswer(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	answer.Upvote++

	initializers.DB.Save(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func DownvoteAnswer(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	answer.Downvote++

	initializers.DB.Save(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

func AcceptAnswer(c *gin.Context) {
	var user models.User = c.MustGet("user").(models.User)
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}	

	if user.ID == answer.UserID {
		answer.Verified = true
		initializers.DB.Save(&answer)
		c.JSON(http.StatusOK, gin.H{"data": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": false})
	}

}