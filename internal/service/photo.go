package service

import (
	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
)

type PhotoService interface {
	// photos
	GetPhotos(ctx *gin.Context, id uint64) ([]model.Photo, error)
	GetPhotosById(ctx *gin.Context, id uint64) (model.Photo, error)
	DeletePhotosById(ctx *gin.Context, id uint64) error
	UpdatePhotosById(ctx *gin.Context, id uint64, photo model.Photo) (model.PhotoUpdate, error)
	AddPhotos(ctx *gin.Context, photo model.Photo) (model.PhotoResCreate, error)
}

type photoServiceImpl struct {
	repo repository.PhotoQuery
}

func NewPhotoService(repo repository.PhotoQuery) PhotoService {
	return &photoServiceImpl{
		repo: repo,
	}
}

func (s *photoServiceImpl) AddPhotos(ctx *gin.Context, photo model.Photo) (model.PhotoResCreate, error) {
	err := s.repo.AddPhotos(ctx, photo)
	if err != nil {
		return model.PhotoResCreate{}, err
	}
	return model.PhotoResCreate{}, nil
}

func (s *photoServiceImpl) GetPhotos(ctx *gin.Context, id uint64) ([]model.Photo, error) {
	photos, err := s.repo.GetPhotos(ctx, id)
	if err != nil {
		return nil, err
	}
	return photos, nil
}

func (s *photoServiceImpl) GetPhotosById(ctx *gin.Context, id uint64) (model.Photo, error) {
	photo, err := s.repo.GetPhotosById(ctx, id)
	if err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (s *photoServiceImpl) DeletePhotosById(ctx *gin.Context, id uint64) error {
	err := s.repo.DeletePhotosById(ctx, id)

	return err
}

func (s *photoServiceImpl) UpdatePhotosById(ctx *gin.Context, id uint64, photo model.Photo) (model.PhotoUpdate, error) {
	err := s.repo.UpdatePhotosById(ctx, id, photo)
	if err != nil {
		return model.PhotoUpdate{}, err
	}
	return model.PhotoUpdate{}, nil
}
