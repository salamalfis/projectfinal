package handler

import (
	"net/http"
	"strconv"

	// "github.com/salamalfis/projectfinal/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
	"github.com/salamalfis/projectfinal/pkg/helper"
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

// AddPhotos godoc
// @Summary Add a photo
// @Description Add a photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param photo body Photo true "Photo"
// @Success 200 {object} Photo
// @Success 201 {object} Photo
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
// @Router /photos [post]

func (h *photoHandlerImpl) AddPhotos(ctx *gin.Context) {
	userId, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	data := model.AddPhotos{}
	err = ctx.Bind(&data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photo := model.Photo{}
	photo.UserId = uint64(userId)
	photo.Title = data.Title
	photo.Url = data.Url
	photo.Caption = data.Caption

	photoRes, err := h.svc.AddPhotos(ctx, photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, photoRes)
}

// GetPhotos godoc
// @Summary Get all photos
// @Description Get all photos
// @Tags photos
// @Accept  json
// @Produce  json
// @Success 200 {array} Photo
// @Success 201 {array} Photo
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
// @Router /photos [get]

func (h *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "userId is required"})
		return
	}

	iduser, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photos, err := h.svc.GetPhotos(ctx, uint64(iduser))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

// GetPhotosById godoc
// @Summary Get a photo by id
// @Description Get a photo by id
// @Tags photos
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} Photo
// @Success 201 {object} Photo
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
// @Router /photos/{id} [get]

func (h *photoHandlerImpl) GetPhotosById(ctx *gin.Context) {
	id := ctx.Request.URL.Query().Get("userId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "userId is required"})
		return
	}

	iduser, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photos, err := h.svc.GetPhotosById(ctx, uint64(iduser))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photos)

}

// UpdatePhotosById godoc
// @Summary Update a photo by id
// @Description Update a photo by id
// @Tags photos
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Param photo body Photo true "Photo"
// @Success 200 {object} Photo
// @Success 201 {object} Photo
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
// @Router /photos/{id} [put]

func (h *photoHandlerImpl) UpdatePhotosById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photo, err := h.svc.GetPhotosById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if photo.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Photo not found"})
		return
	}

	userid, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if photo.UserId != uint64(userid) {
		ctx.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: "You are not authorized to update this photo"})
		return
	}

	data := model.UpdatePhotos{}
	err = ctx.Bind(&data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photoupdate := model.Photo{}
	photoupdate.ID = uint64(id)
	photoupdate.Title = data.Title
	photoupdate.Url = data.Url
	photoupdate.Caption = data.Caption

	photoRes, err := h.svc.UpdatePhotosById(ctx, photoupdate.ID, photoupdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, photoRes)
}

// DeletePhotosById godoc
// @Summary Delete a photo by id
// @Description Delete a photo by id
// @Tags photos
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} Photo
// @Success 201 {object} Photo
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
// @Router /photos/{id} [delete]

func (h *photoHandlerImpl) DeletePhotosById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	photo, err := h.svc.GetPhotosById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if photo.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Photo not found"})
		return
	}

	userid, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if photo.UserId != uint64(userid) {
		ctx.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: "You are not authorized to delete this photo"})
		return
	}

	err = h.svc.DeletePhotosById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
