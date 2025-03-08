package usecase

import (
	"context"

	"github.com/ryota1119/gin_webapi/internal/domain"
	"github.com/ryota1119/gin_webapi/internal/domain/repository"
	"github.com/ryota1119/gin_webapi/internal/domain/usecase"
)

type PostUsecaseImpl struct {
	postRepo repository.PostRepository
}

var _ usecase.PostUsecase = (*PostUsecaseImpl)(nil)

func NewPostUsecase(postRepo repository.PostRepository) usecase.PostUsecase {
	return &PostUsecaseImpl{postRepo}
}

func (u *PostUsecaseImpl) Create(ctx context.Context, post domain.Post) error {
	return u.postRepo.Create(ctx, post)
}

func (u *PostUsecaseImpl) GetAll(ctx context.Context) ([]domain.Post, error) {
	return u.postRepo.GetAll(ctx)
}
