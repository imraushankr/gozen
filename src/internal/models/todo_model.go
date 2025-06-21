package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoStatus string

const (
    TODO_PENDING    TodoStatus = "pending"
    TODO_IN_PROGRESS TodoStatus = "in_progress" 
    TODO_COMPLETED  TodoStatus = "completed"
)

type TodoPriority string

const (
    PRIORITY_LOW    TodoPriority = "low"
    PRIORITY_MEDIUM TodoPriority = "medium"
    PRIORITY_HIGH   TodoPriority = "high"
)

// Todo struct for SQL databases
type Todo struct {
    BaseModel
    Title       string       `json:"title" gorm:"not null" bson:"title" validate:"required,min=1,max=200"`
    Description string       `json:"description" bson:"description" validate:"max=1000"`
    Status      TodoStatus   `json:"status" gorm:"default:pending" bson:"status"`
    Priority    TodoPriority `json:"priority" gorm:"default:medium" bson:"priority"`
    DueDate     *time.Time   `json:"due_date,omitempty" bson:"due_date,omitempty"`
    CompletedAt *time.Time   `json:"completed_at,omitempty" bson:"completed_at,omitempty"`
    UserID      string       `json:"user_id" gorm:"not null;index" bson:"user_id" validate:"required"`
    Tags        []string     `json:"tags" gorm:"serializer:json" bson:"tags"`
    
    // Relationships
    User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// MongoTodo for MongoDB
type MongoTodo struct {
    MongoBaseModel
    Title       string             `json:"title" bson:"title" validate:"required,min=1,max=200"`
    Description string             `json:"description" bson:"description" validate:"max=1000"`
    Status      TodoStatus         `json:"status" bson:"status"`
    Priority    TodoPriority       `json:"priority" bson:"priority"`
    DueDate     *time.Time         `json:"due_date,omitempty" bson:"due_date,omitempty"`
    CompletedAt *time.Time         `json:"completed_at,omitempty" bson:"completed_at,omitempty"`
    UserID      primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
    Tags        []string           `json:"tags" bson:"tags"`
}

// TodoCreateRequest for creating todos
type TodoCreateRequest struct {
    Title       string       `json:"title" validate:"required,min=1,max=200"`
    Description string       `json:"description" validate:"max=1000"`
    Priority    TodoPriority `json:"priority"`
    DueDate     *time.Time   `json:"due_date,omitempty"`
    Tags        []string     `json:"tags"`
}

// TodoUpdateRequest for updating todos
type TodoUpdateRequest struct {
    Title       *string       `json:"title,omitempty" validate:"omitempty,min=1,max=200"`
    Description *string       `json:"description,omitempty" validate:"omitempty,max=1000"`
    Status      *TodoStatus   `json:"status,omitempty"`
    Priority    *TodoPriority `json:"priority,omitempty"`
    DueDate     *time.Time    `json:"due_date,omitempty"`
    Tags        []string      `json:"tags,omitempty"`
}

// TodoResponse for API responses
type TodoResponse struct {
    ID          string       `json:"id"`
    Title       string       `json:"title"`
    Description string       `json:"description"`
    Status      TodoStatus   `json:"status"`
    Priority    TodoPriority `json:"priority"`
    DueDate     *time.Time   `json:"due_date,omitempty"`
    CompletedAt *time.Time   `json:"completed_at,omitempty"`
    UserID      string       `json:"user_id"`
    Tags        []string     `json:"tags"`
    CreatedAt   time.Time    `json:"created_at"`
    UpdatedAt   time.Time    `json:"updated_at"`
    User        *UserResponse `json:"user,omitempty"`
}