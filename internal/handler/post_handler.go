package handler

import "github.com/gin-gonic/gin"

// PostHandler はPostControllerのインターフェース
type PostHandler interface {
	Create(ctx *gin.Context)
	GetAll(c *gin.Context)
}
