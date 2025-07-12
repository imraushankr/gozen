package middleware

import (
	"gozen/src/internal/pkg/auth"
	"gozen/src/internal/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("Authorization header is missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			logger.Warn("Bearer token not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			return
		}

		claims, err := auth.VerifyToken(tokenString, jwtSecret)
		if err != nil {
			logger.Warn("Invalid token", zap.Error(err))
			status := http.StatusUnauthorized
			if err == auth.ErrExpiredToken {
				status = http.StatusForbidden
			}
			c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}

		// Add user info to context
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}
