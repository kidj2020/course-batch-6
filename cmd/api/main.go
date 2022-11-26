package main

import (
	"latihan-course-batch-6/cmd/internal/app/database"
	"latihan-course-batch-6/cmd/internal/app/exercise/handler"
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
	r.GET("/exercises/:id", middleware.WithAuth(), exerciseHandler.GetExerciseByID)
	r.GET("/exercises/:id/score", middleware.WithAuth(), exerciseHandler.GetScore)
	r.POST("/register", userHandler.Register)
	r.Run(":1234")
}
