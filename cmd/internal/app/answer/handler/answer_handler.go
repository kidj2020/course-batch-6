package handler

import (
	"latihan-course-batch-6/cmd/internal/app/exercise/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnswerHandler struct {
	db *gorm.DB
}

func NewAnswerHandler(db *gorm.DB) *AnswerHandler {
	return &AnswerHandler{db: db}
}

func (uh AnswerHandler) NewAnswer(c *gin.Context) {
	exerciseIdString := c.Param("exerciseId")
	questionIdString := c.Param("questionId")
	userId, ok := c.Request.Context().Value("user_id").(int)
	if !ok {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid userId",
		})
		return
	}

	exerciseId, err := strconv.Atoi(exerciseIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid exerciseId",
		})
		return
	}

	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid questionId",
		})
		return
	}

	var newAnswer domain.CreateAnswersRequest
	if err := c.ShouldBind(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}
	answer, err := domain.CreateNewAnswer(
		exerciseId,
		questionId,
		userId,
		newAnswer.Answer)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := uh.db.Create(answer).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"message": "sukses",
	})
}
