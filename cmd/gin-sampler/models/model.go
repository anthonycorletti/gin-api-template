package models

import (
	"time"
)

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	ID        uint      `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// User Model
type User struct {
	Model
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	UserName  string `gorm:"column:user_name" json:"user_name"`
	Email     string `gorm:"column:email" json:"email"`
}

// UserCreate struct for creating users
type UserCreate struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	UserName  string `json:"user_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

// UserUpdate struct for updating users
type UserUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
}
