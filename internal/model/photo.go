package model

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint64         `gorm:"primaryKey;type:bigint" json:"id"`
	Title     string         `gorm:"not null;type:varchar(100)" json:"title"`
	Url       string         `gorm:"not null;type:varchar(200)" json:"url"`
	Caption   string         `gorm:"type:varchar(200)" json:"caption"`
	UserId    uint64         `gorm:"not null;type:bigint" json:"user_id"`
	Comments  []Comment      `gorm:"foreignKey:PhotoId" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type AddPhotos struct {
	Title   string `json:"title" binding:"required"`
	Url     string `json:"url" binding:"required"`
	Caption string `json:"caption"`
}

type UpdatePhotos struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Caption string `json:"caption"`
}

type ViewPhotos struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Caption   string    `json:"caption"`
	UserId    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      UserItem  `json:"user" gorm:"foreignKey:UserId;references:ID"`
}

type PhotoResCreate struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"url"`
	UserId    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePhotosById struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Url       string    `json:"url"`
	UserId    uint64    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoUpdate struct {
	ID        uint32    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint32    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostComment struct {
	ID       uint64    `json:"id"`
	UserId   uint64    `json:"user_id"`
	Message  string    `json:"message"`
	PhotoId  uint64    `json:"photo_id"`
	CreateAt time.Time `json:"created_at"`
}

type PhotoItem struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint64 `json:"user_id"`
}
