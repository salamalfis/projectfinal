package model

import (
	"time"
)

// Comment represents the data about a comment
type Comment struct {
	ID        uint64    `gorm:"primaryKey;type:bigint" json:"id"`
	Message   string    `gorm:"not null;type:varchar(200)" json:"message"`
	Photo_id  uint64    `gorm:"not null;type:bigint" json:"photo_id"`
	User_id   uint64    `gorm:"not null;type:bigint" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostComment represents the data about a comment
type CreateComment struct {
	Photo_id uint64 `json:"photo_id" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

// CommentResponse represents the data about a comment
type CreateCommentRespon struct {
	ID        uint64    `json:"id"`
	User_id   uint64    `json:"user_id"`
	Message   string    `json:"message"`
	Photo_id  uint64    `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
}
type CommentView struct {
	ID        uint64    `json:"id"`
	User_id   uint64    `json:"user_id"`
	Photo_id  uint64    `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserItem  `json:"user" gorm:"foreignKey:UserId;references:ID"`
	Photo     PhotoItem `json:"photo" gorm:"foreignKey:PhotoId;references:ID"`
}

type UpdateCommentRes struct {
	ID        uint64    `json:"id"`
	UserId    uint64    `json:"user_id"`
	Message   string    `json:"message"`
	PhotoId   uint64    `json:"photo_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateComment struct {
	Message string `json:"message" validate:"required"`
}
