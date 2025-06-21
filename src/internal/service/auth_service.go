// package service

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"time"

// 	"github.com/imraushankr/gozen/src/internal/config"
// 	"github.com/imraushankr/gozen/src/internal/models"
// 	"github.com/imraushankr/gozen/src/internal/repo"
// 	"github.com/imraushankr/gozen/src/internal/security"
// 	"github.com/imraushankr/gozen/src/pkg/logger"
// 	"github.com/imraushankr/gozen/src/pkg/utils"
// 	"go.uber.org/zap"
// )

// type authService struct {
// 	userRepo     repo.UserRepository
// 	config       *config.Config
// 	emailService EmailService
// }

// // NewAuthService creates a new auth service
// func NewAuthService(userRepo repo.UserRepository, cfg *config.Config, emailService EmailService) AuthService {
// 	return &authService{
// 		userRepo:     userRepo,
// 		config:       cfg,
// 		emailService: emailService,
// 	}
// }

// func (s *authService) Register(ctx context.Context, req *models.UserRegisterRequest) (*models.AuthResponse, error) {
// 	// validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		logger.Error("Validation failed for user registration", zap.Error(err))
// 		return nil, fmt.Errorf("validation failed: %w", err)
// 	}

// 	// 1. Check if the user already exists
// 	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
// 	if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 		logger.Error("Failed to check existing user", zap.Error(err), zap.String("email", req.Email))
// 		return nil, fmt.Errorf("failed to check existing user: %w", err)
// 	}

// 	if existingUser != nil {
// 		logger.Warn("User already exists", zap.String("email", req.Email))
// 		return nil, errors.New("user already exists with this email")
// 	}

// 	// 2. Hash the password
// 	hashedPassword, err := security.HashPassword(req.Password)
// 	if err != nil {
// 		logger.Error("failed to hash password", zap.Error(err))
// 		return nil, fmt.Errorf("failed to hash password: %w", err)
// 	}

// 	// 3. Create new user model
// 	user := &models.User{
// 		BaseModel: models.BaseModel{
// 			ID: utils.GenerateUUID(),
// 		},
// 		FirstName:  req.FirstName,
// 		LastName:   req.LastName,
// 		Username:   req.Username,
// 		Email:      req.Email,
// 		Phone:      req.Phone,
// 		Password:   hashedPassword,
// 		Role:       models.USER,
// 		IsActive:   false, // Should be false initially until email verification
// 		IsVerified: false, // Should be false initially until email verification
// 	}

// 	// 4. Save user to the database
// 	if err := s.userRepo.Create(ctx, user); err != nil {
// 		logger.Error("Failed to create user", zap.Error(err))
// 		return nil, fmt.Errorf("failed to create user: %w", err)
// 	}

// 	// 5. Send welcome email (async) - no tokens generated during registration
// 	go func() {
// 		if err := s.emailService.SendWelcomeEmail(user.Email, user.FirstName); err != nil {
// 			logger.Error("Failed to send welcome email", zap.Error(err), zap.String("email", user.Email))
// 		}
// 	}()

// 	// 6. Create user response
// 	userResponse := &models.UserResponse{
// 		ID:         user.ID,
// 		FirstName:  user.FirstName,
// 		LastName:   user.LastName,
// 		Username:   user.Username,
// 		Role:       user.Role,
// 		Email:      user.Email,
// 		Phone:      user.Phone,
// 		Avatar:     user.Avatar,
// 		IsVerified: user.IsVerified,
// 		CreatedAt:  user.CreatedAt,
// 		UpdatedAt:  user.UpdatedAt,
// 	}

// 	// 7. Return response without tokens (tokens only generated during login)
// 	return &models.AuthResponse{
// 		User:         userResponse,
// 		AccessToken:  "", // No access token during registration
// 		RefreshToken: "", // No refresh token during registration
// 	}, nil
// }

// func (s *authService) Login(ctx context.Context, req *models.UserLoginRequest) (*models.AuthResponse, error) {
// 	// Validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		logger.Error("Validation failed for user login", zap.Error(err))
// 		return nil, fmt.Errorf("validation failed: %w", err)
// 	}

// 	// 1. Check if the user exists
// 	user, err := s.userRepo.GetByEmail(ctx, req.Email)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			logger.Warn("User not found", zap.String("email", req.Email))
// 			return nil, errors.New("invalid credentials")
// 		}
// 		logger.Error("Failed to get user", zap.Error(err))
// 		return nil, fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// 2. Check password
// 	if !security.IsPasswordCorrect(req.Password, user.Password) {
// 		logger.Warn("Invalid password", zap.String("email", req.Email))
// 		return nil, errors.New("invalid credentials")
// 	}

