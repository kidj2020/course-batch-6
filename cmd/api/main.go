package main

import (
	answerHandler "latihan-course-batch-6/cmd/internal/app/answer/handler"
	"latihan-course-batch-6/cmd/internal/app/database"
	"latihan-course-batch-6/cmd/internal/app/exercise/handler"
	questionHandler "latihan-course-batch-6/cmd/internal/app/question/handler"
	userHandler "latihan-course-batch-6/cmd/internal/app/user/handler"
	"latihan-course-batch-6/cmd/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "hello world",
		})
	})

	db := database.NewConnDatabase()
	exerciseHandler := handler.NewExerciseHandler(db)
	userHandler := userHandler.NewUserHandler(db)
	questionHandler1 := questionHandler.NewQuestionHandler(db)
	answerHandler := answerHandler.NewAnswerHandler(db)

	r.GET("/exercises/:id", middleware.WithAuth(), exerciseHandler.GetExerciseByID)
	r.GET("/exercises/:id/score", middleware.WithAuth(), exerciseHandler.GetScore)

	r.POST("/exercises/:id/questions", middleware.WithAuth(), questionHandler1.CreateQuestion)
	r.POST("/exercises/:exerciseId/questions/:questionId/answer", answerHandler.NewAnswer)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	r.Run(":1234")
}
