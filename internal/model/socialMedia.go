package model

import(
	"time"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID		uint   `gorm:"primaryKey;type:bigint" json:"id"`
	Name	string `gorm:"not null;type:varchar(100)" json:"name"`
	Url		string `gorm:"not null;type:varchar(200)" json:"url"`
	User_id uint `gorm:"not null;type:bigint" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type DefaultColumnSocialMedia struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

