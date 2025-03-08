package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/ryota1119/gin_webapi/internal/handler"
	"github.com/ryota1119/gin_webapi/internal/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler はHandlerを実装
type Handler struct {
	authMiddleware middleware.AuthMiddleware
	authHandler    handler.AuthHandler
	postHandler    handler.PostHandler
}

// NewHandler はHandlerを初期化
func NewHandler(
	authMiddleware middleware.AuthMiddleware,
	authHandler handler.AuthHandler,
	postHandler handler.PostHandler,
) *Handler {
	return &Handler{
		authMiddleware: authMiddleware,
		authHandler:    authHandler,
		postHandler:    postHandler,
	}
}

// SetupRouter はRouterをセットアップする
func (h *Handler) SetupRouter(r *gin.Engine) {
	// **Swagger**
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")

	// ** 認証 API **
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/register", h.authHandler.Register)
		authGroup.POST("/login", h.authHandler.Login)
		authGroup.POST("/refresh", h.authHandler.RefreshToken)
		authGroup.POST("/logout", h.authHandler.Logout)
	}

	// ** 投稿 API **
	postGroup := v1.Group("/posts")
	postGroup.GET("", h.postHandler.GetAll)

	// 認証が必要なルートをリスト化
	protectedRoutes := []struct {
		method  string
		path    string
		handler gin.HandlerFunc
	}{
		{"POST", "/posts", h.postHandler.Create},
		// 他にも認証が必要なエンドポイントを追加可能
	}

	// 認証が必要なルートにミドルウェアを適用
	for _, route := range protectedRoutes {
		v1.Handle(route.method, route.path, h.authMiddleware.AuthMiddleware(), route.handler)
	}
}
