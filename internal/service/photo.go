package service

import (
	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
)

type PhotoService interface {
	// photos
	GetPhotos(ctx *gin.Context) (model.Photo, error)
	GetPhotosById(ctx *gin.Context)
	DeletePhotosById(ctx *gin.Context)
	UpdatePhotosById(ctx *gin.Context)
	AddPhotos(ctx *gin.Context)
}

type photoServiceImpl struct {
	repo repository.PhotoQuery
}
