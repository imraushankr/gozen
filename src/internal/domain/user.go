package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleUser    Role = "user"
	RoleManager Role = "manager"
)

var (
	validate = validator.New()
)

type User struct {
	ID                  string     `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName           string     `json:"first_name" validate:"required,min=2,max=50"`
	LastName            string     `json:"last_name" validate:"required,min=2,max=50"`
	Username            string     `json:"username" validate:"required,min=3,max=30,alphanum" gorm:"unique"`
	Role                Role       `json:"role" validate:"required,oneof=admin user manager" gorm:"type:varchar(20)"`
	Email               string     `json:"email" validate:"required,email" gorm:"unique"`
	Phone               string     `json:"phone" validate:"required,min=10,max=15"`
	Avatar              string     `json:"avatar" gorm:"default:https://res.cloudinary.com/cloud-alpha/image/upload/v1739464346/Tours/user_oxe2tu.png"`
	Password            string     `json:"-" validate:"required,min=8"`
	IsActive            bool       `json:"is_active" gorm:"default:true"`
	IsVerified          bool       `json:"is_verified" gorm:"default:false"`
	RefreshToken        string     `json:"-" gorm:"-:all"`
	ResetPasswordToken  string     `json:"-" gorm:"-:all"`
	ResetPasswordExpire *time.Time `json:"-" gorm:"-:all"`
	LastLoginAt         *time.Time `json:"last_login_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt           time.Time  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt           *time.Time `json:"-" gorm:"index"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (u *User) Sanitize() {
	u.Password = ""
	u.RefreshToken = ""
	u.ResetPasswordToken = ""
	u.ResetPasswordExpire = nil
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) HasRole(role Role) bool {
	return u.Role == role
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = RoleUser
	}
	return u.Validate()
}