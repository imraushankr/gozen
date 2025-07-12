package auth

import (
	"errors"
	"fmt"
	"gozen/src/configs"
	"gozen/src/internal/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateAccessAndRefreshToken(userID, role string, cfg *configs.JWTConfig) (*TokenPair, error) {
	// Generate access token
	accessToken, err := generateToken(
		userID,
		role,
		cfg.AccessTokenSecret,
		cfg.AccessTokenExpiry,
		cfg.Issuer,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token
	refreshToken, err := generateToken(
		userID,
		role,
		cfg.RefreshTokenSecret,
		cfg.RefreshTokenExpiry,
		cfg.Issuer,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(userID, role, secret string, expiry time.Duration, issuer string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		logger.Warn("Failed to parse token", zap.Error(err))
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func RefreshToken(refreshToken string, cfg *configs.JWTConfig) (*TokenPair, error) {
	claims, err := VerifyToken(refreshToken, cfg.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return GenerateAccessAndRefreshToken(claims.UserID, claims.Role, cfg)
}

func GenerateRefreshToken(userId, role string, cfg *configs.JWTConfig) (string, error) {
	refreshToken, err := generateToken(
		userId,
		role,
		cfg.RefreshTokenSecret,
		cfg.RefreshTokenExpiry,
		cfg.Issuer,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return refreshToken, nil
}