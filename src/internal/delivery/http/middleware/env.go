package middleware

import (
	"github.com/gin-gonic/gin"
	"gozen/src/configs"
)

func Environment(cfg *configs.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("environment", cfg.App.Environment)
		c.Next()
	}
}