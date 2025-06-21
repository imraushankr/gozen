package utils

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/imraushankr/gozen/src/pkg/response"
)

type PaginationParams struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
	Search  string `json:"search"`
}

// GetPaginationParams extracts pagination parameters from query string
func GetPaginationParams(c *fiber.Ctx) PaginationParams {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sort_by", "created_at")
	sortDir := c.Query("sort_dir", "desc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if sortDir != "asc" && sortDir != "desc" {
		sortDir = "desc"
	}

	return PaginationParams{
		Page:    page,
		Limit:   limit,
		SortBy:  sortBy,
		SortDir: sortDir,
		Search:  search,
	}
}

// CalculateOffset calculates the offset for pagination
func (p *PaginationParams) CalculateOffset() int {
	return (p.Page - 1) * p.Limit
}

// CreateMeta creates pagination metadata
func (p *PaginationParams) CreateMeta(total int64) *response.Meta {
	totalPages := int(math.Ceil(float64(total) / float64(p.Limit)))
	return &response.Meta{
		Page:       p.Page,
		Limit:      p.Limit,
		Total:      total,
		TotalPages: totalPages,
	}
}