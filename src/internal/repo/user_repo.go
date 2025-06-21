// package repo

// import (
// 	"context"
// 	"fmt"

// 	"github.com/imraushankr/gozen/src/internal/models"
// 	"github.com/imraushankr/gozen/src/pkg/utils"
// 	"gorm.io/gorm"
// )

// type userRepository struct {
// 	db *gorm.DB
// }

// // NewUserRepository creates a new user repository
// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userRepository{db: db}
// }

// func (r *userRepository) Create(ctx context.Context, user *models.User) error {
// 	return r.db.WithContext(ctx).Create(user).Error
// }

// func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) GetByResetToken(ctx context.Context, token string) (*models.User, error) {
// 	var user models.User
// 	err := r.db.WithContext(ctx).Where("reset_password_token = ?", token).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *userRepository) Update(ctx context.Context, id string, updates map[string]interface{}) error {
// 	return r.db.WithContext(ctx).Model(&models.User{}).
// 		Where("id = ?", id).
// 		Updates(updates).Error
// }

// func (r *userRepository) Delete(ctx context.Context, id string) error {
// 	return r.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
// }

// func (r *userRepository) List(ctx context.Context, params utils.PaginationParams) ([]*models.User, int64, error) {
// 	var users []*models.User
// 	var total int64

// 	query := r.db.WithContext(ctx).Model(&models.User{})

// 	// Apply search if provided
// 	if params.Search != "" {
// 		searchTerm := "%" + params.Search + "%"
// 		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ? OR username ILIKE ?",
// 			searchTerm, searchTerm, searchTerm, searchTerm)
// 	}

// 	// Count total records
// 	if err := query.Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// Apply pagination and sorting
// 	err := query.Order(fmt.Sprintf("%s %s", params.SortBy, params.SortDir)).
// 		Offset(params.CalculateOffset()).
// 		Limit(params.Limit).
// 		Find(&users).Error

// 	return users, total, err
// }

// func (r *userRepository) UpdateRefreshToken(ctx context.Context, id, token string) error {
// 	return r.db.WithContext(ctx).Model(&models.User{}).
// 		Where("id = ?", id).
// 		Update("refresh_token", token).Error
// }

// func (r *userRepository) UpdateResetToken(ctx context.Context, id, token string, expiry int64) error {
// 	updates := map[string]interface{}{
// 		"reset_password_token":  token,
// 		"reset_password_expire": expiry,
// 	}
// 	return r.db.WithContext(ctx).Model(&models.User{}).
// 		Where("id = ?", id).
// 		Updates(updates).Error
// }

// func (r *userRepository) UpdateLastLogin(ctx context.Context, id string) error {
// 	return r.db.WithContext(ctx).Model(&models.User{}).
// 		Where("id = ?", id).
// 		Update("last_login_at", "NOW()").Error
// }

package repo