package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ryota1119/gin_webapi/internal/domain/usecase"
	"github.com/ryota1119/gin_webapi/internal/infrastructure/jwt_auth"

	"github.com/ryota1119/gin_webapi/internal/domain"
	"github.com/ryota1119/gin_webapi/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
	jwtAuth  jwt_auth.JwtAuth
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
}

var _ usecase.AuthUsecase = (*AuthUsecaseImpl)(nil)

func NewAuthUsecase(
	jwtAuth jwt_auth.JwtAuth,
	authRepo repository.AuthRepository,
	userRepo repository.UserRepository,
) usecase.AuthUsecase {
	return &AuthUsecaseImpl{
		jwtAuth:  jwtAuth,
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

// Register
func (a *AuthUsecaseImpl) Register(ctx context.Context, username, email, password string) error {
	_, err := a.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     username,
		Email:    email,
		Password: hashedPassword,
	}

	return a.userRepo.Create(user)
}

// Login ログイン
func (a *AuthUsecaseImpl) Login(ctx context.Context, email, password string) (*domain.AuthToken, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// パスワードの比較
	err = checkHashedPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	// アクセストークン生成
	accessToken, accessJti, err := a.jwtAuth.GenerateToken(user.ID, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	// リフレッシュトークン生成
	refreshToken, refreshJti, err := a.jwtAuth.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// アクセストークンをredisに保存
	err = a.authRepo.SaveAccessToken(ctx, user.ID, accessJti, 15*time.Hour)
	if err != nil {
		return nil, err
	}

	// リフレッシュトークンをredisに保存
	err = a.authRepo.SaveRefreshToken(ctx, user.ID, refreshJti, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &domain.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(15 * time.Minute),
	}, nil
}

// RefreshToken で新しい accessToken を発行
func (a *AuthUsecaseImpl) RefreshToken(ctx context.Context, oldToken string) (*domain.AuthToken, error) {
	return &domain.AuthToken{}, nil
}

// ログアウト
func (a *AuthUsecaseImpl) Logout(ctx context.Context, refreshToken string) error {
	return a.authRepo.DeleteRefreshToken(ctx, refreshToken)
}

func (a *AuthUsecaseImpl) ValidateUser(r *http.Request) (*domain.UserID, error) {
	ctx := r.Context()

	// Authorization ヘッダーを取得
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("missing Authorization header")
	}

	// "Bearer <token>" の形式チェック
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid token format")
	}
	accessToken := parts[1]

	// JWT トークンの検証
	claims, err := a.jwtAuth.ValidateToken(accessToken)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	// jti取得
	jti := claims.ID
	tmpUserID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("error: failed to convert userID (%s) to int: %w", tmpUserID, err)
	}
	userID := domain.UserID(tmpUserID)

	// Redis でトークンが有効か確認
	redisUserID, err := a.authRepo.GetUserIDByAccessJti(ctx, jti)
	if errors.Is(err, redis.Nil) {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	// emailからユーザー情報を取得
	if userID != *redisUserID {
		return nil, errors.New("not authorized")
	}

	return &userID, nil
}

// hashPassword パスワードをハッシュ化
func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// checkHashedPassword パスワードを比較
func checkHashedPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
