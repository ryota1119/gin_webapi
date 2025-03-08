package repository

import (
	"context"
	"database/sql"

	"github.com/ryota1119/gin_webapi/internal/domain"
	"github.com/ryota1119/gin_webapi/internal/domain/repository"
)

type PostRepositoryImpl struct {
	db *sql.DB
}

var _ repository.PostRepository = (*PostRepositoryImpl)(nil)

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &PostRepositoryImpl{db}
}

func (r *PostRepositoryImpl) Create(ctx context.Context, post domain.Post) error {
	return nil
}

func (r *PostRepositoryImpl) GetAll(ctx context.Context) ([]domain.Post, error) {
	var posts []domain.Post
	return posts, nil
}
