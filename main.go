package main

import (
	_"net/http"
	"github.com/gin-gonic/gin"
	"dblab/questlist/controllers"
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://hashdotlee.cyou", "http://localhost:3000" , "https://quest.hashdotlee.cyou"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))
	// Initialize the database.

	r.GET("/answers", controllers.GetAnswers)
	r.POST("/answers", middlewares.RequireAuth, controllers.CreateAnswer)
	r.GET("/answers/:id", controllers.GetAnswer)
	r.DELETE("/answers/:id", middlewares.RequireAuth, controllers.DeleteAnswer)
	r.POST("/answers/:id/vote", middlewares.RequireAuth, controllers.VoteAnswer)
	r.PATCH("/answers/:id/accept", middlewares.RequireAuth, controllers.AcceptAnswer)
	r.PUT("/answers/:id", middlewares.RequireAuth,controllers.UpdateAnswer)

	r.GET("/questions", controllers.GetQuestions)
	r.POST("/questions", middlewares.RequireAuth, controllers.CreateQuestion)
	r.GET("/questions/:id", controllers.GetQuestion)
	r.GET("/questions/:id/answers", controllers.GetAnswerByQuestion)
	r.DELETE("/questions/:id", middlewares.RequireAuth, controllers.DeleteQuestion)
	r.POST("/questions/:id/vote", middlewares.RequireAuth, controllers.VoteQuestion)
	r.PUT("/questions/:id", middlewares.RequireAuth, controllers.UpdateQuestion)

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

