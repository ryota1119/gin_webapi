package repository

import (
	"context"
	"time"

	"github.com/ryota1119/gin_webapi/internal/domain"
)

// AuthRepository AuthRepositoryのインターフェースを定義
type AuthRepository interface {
	// SaveAccessToken はアクセストークンを保存する
	SaveAccessToken(ctx context.Context, userID domain.UserID, jti string, duration time.Duration) error
	// SaveRefreshToken はリフレッシュトークンを保存する
	SaveRefreshToken(ctx context.Context, userID domain.UserID, jti string, duration time.Duration) error
	// GetUserIDByAccessJti はアクセストークンから取得したjtiを使用してユーザーIDを取得する
	GetUserIDByAccessJti(ctx context.Context, jti string) (*domain.UserID, error)
	// GetUserIDByRefreshToken はリフレッシュトークンをから取得したjtiを使用してユーザーIDを取得する
	// TODO 未実装
	GetUserIDByRefreshToken(ctx context.Context, token string) (string, error)
	// DeleteRefreshToken はリフレッシュトークンを削除する
	// TODO 未実装
	DeleteRefreshToken(ctx context.Context, token string) error
}
