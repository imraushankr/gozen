package service

import (
	"context"

	"github.com/imraushankr/gozen/src/internal/models"
	"github.com/imraushankr/gozen/src/pkg/utils"
)

// AuthService interface
type AuthService interface {
	Register(ctx context.Context, req *models.UserRegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req *models.UserLoginRequest) (*models.AuthResponse, error)
	RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.AuthResponse, error)
	Logout(ctx context.Context, userID string) error
	ForgotPassword(ctx context.Context, req *models.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req *models.ResetPasswordRequest) error
	ChangePassword(ctx context.Context, userID string, req *models.ChangePasswordRequest) error
}

// UserService interface
type UserService interface {
	GetProfile(ctx context.Context, userID string) (*models.UserResponse, error)
	UpdateProfile(ctx context.Context, userID string, req *models.UserUpdateRequest) (*models.UserResponse, error)
	DeleteAccount(ctx context.Context, userID string) error
	ListUsers(ctx context.Context, params utils.PaginationParams) ([]*models.UserResponse, int64, error)
	GetUserByID(ctx context.Context, userID string) (*models.UserResponse, error)
}

// TodoService interface
type TodoService interface {
	CreateTodo(ctx context.Context, userID string, req *models.TodoCreateRequest) (*models.TodoResponse, error)
	GetTodo(ctx context.Context, userID, todoID string) (*models.TodoResponse, error)
	UpdateTodo(ctx context.Context, userID, todoID string, req *models.TodoUpdateRequest) (*models.TodoResponse, error)
	DeleteTodo(ctx context.Context, userID, todoID string) error
	ListTodos(ctx context.Context, userID string, params utils.PaginationParams) ([]*models.TodoResponse, int64, error)
	ListAllTodos(ctx context.Context, params utils.PaginationParams) ([]*models.TodoResponse, int64, error)
	GetTodosByStatus(ctx context.Context, userID string, status models.TodoStatus, params utils.PaginationParams) ([]*models.TodoResponse, int64, error)
	GetTodosByPriority(ctx context.Context, userID string, priority models.TodoPriority, params utils.PaginationParams) ([]*models.TodoResponse, int64, error)
	SearchTodos(ctx context.Context, userID, query string, params utils.PaginationParams) ([]*models.TodoResponse, int64, error)
	GetTodoStats(ctx context.Context, userID string) (map[string]int64, error)
}

// EmailService interface
type EmailService interface {
	SendPasswordResetEmail(ctx context.Context, email, token string) error
	SendWelcomeEmail(ctx context.Context, email, name string) error
	SendVerificationEmail(ctx context.Context, email, token string) error
}