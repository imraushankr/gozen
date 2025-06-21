package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

// BaseModel contains common fields for all models
type BaseModel struct {
	ID        string     `json:"id" gorm:"primaryKey;type:varchar(36)" bson:"_id,omitempty"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" bson:"deleted_at,omitempty"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}

// MongoBaseModel for MongoDB
type MongoBaseModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time         `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}