package models

import (
	"time"
)

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

// User struct for SQL databases (GORM)
type User struct {
	BaseModel
	FirstName           string     `json:"first_name" gorm:"not null" bson:"first_name" validate:"required,min=2,max=50"`
	LastName            string     `json:"last_name" gorm:"not null" bson:"last_name" validate:"required,min=2,max=50"`
	Username            string     `json:"username" gorm:"uniqueIndex;not null" bson:"username" validate:"required,min=3,max=30"`
	Role                Role       `json:"role" gorm:"default:user" bson:"role"`
	Email               string     `json:"email" gorm:"uniqueIndex;not null" bson:"email" validate:"required,email"`
	Phone               string     `json:"phone" bson:"phone" validate:"omitempty,min=10,max=15"`
	Avatar              string     `json:"avatar" bson:"avatar"`
	Password            string     `json:"-" gorm:"not null" bson:"password" validate:"required,min=6"`
	IsActive            bool       `json:"is_active" gorm:"default:false" bson:"is_active"`
	IsVerified          bool       `json:"is_verified" gorm:"default:false" bson:"is_verified"`
	RefreshToken        string     `json:"-" bson:"refresh_token"`
	ResetPasswordToken  string     `json:"-" bson:"reset_password_token"`
	ResetPasswordExpire time.Time  `json:"-" bson:"reset_password_expire"`
	LastLoginAt         *time.Time `json:"last_login_at,omitempty" bson:"last_login_at,omitempty"`

	// Relationships
	Todos []Todo `json:"todos,omitempty" gorm:"foreignKey:UserID"`
}

// MongoDB specific user struct
type MongoUser struct {
	MongoBaseModel
	FirstName           string     `json:"first_name" bson:"first_name" validate:"required,min=2,max=50"`
	LastName            string     `json:"last_name" bson:"last_name" validate:"required,min=2,max=50"`
	Username            string     `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Role                Role       `json:"role" bson:"role"`
	Email               string     `json:"email" bson:"email" validate:"required,email"`
	Phone               string     `json:"phone" bson:"phone" validate:"omitempty,min=10,max=15"`
	Avatar              string     `json:"avatar" bson:"avatar"`
	Password            string     `json:"-" bson:"password" validate:"required,min=6"`
	IsActive            bool       `json:"is_active" bson:"is_active"`
	IsVerified          bool       `json:"is_verified" bson:"is_verified"`
	RefreshToken        string     `json:"-" bson:"refresh_token"`
	ResetPasswordToken  string     `json:"-" bson:"reset_password_token"`
	ResetPasswordExpire time.Time  `json:"-" bson:"reset_password_expire"`
	LastLoginAt         *time.Time `json:"last_login_at,omitempty" bson:"last_login_at,omitempty"`
}

// UserRegisterRequest for registration
type UserRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Username  string `json:"username" validate:"required,min=3,max=30"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"omitempty,min=10,max=15"`
	Password  string `json:"password" validate:"required,min=6"`
}

// UserLoginRequest for login
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserUpdateRequest for updating user profile
type UserUpdateRequest struct {
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Username  *string `json:"username,omitempty" validate:"omitempty,min=3,max=30"`
	Phone     *string `json:"phone,omitempty" validate:"omitempty,min=10,max=15"`
	Avatar    *string `json:"avatar,omitempty"`
}

// UserResponse for API responses
type UserResponse struct {
	ID          string     `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Username    string     `json:"username"`
	Role        Role       `json:"role"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	Avatar      string     `json:"avatar"`
	IsVerified  bool       `json:"is_verified"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// AuthResponse for authentication responses
type AuthResponse struct {
	User         *UserResponse `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

// RefreshTokenRequest for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// ForgotPasswordRequest for password reset
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPasswordRequest for password reset
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

// ChangePasswordRequest for changing password
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}