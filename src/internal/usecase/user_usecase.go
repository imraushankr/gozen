package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gozen/src/configs"
	"gozen/src/internal/domain"
	"gozen/src/internal/pkg/auth"
	"gozen/src/internal/repository"

	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo repository.UserRepository
	cfg      *configs.JWTConfig
}

func NewUserUsecase(userRepo repository.UserRepository, cfg *configs.JWTConfig) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

// func (uc *userUsecase) Signup(ctx context.Context, user *domain.User) error {
// 	if user.Email == "" || user.Username == "" || user.Password == "" {
// 		return domain.ErrInvalidInput
// 	}

// 	// existingUser, err := uc.userRepo.FindByEmail(ctx, user.Email)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to check email existence: %w", err)
// 	// }
// 	// if existingUser != nil {
// 	// 	return domain.ErrEmailAlreadyExists
// 	// }

// 	existingUser, err := uc.userRepo.FindByEmail(ctx, user.Email)
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return fmt.Errorf("failed to check email existence: %w", err)
// 	}
// 	if existingUser != nil && existingUser.Email == user.Email {
// 		return domain.ErrEmailAlreadyExists
// 	}

// 	existingUser, err = uc.userRepo.FindByUsername(ctx, user.Username)
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return fmt.Errorf("failed to check username existence: %w", err)
// 	}
// 	if existingUser != nil && existingUser.Username == user.Username {
// 		return domain.ErrUsernameAlreadyExists
// 	}

// 	hashPassword, err := auth.HashPassword(user.Password)
// 	if err != nil {
// 		return fmt.Errorf("failed to hash password: %w", err)
// 	}
// 	user.Password = hashPassword

// 	if user.Role == "" {
// 		user.Role = domain.RoleUser
// 	}

// 	now := time.Now()
// 	user.CreatedAt = now
// 	user.UpdatedAt = now

// 	if err := uc.userRepo.Create(ctx, user); err != nil {
// 		return fmt.Errorf("failed to create user: %w", err)
// 	}

// 	return nil
// }

func (uc *userUsecase) Signup(ctx context.Context, user *domain.User) error {
	if user.Email == "" || user.Username == "" || user.Password == "" || user.Phone == "" {
		return domain.ErrInvalidInput
	}

	existingUser, err := uc.userRepo.FindByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check email existence: %w", err)
	}
	if existingUser != nil && existingUser.Email == user.Email {
		return domain.ErrEmailAlreadyExists
	}

	existingUser, err = uc.userRepo.FindByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check username existence: %w", err)
	}
	if existingUser != nil && existingUser.Username == user.Username {
		return domain.ErrUsernameAlreadyExists
	}

	hashPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashPassword

	if user.Role == "" {
		user.Role = domain.RoleUser
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (uc *userUsecase) Signin(ctx context.Context, identifier string, password string) (*domain.User, string, error) {
	if identifier == "" || password == "" {
		return nil, "", domain.ErrInvalidInput
	}

	user, err := uc.userRepo.FindUser(ctx, identifier)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, "", domain.ErrInvalidCredentials
	}

	if err := auth.CheckPasswordHash(password, user.Password); err != nil {
		return nil, "", domain.ErrInvalidCredentials
	}

	tokenPair, err := auth.GenerateAccessAndRefreshToken(user.ID, string(user.Role), uc.cfg)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate tokens: %w", err)
	}

	now := time.Now()
	user.LastLoginAt = &now
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, "", fmt.Errorf("failed to update last login: %w", err)
	}

	user.RefreshToken = tokenPair.RefreshToken

	return user, tokenPair.AccessToken, nil
}

func (uc *userUsecase) Signout(ctx context.Context, userId string) error {
	return nil
}

func (uc *userUsecase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	if user == nil || user.ID == "" {
		return domain.ErrInvalidInput
	}

	existingUser, err := uc.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}
	if existingUser == nil {
		return domain.ErrUserNotFound
	}

	user.UpdatedAt = time.Now()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return domain.ErrInvalidInput
	}

	existingUser, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}
	if existingUser == nil {
		return domain.ErrUserNotFound
	}

	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (uc *userUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	if refreshToken == "" {
		return "", domain.ErrInvalidInput
	}

	claims, err := auth.VerifyToken(refreshToken, uc.cfg.RefreshTokenSecret)
	if err != nil {
		return "", domain.ErrInvalidToken
	}

	tokenPair, err := auth.GenerateAccessAndRefreshToken(claims.UserID, claims.Role, uc.cfg)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}

	return tokenPair.AccessToken, nil
}
