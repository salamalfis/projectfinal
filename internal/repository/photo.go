package repository

import (
	"context"

	//"github.com/salamalfis/projectfinal/internal/infrastructure"
	"github.com/salamalfis/projectfinal/internal/model"
	//"gorm.io/gorm"

)

type PhotoQuery interface {
	// photos
	GetPhotos(ctx context.Context) ([]model.Photo, error)
	GetPhotosByID(ctx context.Context, id uint64) (model.Photo, error)
	
	DeletePhotosByID(ctx context.Context, id uint64) error

}

type PhotoCommand interface {
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
}