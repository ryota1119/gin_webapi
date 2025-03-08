package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota1119/gin_webapi/internal/domain/usecase"
	"net/http"
)

type AuthMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
}

type AuthMiddlewareImpl struct {
	authUsecase usecase.AuthUsecase
}

var _ AuthMiddleware = (*AuthMiddlewareImpl)(nil)

func NewAuthMiddleware(authUsecase usecase.AuthUsecase) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		authUsecase: authUsecase,
	}
}

// AuthMiddleware は AccessToken を検証するミドルウェア
func (a *AuthMiddlewareImpl) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := a.authUsecase.ValidateUser(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
