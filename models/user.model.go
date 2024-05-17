package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Role      string         `json:"role"` // admin,user
	Password  string         `json:"-" gorm:"column:password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
