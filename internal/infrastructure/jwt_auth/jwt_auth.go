package jwt_auth

import (
	"errors"
	"github.com/ryota1119/gin_webapi/internal/domain"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtAuth interface {
	GenerateToken(userID domain.UserID, expires time.Duration) (string, string, error)
	ValidateToken(token string) (*jwt.RegisteredClaims, error)
}

type JwtAuthImpl struct {
	secretKey []byte
}

var _ JwtAuth = (*JwtAuthImpl)(nil)

func NewJwtAuth(secretKey []byte) JwtAuth {
	return &JwtAuthImpl{secretKey: secretKey}
}

// GenerateToken はトークンの生成
func (j *JwtAuthImpl) GenerateToken(userID domain.UserID, expires time.Duration) (string, string, error) {
	jti := uuid.New().String()
	log.Printf(jti)
	expirationTime := time.Now().Add(expires)

	claims := &jwt.RegisteredClaims{
		ID:        jti,
		Subject:   userID.String(),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    "gin_webapi",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", "", err
	}
	return signedToken, jti, nil
}

// ValidateToken はトークンの検証
func (j *JwtAuthImpl) ValidateToken(token string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil || !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