// 	// 3. Check if user is active
// 	if !user.IsActive {
// 		logger.Warn("User account is not active", zap.String("email", req.Email))
// 		return nil, errors.New("account is not active")
// 	}

// 	// 4. Generate tokens
// 	accessToken, err := security.GenerateAccessToken(user, s.config)
// 	if err != nil {
// 		logger.Error("Failed to generate access token", zap.Error(err))
// 		return nil, fmt.Errorf("failed to generate access token: %w", err)
// 	}

// 	refreshToken, err := security.GenerateRefreshToken(user.ID, s.config)
// 	if err != nil {
// 		logger.Error("Failed to generate refresh token", zap.Error(err))
// 		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
// 	}

// 	// 5. Update user with refresh token and last login
// 	now := time.Now()
// 	user.RefreshToken = refreshToken
// 	user.LastLoginAt = &now

// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		logger.Error("Failed to update user", zap.Error(err))
// 		return nil, fmt.Errorf("failed to update user: %w", err)
// 	}

// 	// 6. Create user response
// 	userResponse := &models.UserResponse{
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
// 	}

// 	// 7. Return AuthResponse
// 	return &models.AuthResponse{
// 		User:         userResponse,
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 	}, nil
// }

// func (s *authService) RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.AuthResponse, error) {
// 	// Validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		logger.Error("Validation failed for refresh token", zap.Error(err))
// 		return nil, fmt.Errorf("validation failed: %w", err)
// 	}

// 	// 1. Validate refresh token
// 	claims, err := security.ValidateRefreshToken(req.RefreshToken, s.config)
// 	if err != nil {
// 		logger.Error("Invalid refresh token", zap.Error(err))
// 		return nil, errors.New("invalid refresh token")
// 	}

// 	// 2. Get user by ID from token claims
// 	user, err := s.userRepo.GetByID(ctx, claims.ID)
// 	if err != nil {
// 		logger.Error("Failed to get user", zap.Error(err))
// 		return nil, fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// 3. Check if stored refresh token matches
// 	if user.RefreshToken != req.RefreshToken {
// 		logger.Warn("Refresh token mismatch", zap.String("userID", user.ID))
// 		return nil, errors.New("invalid refresh token")
// 	}

// 	// 4. Check if user is active
// 	if !user.IsActive {
// 		logger.Warn("User account is not active", zap.String("userID", user.ID))
// 		return nil, errors.New("account is not active")
// 	}

// 	// 5. Generate new tokens
// 	accessToken, err := security.GenerateAccessToken(user, s.config)
// 	if err != nil {
// 		logger.Error("Failed to generate access token", zap.Error(err))
// 		return nil, fmt.Errorf("failed to generate access token: %w", err)
// 	}

// 	newRefreshToken, err := security.GenerateRefreshToken(user.ID, s.config)
// 	if err != nil {
// 		logger.Error("Failed to generate refresh token", zap.Error(err))
// 		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
// 	}

// 	// 6. Update user with new refresh token
// 	user.RefreshToken = newRefreshToken
// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		logger.Error("Failed to update user", zap.Error(err))
// 		return nil, fmt.Errorf("failed to update user: %w", err)
// 	}

// 	// 7. Create user response
// 	userResponse := &models.UserResponse{
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
// 	}

// 	// 8. Return AuthResponse
// 	return &models.AuthResponse{
// 		User:         userResponse,
// 		AccessToken:  accessToken,
// 		RefreshToken: newRefreshToken,
// 	}, nil
// }

// func (s *authService) Logout(ctx context.Context, userId string) error {
// 	// Get user
// 	user, err := s.userRepo.GetByID(ctx, userId)
// 	if err != nil {
// 		logger.Error("Failed to get user", zap.Error(err))
// 		return fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// Clear refresh token
// 	user.RefreshToken = ""
// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		logger.Error("Failed to clear refresh token", zap.Error(err))
// 		return fmt.Errorf("failed to clear refresh token: %w", err)
// 	}

// 	logger.Info("User logged out successfully", zap.String("userID", userId))
// 	return nil
// }

// func (s *authService) ForgotPassword(ctx context.Context, req *models.ForgotPasswordRequest) error {
// 	// Validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		logger.Error("Validation failed for forgot password", zap.Error(err))
// 		return fmt.Errorf("validation failed: %w", err)
// 	}

