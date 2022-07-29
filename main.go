package main

import (
	_"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"dblab/questlist/controllers"
	"dblab/questlist/initializers"
	"dblab/questlist/middlewares"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDb()
}

func main() {
    r := gin.Default()
	
	// Initialize the database.

	  r.GET("/answers", controllers.GetAnswers)
	  r.POST("/answers", middlewares.RequireAuth, controllers.CreateAnswer)
	  r.GET("/answers/:id", controllers.GetAnswer)
	  r.DELETE("/answers/:id", middlewares.RequireAuth, controllers.DeleteAnswer)
	  r.PATCH("/answers/:id/upvote", middlewares.RequireAuth, controllers.UpvoteAnswer)
	  r.PATCH("/answers/:id/downvote", middlewares.RequireAuth, controllers.DownvoteAnswer)
	  r.PATCH("/answers/:id/accept", middlewares.RequireAuth, controllers.AcceptAnswer)
	  r.PUT("/answers/:id/update", middlewares.RequireAuth,controllers.UpdateAnswer)

	  r.GET("/questions", controllers.GetQuestions)
	  r.POST("/questions", middlewares.RequireAuth, controllers.CreateQuestion)
	  r.GET("/questions/:id", controllers.GetQuestion)
	  r.DELETE("/questions/:id", middlewares.RequireAuth, controllers.DeleteQuestion)
	  r.PATCH("/questions/:id/upvote", middlewares.RequireAuth, controllers.UpvoteQuestion)
	  r.PATCH("/questions/:id/downvote", middlewares.RequireAuth, controllers.DownvoteQuestion)
	  r.PUT("/questions/:id/update", middlewares.RequireAuth, controllers.UpdateQuestion)

	  r.POST("/signup", controllers.Signup)
	  r.POST("/login", controllers.Login)

	  r.PUT("/users/update", middlewares.RequireAuth, controllers.UpdateUser)
	  r.GET("/users/info", middlewares.RequireAuth, controllers.GetUser)
	  r.GET("/users", middlewares.RequireAdmin, controllers.GetUsers)
	  r.PATCH("/user/verify", middlewares.RequireAdmin, controllers.VerifyUser)

	  r.GET("/topics", controllers.GetTopics)
	  r.POST("/topics", middlewares.RequireAdmin, controllers.CreateTopic)
	  r.GET("/topics/:id", controllers.GetTopic)
	  r.DELETE("/topics/:id", middlewares.RequireAdmin, controllers.DeleteTopic)
	  r.PUT("/topics/:id/update", middlewares.RequireAdmin, controllers.UpdateTopic)
	  
    r.Run()
}

