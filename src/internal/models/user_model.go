package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

// User struct for SQL databases (GORM)
type User struct {
	ID                  string    `json:"id" gorm:"primaryKey;type:varchar(36)" bson:"_id,omitempty"`
	FirstName           string    `json:"first_name" gorm:"not null" bson:"first_name" validate:"required,min=2,max=50"`
	LastName            string    `json:"last_name" gorm:"not null" bson:"last_name" validate:"required,min=2,max=50"`
	Username            string    `json:"username" gorm:"uniqueIndex;not null" bson:"username" validate:"required,min=3,max=30"`
	Role                Role      `json:"role" gorm:"default:USER" bson:"role"`
	Email               string    `json:"email" gorm:"uniqueIndex;not null" bson:"email" validate:"required,email"`
	Phone               string    `json:"phone" bson:"phone" validate:"omitempty,min=10,max=15"`
	Avatar              string    `json:"avatar" bson:"avatar"`
	Password            string    `json:"-" gorm:"not null" bson:"password" validate:"required,min=6"`
	IsVerified          bool      `json:"is_verified" gorm:"default:false" bson:"is_verified"`
	RefreshToken        string    `json:"-" bson:"refresh_token"`
	ResetPasswordToken  string    `json:"-" bson:"reset_password_token"`
	ResetPasswordExpire time.Time `json:"-" bson:"reset_password_expire"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime" bson:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoUpdateTime" bson:"updated_at"`
}

// MongoDB specific user struct
type MongoUser struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName           string             `json:"first_name" bson:"first_name" validate:"required,min=2,max=50"`
	LastName            string             `json:"last_name" bson:"last_name" validate:"required,min=2,max=50"`
	Username            string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Role                Role               `json:"role" bson:"role"`
	Email               string             `json:"email" bson:"email" validate:"required,email"`
	Phone               string             `json:"phone" bson:"phone" validate:"omitempty,min=10,max=15"`
	Avatar              string             `json:"avatar" bson:"avatar"`
	Password            string             `json:"-" bson:"password" validate:"required,min=6"`
	IsVerified          bool               `json:"is_verified" bson:"is_verified"`
	RefreshToken        string             `json:"-" bson:"refresh_token"`
	ResetPasswordToken  string             `json:"-" bson:"reset_password_token"`
	ResetPasswordExpire time.Time          `json:"-" bson:"reset_password_expire"`
	CreatedAt           time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at" bson:"updated_at"`
}