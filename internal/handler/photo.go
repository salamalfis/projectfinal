package handler

import (
	//"net/http"
	// "strconv"

	// "github.com/salamalfis/projectfinal/internal/middleware"
	"github.com/gin-gonic/gin"
	//"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	//"github.com/salamalfis/projectfinal/pkg"
)

// PhotoHandler ...
type PhotoHandler interface {
	// photos
	GetPhotos(ctx *gin.Context)
	GetPhotosById(ctx *gin.Context)
	DeletePhotosById(ctx *gin.Context)
	UpdatePhotosById(ctx *gin.Context)
	AddPhotos(ctx *gin.Context)
}

type photoHandlerImpl struct {
	svc service.PhotoService
}

// func NewPhotoHandler(svc service.PhotoService) PhotoHandler {
// 	// return &photoHandlerImpl{
// 	// 	svc: svc,
// 	// }
// }

// GetPhotos godoc
// @Summary Show photos list
// @Description will fetch photos data from the database
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} []model.Photo
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /photos [get]
func (p *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	// photos, err := p.svc.GetPhotosFromDB()
	// if err != nil {
	// 	// handle error
	// 	ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
	// 		Message: "Failed to get photos",
	// 		Error:   err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, photos)
}

// GetPhotosFromDB retrieves photos from the database
// func (s *PhotoService) GetPhotosFromDB() ([]model.Photo, error) {
// 	// implementation to retrieve photos from the database

// }

// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} model.Photo
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /photos/{id} [get]

func (p *photoHandlerImpl) GetPhotosById(ctx *gin.Context) {
	p.svc.GetPhotosById(ctx)
}
