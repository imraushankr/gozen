package utils

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

func NewPagination(page, limit, total int) *Pagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalPage := total / limit
	if total%limit != 0 {
		totalPage++
	}

	return &Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPage,
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}