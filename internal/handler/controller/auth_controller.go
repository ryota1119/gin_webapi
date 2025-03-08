package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/gin_webapi/internal/domain/usecase"
	"github.com/ryota1119/gin_webapi/internal/handler"
)

// AuthController の実装
type AuthController struct {
	authUsecase usecase.AuthUsecase
}

var _ handler.AuthHandler = (*AuthController)(nil)

// NewAuthController はAuthControllerの初期化を行う
func NewAuthController(authUsecase usecase.AuthUsecase) handler.AuthHandler {
	return &AuthController{authUsecase}
}

// Register はユーザーを新規作成する
//
//	@Summary		新規ユーザーを作成
//	@Description	新規ユーザーを作成
//	@Tags			users
//	@Success		200
//	@Router			/auth/register [post]
func (h *AuthController) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.authUsecase.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

// Login はログインを実行する
//
//	@Summary		ログイン
//	@Description	ログイン
//	@Tags			users
//	@Success		200
//	@Router			/auth/login [post]
func (h *AuthController) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUsecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *AuthController) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUsecase.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *AuthController) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	if err := h.authUsecase.Logout(ctx, req.RefreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
