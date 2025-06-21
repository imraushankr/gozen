package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

/*
BaseModel
  - Description: BaseModel is a shared structure used across all SQL-based GORM models.
    It includes common fields like ID, timestamps, and soft-delete functionality.
*/
type BaseModel struct {
	ID        string     `json:"id" gorm:"primaryKey;type:varchar(36)" bson:"_id,omitempty"`    // Unique UUID identifier
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime" bson:"created_at"`            // Timestamp of creation
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime" bson:"updated_at"`            // Timestamp of last update
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" bson:"deleted_at,omitempty"` // Soft delete field
}

/*
BeforeCreate
- Description: GORM hook to automatically assign a UUID to the ID field before record creation.
*/
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}

/*
MongoBaseModel
  - Description: MongoBaseModel provides the base structure for MongoDB documents,
    including ObjectID and timestamp fields.
*/
type MongoBaseModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`                          // MongoDB ObjectID
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`                     // Document creation timestamp
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`                     // Document update timestamp
	DeletedAt *time.Time         `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"` // Soft delete (optional)
}