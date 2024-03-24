package model

import (
	"time"

	"gorm.io/gorm"
)

// Photo ...

type Photo struct {
	ID        uint           `gorm:"primaryKey;type:bigint" json:"id"`
	Title     string         `gorm:"not null;type:varchar(100)" json:"title"`
	Url 	 string         `gorm:"not null;type:varchar(200)" json:"url"`
	Caption   string         `gorm:"type:varchar(200)" json:"caption"`
	UserId    uint           `gorm:"not null;type:bigint" json:"user_id"`
	Comments  []Comment      `gorm:"foreignKey:PhotoId" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type DefaultColumnPhoto struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
