package usecase

import (
	"context"

	"github.com/ryota1119/gin_webapi/internal/domain"
)

// PostUsecase PostUsecaseのインンターフェースを定義
type PostUsecase interface {
	// Create は投稿を作成する
	Create(ctx context.Context, post domain.Post) error
	// GetAll は投稿を取得する
	GetAll(ctx context.Context) ([]domain.Post, error)
}
