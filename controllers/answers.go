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
}

func CreateAnswer(c *gin.Context) {
	var user models.User = c.MustGet("user").(models.User)
	// Validate input
	var input CreateAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create answer
	answer := models.Answer{Content: input.Content, QuestionID: input.QuestionId, UserID: user.ID}
	initializers.DB.Create(&answer)

	c.JSON(http.StatusCreated, gin.H{"data": answer})
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
	user := c.MustGet("user").(models.User)
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if user.ID != answer.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this answer!"})
		return
	}


	initializers.DB.Delete(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

type UpdateAnswerInput struct {
	Content string `json:"content" binding:"required"`
}

func UpdateAnswer(c *gin.Context) {
	var user models.User = c.MustGet("user").(models.User)
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if user.ID != answer.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this answer!"})
		return
	}

	// Validate input
	var input UpdateAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer.Content = input.Content
	answer.UserID = user.ID 

	initializers.DB.Save(&answer)

	c.JSON(http.StatusOK, gin.H{"data": answer})
}

type AnswerWithUser struct {
	models.Answer
	User NestedUserReturn `json:"user"`
}

func GetAnswerByQuestion(c *gin.Context) {
	var answers []models.Answer
	initializers.DB.Where("question_id = ?", c.Param("id")).Find(&answers)

	var answerWithUser []AnswerWithUser
	for _, answer := range answers {
		var user models.User
		initializers.DB.Where("id = ?", answer.UserID).First(&user)
		myAnswer := AnswerWithUser{answer, NestedUserReturn{
			ID: user.ID,
			Username: user.Username,
		}}
		answerWithUser = append(answerWithUser, myAnswer)
	}

	c.JSON(http.StatusOK, gin.H{"data": answerWithUser})
}

func GetAnswerByUser(c *gin.Context) {
	var answers []models.Answer
	initializers.DB.Where("user_id = ?", c.Param("id")).Find(&answers)

	c.JSON(http.StatusOK, gin.H{"data": answers})
}

type VoteAnswerInput struct {
	Type models.VoteAnswerType `json:"type" binding:"required"`
}

func VoteAnswer(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input VoteAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vote models.VoteAnswer
	vote = models.VoteAnswer{AnswerID: answer.ID, Type: input.Type, UserID: c.MustGet("user").(models.User).ID}

	// Check if vote exists and update it
	if err := initializers.DB.Where("answer_id = ? AND user_id = ?", answer.ID, c.MustGet("user").(models.User).ID).First(&vote).Error; err == nil {
		vote.Type = input.Type
		initializers.DB.Save(&vote)
	} else {
		initializers.DB.Create(&vote)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully voted!"})
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

func GetAnswerVotes(c *gin.Context) {
	var answer models.Answer

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&answer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var votes []models.VoteAnswer
	initializers.DB.Where("answer_id = ?", answer.ID).Find(&votes)

	c.JSON(http.StatusOK, gin.H{"data": votes})
}

func DeleteVoteAnswer(c *gin.Context) {
	var vote models.VoteAnswer
	var user models.User = c.MustGet("user").(models.User)

	if err := initializers.DB.Where("question_id = ? AND user_id = ?", c.Param("id"), user.ID).First(&vote).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// check if user is the owner of the vote
	if vote.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this vote!"})
		return
	}

	initializers.DB.Delete(&vote)

	c.JSON(http.StatusOK, gin.H{"data": vote})
}
