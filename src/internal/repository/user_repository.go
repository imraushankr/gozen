package repository

import (
	"context"
	"gozen/src/internal/domain"
	"gozen/src/internal/pkg/database"
)

type userRepository struct {
	db *database.DB
}

func NewUserRespository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	result := r.db.GORM.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	result := r.db.GORM.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	result := r.db.GORM.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	result := r.db.GORM.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindUser(ctx context.Context, identifier string) (*domain.User, error) {
	var user domain.User
	result := r.db.GORM.WithContext(ctx).Where("id = ? OR email = ? OR username = ?", identifier, identifier, identifier).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	result := r.db.GORM.WithContext(ctx).Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	result := r.db.GORM.WithContext(ctx).Where("id = ?", id).Delete(&domain.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}