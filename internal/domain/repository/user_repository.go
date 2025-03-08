package repository

import "github.com/ryota1119/gin_webapi/internal/domain"

// UserRepository UserRepositoryのインターフェースを定義
type UserRepository interface {
	// Create はユーザー情報を作成する
	Create(user *domain.User) error
	// Find はユーザーIDからユーザー情報を取得する
	Find(id domain.UserID) (*domain.User, error)
	// FindByEmail はEmailからユーザー情報を取得する
	FindByEmail(email string) (*domain.User, error)
	// FindByUsername はユーザーネームからユーザー情報を取得する
	FindByUsername(username string) (*domain.User, error)
}
