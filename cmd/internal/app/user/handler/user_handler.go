package handler

import (
	"latihan-course-batch-6/cmd/internal/app/exercise/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (uh UserHandler) Register(c *gin.Context) {
	var userRegister domain.UserRegister
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	user, err := domain.NewUser(userRegister.Name, userRegister.Email, userRegister.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := uh.db.Create(user).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (uh UserHandler) Login(c *gin.Context) {
	var userLogin domain.UserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	var user domain.User
	err := uh.db.Where("email = ?", userLogin.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid username",
		})
		return
	}

	// err := user.Login(userLogin.Email, userLogin.Password).Take(&user)
	if err := user.Login(userLogin.Email, userLogin.Password); !err {
		// if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"message": user.Email + " invalid in uh",
		})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error generating token",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
