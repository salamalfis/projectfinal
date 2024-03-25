package model

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint64    `json:"id"`
	UserId         uint64    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      gorm.DeletedAt
}

type NewSocialMedia struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type CreateSocialMediaRes struct {
	ID             uint64    `json:"id"`
	UserId         uint64    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateSocialMediaRes struct {
	ID             uint64    `json:"id"`
	UserId         uint64    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaView struct {
	ID             uint64    `json:"id"`
	UserId         uint64    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UserItem  `json:"user" gorm:"foreignKey:UserId;references:ID"`
}
