package handler

import (
	"net/http"
	"strconv"

	"github.com/salamalfis/projectfinal/internal/middleware"
	//"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
	"github.com/gin-gonic/gin"
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

func NewPhotoHandler(svc service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{
		svc: svc,
	}
}

// GetPhotos godoc
// @Summary Show photos list
// @Description will fetch 3rd party server to get photos data
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} []model.Photo
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /photos [get]
func (p *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	p.svc.GetPhotos(ctx)
}

// GetPhotosById godoc
// @Summary Show photos detail
// @Description will fetch 3rd party server to get photos data to get detail photo
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
