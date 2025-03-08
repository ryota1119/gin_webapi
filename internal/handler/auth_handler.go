package handler

import "github.com/gin-gonic/gin"

// AuthHandler はAuthControllerのインターフェース
type AuthHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
}
