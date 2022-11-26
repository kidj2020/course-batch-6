package middleware

import (
	"context"
	"latihan-course-batch-6/cmd/internal/app/exercise/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized 1",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized 2",
			})
			c.Abort()
			return
		}

		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized 3",
			})
			c.Abort()
			return
		}
		var user domain.User
		data, err := user.DecryptJWT(auths[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		userID := int(data["user_id"].(float64))
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userID)
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
