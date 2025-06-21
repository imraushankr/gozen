package repo

import (
	"context"

	"github.com/imraushankr/gozen/src/pkg/utils"
)

type UserRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, user interface{}) error
	GetByID(ctx context.Context, id string) (interface{}, error)
	GetByEmail(ctx context.Context, email string) (interface{}, error)
	GetByUsername(ctx context.Context, username string) (interface{}, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) error
	Delete(ctx context.Context, id string) error

	// Advanced queries
	List(ctx context.Context, params utils.PaginationParams) ([]interface{}, int64, error)
	UpdateRefreshToken(ctx context.Context, id, token string) error
	UpdatePasswordResetToken(ctx context.Context, email, token string, expiry interface{}) error
	GetByPasswordResetToken(ctx context.Context, token string) (interface{}, error)
	ClearPasswordResetToken(ctx context.Context, id string) error
	UpdateLastLogin(ctx context.Context, id string) error

	// User-specific operations
	ActivateUser(ctx context.Context, id string) error
	VerifyUser(ctx context.Context, id string) error
	IsEmailExists(ctx context.Context, email string) (bool, error)
	IsUsernameExists(ctx context.Context, username string) (bool, error)
}