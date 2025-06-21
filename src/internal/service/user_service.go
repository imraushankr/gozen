// package service

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"github.com/imraushankr/gozen/src/internal/models"
// 	"github.com/imraushankr/gozen/src/internal/repo"
// 	"github.com/imraushankr/gozen/src/internal/security"
// 	"github.com/imraushankr/gozen/src/pkg/logger"
// 	"github.com/imraushankr/gozen/src/pkg/utils"
// 	"go.uber.org/zap"
// )

// type userService struct {
// 	userRepo repo.UserRepository
// }

// // NewUserService creates a new user service
// func NewUserService(userRepo repo.UserRepository) UserService {
// 	return &userService{
// 		userRepo: userRepo,
// 	}
// }

// func (s *userService) GetProfile(ctx context.Context, userID string) (*models.UserResponse, error) {
// 	user, err := s.userRepo.GetByID(ctx, userID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, errors.New("user not found")
// 		}
// 		logger.Error("Failed to get user", zap.Error(err))
// 		return nil, fmt.Errorf("failed to get user: %w", err)
// 	}

// 	return &models.UserResponse{
// 		ID:          user.ID,
// 		FirstName:   user.FirstName,
// 		LastName:    user.LastName,
// 		Username:    user.Username,
// 		Role:        user.Role,
// 		Email:       user.Email,
// 		Phone:       user.Phone,
// 		Avatar:      user.Avatar,
// 		IsVerified:  user.IsVerified,
// 		LastLoginAt: user.LastLoginAt,
// 		CreatedAt:   user.CreatedAt,
// 		UpdatedAt:   user.UpdatedAt,
// 	}, nil
// }

// func (s *userService) UpdateProfile(ctx context.Context, userID string, req *models.UserUpdateRequest) (*models.UserResponse, error) {
// 	// Validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		return nil, fmt.Errorf("validation failed: %w", err)
// 	}

// 	// Get existing user
// 	user, err := s.userRepo.GetByID(ctx, userID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// Check username uniqueness if provided
// 	if req.Username != nil && *req.Username != user.Username {
// 		existingUser, err := s.userRepo.GetByUsername(ctx, *req.Username)
// 		if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 			return nil, fmt.Errorf("failed to check username: %w", err)
// 		}
// 		if existingUser != nil {
// 			return nil, errors.New("username already exists")
// 		}
// 	}

// 	// Update fields
// 	if req.FirstName != nil {
// 		user.FirstName = *req.FirstName
// 	}
// 	if req.LastName != nil {
// 		user.LastName = *req.LastName
// 	}
// 	if req.Username != nil {
// 		user.Username = *req.Username
// 	}
// 	if req.Phone != nil {
// 		user.Phone = *req.Phone
// 	}
// 	if req.Avatar != nil {
// 		user.Avatar = *req.Avatar
// 	}

// 	// Save changes
// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		return nil, fmt.Errorf("failed to update user: %w", err)
// 	}

// 	return &models.UserResponse{
// 		ID:          user.ID,
// 		FirstName:   user.FirstName,
// 		LastName:    user.LastName,
// 		Username:    user.Username,
// 		Role:        user.Role,
// 		Email:       user.Email,
// 		Phone:       user.Phone,
// 		Avatar:      user.Avatar,
// 		IsVerified:  user.IsVerified,
// 		LastLoginAt: user.LastLoginAt,
// 		CreatedAt:   user.CreatedAt,
// 		UpdatedAt:   user.UpdatedAt,
// 	}, nil
// }

// func (s *userService) DeleteAccount(ctx context.Context, userID string) error {
// 	// Check if user exists
// 	_, err := s.userRepo.GetByID(ctx, userID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return errors.New("user not found")
// 		}
// 		return fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// Delete user
// 	if err := s.userRepo.Delete(ctx, userID); err != nil {
// 		return fmt.Errorf("failed to delete user: %w", err)
// 	}

// 	return nil
// }

// func (s *userService) ListUsers(ctx context.Context, params utils.PaginationParams) ([]*models.UserResponse, int64, error) {
// 	users, total, err := s.userRepo.List(ctx, params)
// 	if err != nil {
// 		return nil, 0, fmt.Errorf("failed to list users: %w", err)
// 	}

// 	userResponses := make([]*models.UserResponse, len(users))
// 	for i, user := range users {
// 		userResponses[i] = &models.UserResponse{
// 			ID:          user.ID,
// 			FirstName:   user.FirstName,
// 			LastName:    user.LastName,
// 			Username:    user.Username,
// 			Role:        user.Role,
// 			Email:       user.Email,
// 			Phone:       user.Phone,
// 			Avatar:      user.Avatar,
// 			IsVerified:  user.IsVerified,
// 			LastLoginAt: user.LastLoginAt,
// 			CreatedAt:   user.CreatedAt,
// 			UpdatedAt:   user.UpdatedAt,
// 		}
// 	}

// 	return userResponses, total, nil
// }

// func (s *userService) GetUserByID(ctx context.Context, userID string) (*models.UserResponse, error) {
// 	return s.GetProfile(ctx, userID)
// }


package service