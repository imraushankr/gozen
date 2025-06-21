package repo

import (
	"context"

	"github.com/imraushankr/gozen/src/internal/models"
	"github.com/imraushankr/gozen/src/pkg/utils"
)

// UserRepository interface
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, params utils.PaginationParams) ([]*models.User, int64, error)
	UpdateRefreshToken(ctx context.Context, id, token string) error
	UpdateResetToken(ctx context.Context, id, token string, expiry int64) error
	UpdateLastLogin(ctx context.Context, id string) error
}

// TodoRepository interface
type TodoRepository interface {
	Create(ctx context.Context, todo *models.Todo) error
	GetByID(ctx context.Context, id string) (*models.Todo, error)
	GetByUserID(ctx context.Context, userID string, params utils.PaginationParams) ([]*models.Todo, int64, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, params utils.PaginationParams) ([]*models.Todo, int64, error)
	GetByStatus(ctx context.Context, status models.TodoStatus, params utils.PaginationParams) ([]*models.Todo, int64, error)
	GetByPriority(ctx context.Context, priority models.TodoPriority, params utils.PaginationParams) ([]*models.Todo, int64, error)
	Search(ctx context.Context, query string, params utils.PaginationParams) ([]*models.Todo, int64, error)
	GetStats(ctx context.Context, userID string) (map[string]int64, error)
}