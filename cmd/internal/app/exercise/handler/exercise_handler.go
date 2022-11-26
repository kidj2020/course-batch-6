package handler

import (
	"latihan-course-batch-6/cmd/internal/app/exercise/domain"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseHandler struct {
	db *gorm.DB
}

func NewExerciseHandler(db *gorm.DB) *ExerciseHandler {
	return &ExerciseHandler{db: db}
}

func (eh ExerciseHandler) GetExerciseByID(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
		return
	}
	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, exercise)
}

func (eh ExerciseHandler) GetScore(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
		return
	}
	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}

	userID := ctx.Request.Context().Value("user_id").(int)

	var answers []domain.Answer
	err = eh.db.Where("exercise_id = ? AND user_id = ?", id, userID).Find(&answers).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "not answered yet",
		})
		return
	}

	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score Score
	wg := new(sync.WaitGroup)
	for _, question := range exercise.Questions {
		wg.Add(1)
		go func(question domain.Questions) {
			defer wg.Done()
			if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
				score.Inc(question.Score)
			}
		}(question)
	}
	wg.Wait()
	ctx.JSON(http.StatusOK, map[string]int{
		"score": score.total,
	})
}

type Score struct {
	total int
	mu    sync.Mutex
}

func (s *Score) Inc(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total += value
}
