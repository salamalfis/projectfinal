package repository

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/infrastructure"
	"github.com/salamalfis/projectfinal/internal/model"
	"gorm.io/gorm"
)

type SocialMediaQuery interface {
	AddSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) error
	GetSocialMedia(ctx context.Context, userId uint64) ([]model.SocialMediaView, error)
	GetSocialMediaById(ctx context.Context, socialId uint64) (*model.SocialMedia, error)
	UpdateSocialMediaById(ctx context.Context, social *model.SocialMedia) error
	DeleteSocialMediaById(ctx context.Context, socialId uint64) error
}

type socialMediaRepositoryImpl struct {
	db infrastructure.GormPostgres
}

func NewSocialMediaQuery(db infrastructure.GormPostgres) SocialMediaQuery {
	return &socialMediaRepositoryImpl{
		db: db,
	}
}

func (s *socialMediaRepositoryImpl) AddSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) error {
	db := s.db.GetConnection()

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Create(socialMedia).
		Error
	return err
}

func (s *socialMediaRepositoryImpl) GetSocialMedia(ctx context.Context, userId uint64) ([]model.SocialMediaView, error) {
	db := s.db.GetConnection()
	socials := []model.SocialMediaView{}

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Where("user_id = ?", userId).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, username").Table("users")
		}).
		Find(&socials).
		Error

	if err != nil {
		return nil, err
	}

	return socials, nil
}

func (s *socialMediaRepositoryImpl) GetSocialMediaById(ctx context.Context, socialId uint64) (*model.SocialMedia, error) {
	db := s.db.GetConnection()
	social := &model.SocialMedia{}

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Where("id = ?", socialId).
		First(social).
		Error

	if err != nil {
		return nil, err
	}

	return social, nil
}

func (s *socialMediaRepositoryImpl) UpdateSocialMediaById(ctx context.Context, social *model.SocialMedia) error {
	db := s.db.GetConnection()

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Where("id = ?", social.ID).
		Updates(social).
		Error

	return err
}

func (s *socialMediaRepositoryImpl) DeleteSocialMediaById(ctx context.Context, socialId uint64) error {
	db := s.db.GetConnection()
	social := model.SocialMedia{ID: socialId}

	err := db.
		WithContext(ctx).
		Table("social_medias").
		Delete(&social).
		Error

	return err
}
