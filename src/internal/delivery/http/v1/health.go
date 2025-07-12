package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthHandler struct {
	// You can add dependencies here if needed
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

// HealthCheck godoc
// @Summary Show server status
// @Description Get the status of server
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/v1/health [get]
func (h *healthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "available",
		"message": "Server is up and running",
		"data": gin.H{
			"version":     "1.0.0",
			"environment": c.MustGet("environment").(string),
		},
	})
}