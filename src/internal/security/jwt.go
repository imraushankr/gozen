package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/imraushankr/gozen/src/internal/config"
	"github.com/imraushankr/gozen/src/internal/models"
)

type JWTClaims struct {
	ID    string      `json:"id"`
	Role  models.Role `json:"role"`
	Email string      `json:"email"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

// GenerateAccessToken generates an access token
func GenerateAccessToken(user *models.User, cfg *config.Config) (string, error) {
	claims := JWTClaims{
		ID:    user.ID,
		Role:  user.Role,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.AccessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.AccessTokenSecret))
}

// GenerateRefreshToken generates a refresh token
func GenerateRefreshToken(userID string, cfg *config.Config) (string, error) {
	claims := RefreshClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.RefreshTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.RefreshTokenSecret))
}

// ValidateAccessToken validates an access token
func ValidateAccessToken(tokenString string, cfg *config.Config) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.AccessTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

// ValidateRefreshToken validates a refresh token
func ValidateRefreshToken(tokenString string, cfg *config.Config) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.RefreshTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}