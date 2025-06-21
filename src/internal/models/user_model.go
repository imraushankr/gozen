package models

import (
	"time"
)

/*
Role
- Description: Enum representing available user roles for authorization.
*/
type Role string

const (
	ADMIN Role = "admin" // Admin role with elevated privileges
	USER  Role = "user"  // Regular user role
)

/*
User
  - Description: SQL-compatible User model using GORM.
    Represents registered users with authentication and profile details.
*/
type User struct {
	BaseModel
	FirstName           string     `json:"first_name" gorm:"not null" bson:"first_name" validate:"required,min=2,max=50"`         // First name of the user
	LastName            string     `json:"last_name" gorm:"not null" bson:"last_name" validate:"required,min=2,max=50"`           // Last name of the user
	Username            string     `json:"username" gorm:"uniqueIndex;not null" bson:"username" validate:"required,min=3,max=30"` // Unique username
	Role                Role       `json:"role" gorm:"default:user" bson:"role"`                                                  // Role assigned to the user
	Email               string     `json:"email" gorm:"uniqueIndex;not null" bson:"email" validate:"required,email"`              // User's email address
	Phone               string     `json:"phone" bson:"phone" validate:"omitempty,min=10,max=15"`                                 // Optional phone number
	Avatar              string     `json:"avatar" bson:"avatar"`                                                                  // Profile image URL
	Password            string     `json:"-" gorm:"not null" bson:"password" validate:"required,min=6"`                           // Hashed password (hidden from JSON)
	IsActive            bool       `json:"is_active" gorm:"default:false" bson:"is_active"`                                       // Is the account active
	IsVerified          bool       `json:"is_verified" gorm:"default:false" bson:"is_verified"`                                   // Has the email been verified
	RefreshToken        string     `json:"-" bson:"refresh_token"`                                                                // JWT refresh token (hidden)
	ResetPasswordToken  string     `json:"-" bson:"reset_password_token"`                                                         // Password reset token
	ResetPasswordExpire time.Time  `json:"-" bson:"reset_password_expire"`                                                        // Token expiry
	LastLoginAt         *time.Time `json:"last_login_at,omitempty" bson:"last_login_at,omitempty"`                                // Last login timestamp
}

/*
MongoUser
  - Description: MongoDB-compatible User model using ObjectID.
    Mirrors the SQL model but uses MongoDB-native types.
*/
type MongoUser struct {
	MongoBaseModel
	FirstName           string     `json:"first_name" bson:"first_name" validate:"required,min=2,max=50"` // First name
	LastName            string     `json:"last_name" bson:"last_name" validate:"required,min=2,max=50"`   // Last name
	Username            string     `json:"username" bson:"username" validate:"required,min=3,max=30"`     // Username
	Role                Role       `json:"role" bson:"role"`                                              // Role
	Email               string     `json:"email" bson:"email" validate:"required,email"`                  // Email address
	Phone               string     `json:"phone" bson:"phone" validate:"omitempty,min=10,max=15"`         // Optional phone
	Avatar              string     `json:"avatar" bson:"avatar"`                                          // Avatar URL
	Password            string     `json:"-" bson:"password" validate:"required,min=6"`                   // Hashed password (hidden)
	IsActive            bool       `json:"is_active" bson:"is_active"`                                    // Account status
	IsVerified          bool       `json:"is_verified" bson:"is_verified"`                                // Email verification status
	RefreshToken        string     `json:"-" bson:"refresh_token"`                                        // JWT refresh token
	ResetPasswordToken  string     `json:"-" bson:"reset_password_token"`                                 // Password reset token
	ResetPasswordExpire time.Time  `json:"-" bson:"reset_password_expire"`                                // Token expiry
	LastLoginAt         *time.Time `json:"last_login_at,omitempty" bson:"last_login_at,omitempty"`        // Last login time
}

/*
UserRegisterRequest
- Description: Request payload for registering a new user.
*/
type UserRegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Username  string `json:"username" validate:"required,min=3,max=30"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"omitempty,min=10,max=15"`
	Password  string `json:"password" validate:"required,min=6"`
}

/*
UserLoginRequest
- Description: Request payload for user login.
*/
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

/*
UserUpdateRequest
- Description: Request payload for updating user profile fields.
*/
type UserUpdateRequest struct {
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=2,max=50"`
	Username  *string `json:"username,omitempty" validate:"omitempty,min=3,max=30"`
	Phone     *string `json:"phone,omitempty" validate:"omitempty,min=10,max=15"`
	Avatar    *string `json:"avatar,omitempty"`
}

/*
UserResponse
- Description: Response structure for returning user information.
*/
type UserResponse struct {
	ID          string     `json:"id"`                      // UUID or ObjectID
	FirstName   string     `json:"first_name"`              // First name
	LastName    string     `json:"last_name"`               // Last name
	Username    string     `json:"username"`                // Username
	Role        Role       `json:"role"`                    // Role
	Email       string     `json:"email"`                   // Email
	Phone       string     `json:"phone"`                   // Phone
	Avatar      string     `json:"avatar"`                  // Avatar URL
	IsVerified  bool       `json:"is_verified"`             // Is email verified
	IsActive    bool       `json:"is_active"`               // Is account active
	LastLoginAt *time.Time `json:"last_login_at,omitempty"` // Last login time
	CreatedAt   time.Time  `json:"created_at"`              // Account creation time
	UpdatedAt   time.Time  `json:"updated_at"`              // Last updated time
}

/*
AuthResponse
- Description: Response returned after successful authentication (login/registration).
*/
type AuthResponse struct {
	User         *UserResponse `json:"user"`          // User profile
	AccessToken  string        `json:"access_token"`  // JWT access token
	RefreshToken string        `json:"refresh_token"` // JWT refresh token
}

/*
RefreshTokenRequest
- Description: Request structure to refresh a JWT access token using a refresh token.
*/
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

/*
ForgotPasswordRequest
- Description: Request structure to initiate password reset via email.
*/
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

/*
ResetPasswordRequest
- Description: Request structure to reset the password using a token.
*/
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

/*
ChangePasswordRequest
- Description: Authenticated user password change request structure.
*/
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}