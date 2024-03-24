package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment represents the data about a comment
type Comment struct {
	ID        uint           `gorm:"primaryKey;type:bigint" json:"id"`
	Message   string         `gorm:"not null;type:varchar(200)" json:"message"`
	Photo_id  uint           `gorm:"not null;type:bigint" json:"photo_id"`
	User_id   uint           `gorm:"not null;type:bigint" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type DefaultColumnComment struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
