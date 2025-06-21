package service

import (
	"context"

	"github.com/imraushankr/gozen/src/internal/models"
	"github.com/imraushankr/gozen/src/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, req *models.UserRegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req *models.UserLoginRequest) (*models.AuthResponse, error)
	RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.AuthResponse, error)
	ForgotPassword(ctx context.Context, req *models.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *models.ResetPasswordRequest) error
	ChangePassword(ctx context.Context, userID string, req *models.ChangePasswordRequest) error
	Logout(ctx context.Context, userID string) error
}

type UserService interface {
	GetProfile(ctx context.Context, userID string) (*models.UserResponse, error)
	UpdateProfile(ctx context.Context, userID string, req *models.UserUpdateRequest) (*models.UserResponse, error)
	DeleteAccount(ctx context.Context, userID string) error
	GetUsers(ctx context.Context, params utils.PaginationParams) ([]models.UserResponse, *utils.PaginationResult, error)
	GetUserByID(ctx context.Context, userID string) (*models.UserResponse, error)
	ActivateUser(ctx context.Context, userID string) error
	DeactivateUser(ctx context.Context, userID string) error
}