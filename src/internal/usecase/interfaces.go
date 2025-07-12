

package usecase

import (
	"context"
	"gozen/src/internal/domain"
)

type UserUsecase interface {
	Signup(ctx context.Context, user *domain.User) error
	Signin(ctx context.Context, identifier string, password string) (*domain.User, string, error)
	Signout(ctx context.Context, userId string) error
	GetUser(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}