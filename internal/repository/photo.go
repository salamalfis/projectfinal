package repository

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/infrastructure"
	"github.com/salamalfis/projectfinal/internal/model"
	"gorm.io/gorm"
)

type PhotoQuery interface {
	// photos
	GetPhotos(ctx context.Context, userid uint64) ([]model.Photo, error)
	GetPhotosById(ctx context.Context, id uint64) (model.Photo, error)

	DeletePhotosById(ctx context.Context, id uint64) error
	UpdatePhotosById(ctx context.Context, id uint64, photo model.Photo) error
	AddPhotos(ctx context.Context, photo model.Photo) error
}

// type PhotoCommand interface {
// 	AddPhotos(ctx context.Context, photo model.Photo) error
// 	DeletePhotosByID(ctx context.Context, id uint64) error
// 	UpdatePhotosByID(ctx context.Context, id uint64, photo model.Photo) error
// }

type photoQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewphotoQueryImpl(db infrastructure.GormPostgres) PhotoQuery {
	return &photoQueryImpl{db: db}
}

func (p *photoQueryImpl) AddPhotos(ctx context.Context, photo model.Photo) error {
	db := p.db.GetConnection()

	err := db.
		WithContext(ctx).
		Table("photos").
		Create(&photo).
		Error

	return err
}

func (p *photoQueryImpl) GetPhotos(ctx context.Context, userid uint64) ([]model.Photo, error) {
	db := p.db.GetConnection()

	var photos []model.Photo

	err := db.
		WithContext(ctx).
		Table("photos").
		Where("user_id = ?", userid).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, username").Table("users")
		}).
		Find(&photos).
		Error

	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *photoQueryImpl) GetPhotosById(ctx context.Context, id uint64) (model.Photo, error) {
	db := p.db.GetConnection()
	photos := model.Photo{}
	if err := db.
		WithContext(ctx).
		Table("photos").
		Where("id = ?", id).
		Find(&photos).Error; err != nil {
		// if user not found, return nil error
		if err == gorm.ErrRecordNotFound {
			return model.Photo{}, nil
		}

		return model.Photo{}, err
	}
	return photos, nil
}

func (p *photoQueryImpl) DeletePhotosById(ctx context.Context, id uint64) error {
	db := p.db.GetConnection()

	err := db.
		WithContext(ctx).
		Table("photos").
		Where("id = ?", id).
		Delete(&model.Photo{}).
		Error

	return err
}

func (p *photoQueryImpl) UpdatePhotosById(ctx context.Context, id uint64, photo model.Photo) error {
	db := p.db.GetConnection()

	err := db.
		WithContext(ctx).
		Table("photos").
		Where("id = ?", id).
		Updates(&photo).
		Error

	return err
}
