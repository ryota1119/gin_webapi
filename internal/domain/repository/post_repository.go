package repository

import (
	"context"

	"github.com/ryota1119/gin_webapi/internal/domain"
)

// PostRepository PostRepositoryのインターフェースを定義
type PostRepository interface {
	// Create は投稿を作成する
	Create(ctx context.Context, post domain.Post) error
	// GetAll は投稿を全て取得する
	GetAll(ctx context.Context) ([]domain.Post, error)
}
