package service

import (
	//"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
	//"github.com/salamalfis/projectfinal/pkg"
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

// func (u *photoServiceImpl) GetPhotos(ctx *gin.Context) (model.Photo, error) {
// 	// photo, err := u.repo.GetPhotos(ctx)
// 	// if err != nil {
// 	// 	ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Failed to get photos", Error: err.Error()})
// 	// 	return nil, err
// 	// }
// 	// return photo, nil

// }

func (u *photoServiceImpl) GetPhotosById(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
	// 	return
	// }
}

func (u *photoServiceImpl) DeletePhotosById(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
	// 	return
	// }
}

func (u *photoServiceImpl) UpdatePhotosById(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
	// 	return
	// }
}

func (u *photoServiceImpl) AddPhotos(ctx *gin.Context) {
	// photo := model.Photo{}
	// if err := ctx.ShouldBindJSON(&photo); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
	// 	return
	// }
}
