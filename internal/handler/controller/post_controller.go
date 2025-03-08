package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/gin_webapi/internal/domain"
	"github.com/ryota1119/gin_webapi/internal/domain/usecase"
	"github.com/ryota1119/gin_webapi/internal/handler"
)

// PostController の実装
type PostController struct {
	postUsecase usecase.PostUsecase
}

var _ handler.PostHandler = (*PostController)(nil)

// NewPostController はPostControllerの初期化を行う
func NewPostController(postUsecase usecase.PostUsecase) handler.PostHandler {
	return &PostController{postUsecase}
}

// Create は新しい記事を作成する
//
//	@Summary		記事の作成
//	@Description	新しい記事を作成する
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			post	body		domain.Post	true	"記事データ"
//	@Success		201		{object}	domain.Post
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/posts [post]
func (h *PostController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.postUsecase.Create(ctx, post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// GetAll は全ての記事を取得する
//
//	@Summary		記事一覧の取得
//	@Description	全ての記事を取得する
//	@Tags			posts
//	@Produce		json
//	@Success		200	{array}		domain.Post
//	@Failure		500	{object}	map[string]string
//	@Router			/posts [get]
func (h *PostController) GetAll(c *gin.Context) {
	// ctx := c.Request.Context()
	c.JSON(http.StatusOK, "")

	// posts, err := h.postUsecase.GetAll(ctx)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, posts)
}
