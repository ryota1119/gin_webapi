package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ryota1119/gin_webapi/internal/domain"
	"github.com/ryota1119/gin_webapi/internal/domain/repository"

	"github.com/redis/go-redis/v9"
)

const MaxTokens = 5

type AuthRepositoryImpl struct {
	redisClient *redis.Client
}

func NewAuthRepository(redisClient *redis.Client) repository.AuthRepository {
	return &AuthRepositoryImpl{redisClient}
}

func (r *AuthRepositoryImpl) SaveAccessToken(ctx context.Context, userID domain.UserID, jti string, duration time.Duration) error {
	accessKey := "access:" + jti
	userAccessTokensKey := "user_access_tokens:" + userID.String()

	err := r.redisClient.Set(ctx, accessKey, userID.String(), duration).Err()
	if err != nil {
		return err
	}

	count, err := r.redisClient.LLen(ctx, userAccessTokensKey).Result()
	if err != nil {
		return err
	}
	if count > MaxTokens {
		oldestJTI, err := r.redisClient.RPop(ctx, userAccessTokensKey).Result()
		if err != nil {
			return err
		}

		_ = r.redisClient.Del(ctx, "access:"+oldestJTI).Err()
	}

	return nil
}

func (r *AuthRepositoryImpl) SaveRefreshToken(ctx context.Context, userID domain.UserID, jti string, duration time.Duration) error {
	refreshKey := "refresh:" + jti
	userRefreshTokensKey := "user_refresh_tokens:" + userID.String()

	err := r.redisClient.Set(ctx, refreshKey, userID.String(), duration).Err()
	if err != nil {
		return err
	}

	count, err := r.redisClient.LLen(ctx, userRefreshTokensKey).Result()
	if err != nil {
		return err
	}
	if count > MaxTokens {
		oldestJTI, err := r.redisClient.RPop(ctx, userRefreshTokensKey).Result()
		if err != nil {
			return err
		}

		_ = r.redisClient.Del(ctx, "access:"+oldestJTI).Err()
	}

	return nil
}

func (r *AuthRepositoryImpl) GetUserIDByAccessJti(ctx context.Context, jti string) (*domain.UserID, error) {
	userID, err := r.redisClient.Get(ctx, "access:"+jti).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errors.New("error: Token not found in Redis")
		}
		return nil, err
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("error: failed to convert userID (%s) to int: %w", userID, err)
	}

	returnUserID := domain.UserID(intUserID)

	return &returnUserID, nil
}

func (r *AuthRepositoryImpl) GetUserIDByRefreshToken(ctx context.Context, token string) (string, error) {
	return r.redisClient.Get(ctx, "refresh_token:"+token).Result()
}

func (r *AuthRepositoryImpl) DeleteRefreshToken(ctx context.Context, token string) error {
	return r.redisClient.Del(ctx, "refresh_token:"+token).Err()
}
