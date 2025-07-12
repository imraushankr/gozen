

package repository

import (
	"context"
	"gozen/src/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindUser(ctx context.Context, identifier string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}