// 	// 1. Check if user exists
// 	user, err := s.userRepo.GetByEmail(ctx, req.Email)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			// Don't reveal if email exists or not for security
// 			logger.Info("Password reset requested for non-existent email", zap.String("email", req.Email))
// 			return nil
// 		}
// 		logger.Error("Failed to get user", zap.Error(err))
// 		return fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// 2. Generate reset token using the security package
// 	resetToken, hashedToken, expireTime, err := security.GeneratePasswordResetToken()
// 	if err != nil {
// 		logger.Error("Failed to generate reset token", zap.Error(err))
// 		return fmt.Errorf("failed to generate reset token: %w", err)
// 	}

// 	// 3. Update user with hashed reset token
// 	user.ResetPasswordToken = hashedToken
// 	user.ResetPasswordExpire = expireTime

// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		logger.Error("Failed to save reset token", zap.Error(err))
// 		return fmt.Errorf("failed to save reset token: %w", err)
// 	}

// 	// 4. Send reset email with the original token
// 	if err := s.emailService.SendPasswordResetEmail(user.Email, user.FirstName, resetToken); err != nil {
// 		logger.Error("Failed to send reset email", zap.Error(err))
// 		return fmt.Errorf("failed to send reset email: %w", err)
// 	}

// 	logger.Info("Password reset email sent", zap.String("email", user.Email))
// 	return nil
// }

// func (s *authService) ResetPassword(ctx context.Context, req *models.ResetPasswordRequest) error {
// 	// Validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		logger.Error("Validation failed for reset password", zap.Error(err))
// 		return fmt.Errorf("validation failed: %w", err)
// 	}

// 	// 1. Hash the incoming token to compare with stored hash
// 	hashedToken := security.HashResetToken(req.Token)

// 	// 2. Find user by hashed reset token
// 	user, err := s.userRepo.GetByResetToken(ctx, hashedToken)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			logger.Warn("Invalid reset token", zap.String("token", req.Token))
// 			return errors.New("invalid or expired reset token")
// 		}
// 		logger.Error("Failed to get user by reset token", zap.Error(err))
// 		return fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// 3. Check if token is expired
// 	if time.Now().After(user.ResetPasswordExpire) {
// 		logger.Warn("Reset token expired", zap.String("userID", user.ID))
// 		return errors.New("invalid or expired reset token")
// 	}

// 	// 4. Hash new password
// 	hashedPassword, err := security.HashPassword(req.NewPassword)
// 	if err != nil {
// 		logger.Error("Failed to hash password", zap.Error(err))
// 		return fmt.Errorf("failed to hash password: %w", err)
// 	}

// 	// 5. Update user password and clear reset token
// 	user.Password = hashedPassword
// 	user.ResetPasswordToken = ""
// 	user.ResetPasswordExpire = time.Time{}
// 	user.RefreshToken = "" // Invalidate all sessions

// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		logger.Error("Failed to update password", zap.Error(err))
// 		return fmt.Errorf("failed to update password: %w", err)
// 	}

// 	logger.Info("Password reset successful", zap.String("userID", user.ID))
// 	return nil
// }

// func (s *authService) ChangePassword(ctx context.Context, userId string, req *models.ChangePasswordRequest) error {
// 	// Validate request
// 	if err := security.ValidateStruct(req); err != nil {
// 		logger.Error("Validation failed for change password", zap.Error(err))
// 		return fmt.Errorf("validation failed: %w", err)
// 	}

// 	// 1. Get user
// 	user, err := s.userRepo.GetByID(ctx, userId)
// 	if err != nil {
// 		logger.Error("Failed to get user", zap.Error(err))
// 		return fmt.Errorf("failed to get user: %w", err)
// 	}

// 	// 2. Verify old password
// 	if !security.IsPasswordCorrect(req.OldPassword, user.Password) {
// 		logger.Warn("Invalid old password", zap.String("userID", userId))
// 		return errors.New("invalid old password")
// 	}

// 	// 3. Hash new password
// 	hashedPassword, err := security.HashPassword(req.NewPassword)
// 	if err != nil {
// 		logger.Error("Failed to hash password", zap.Error(err))
// 		return fmt.Errorf("failed to hash password: %w", err)
// 	}

// 	// 4. Update password and clear refresh token (logout from all devices)
// 	user.Password = hashedPassword
// 	user.RefreshToken = ""

// 	if err := s.userRepo.Update(ctx, user); err != nil {
// 		logger.Error("Failed to update password", zap.Error(err))
// 		return fmt.Errorf("failed to update password: %w", err)
// 	}

// 	logger.Info("Password changed successfully", zap.String("userID", userId))
// 	return nil
// }


package service