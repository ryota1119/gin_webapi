package usecase

import (
	"context"
	"net/http"

	"github.com/ryota1119/gin_webapi/internal/domain"
)

// AuthUsecase AuthUsecaseのインターフェースを定義
type AuthUsecase interface {
	// ValidateUser 認証ミドルウェアでリクエストのトークンを読み取って検証を行う
	ValidateUser(r *http.Request) (*domain.UserID, error)
	// Register はユーザー情報を作成する
	Register(ctx context.Context, username, email, password string) error
	// Login はログインを行う
	Login(ctx context.Context, email, password string) (*domain.AuthToken, error)
	// RefreshToken は
	// TODO 未実装
	RefreshToken(ctx context.Context, oldToken string) (*domain.AuthToken, error)
	// Logout はログアウトを行う
	// TODO 未実装
	Logout(ctx context.Context, refreshToken string) error
}
