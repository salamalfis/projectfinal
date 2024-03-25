package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
	"github.com/salamalfis/projectfinal/pkg/helper"
)

type CommentsHandler interface {
	GetComments(ctx *gin.Context)
	CreateComments(ctx *gin.Context)
	DeleteCommentsById(ctx *gin.Context)
	UpdateCommentsById(ctx *gin.Context)
	GetCommentsById(ctx *gin.Context)
}

type commentsHandlerImpl struct {
	svc service.CommentsService
	spt service.PhotoService
}

func NewCommentsHandler(svc service.CommentsService, spt service.PhotoService) CommentsHandler {
	return &commentsHandlerImpl{
		svc: svc, spt: spt,
	}
}

func (h *commentsHandlerImpl) CreateComments(ctx *gin.Context) {
	userid, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	commentnew := model.CreateComment{}
	err = ctx.Bind(&commentnew)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(commentnew)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	photo, err := h.spt.GetPhotosById(ctx, commentnew.Photo_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if photo.ID == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "photo not found"})
		return
	}

	commentRes, err := h.svc.CreateComments(ctx, userid, commentnew)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, commentRes)
}

func (h *commentsHandlerImpl) GetComments(ctx *gin.Context) {
	photoIdStr := ctx.Request.URL.Query().Get("photoId")
	if photoIdStr == "" {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Missing Photo id in query"})
		return
	}
	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comments, err := h.svc.GetComments(ctx, uint64(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (h *commentsHandlerImpl) DeleteCommentsById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid id"})
		return
	}
	photo, err := h.svc.GetCommentsById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if photo.ID == 0 {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "photo not found"})
		return
	}
	userid, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if userid != photo.User_id {
		ctx.JSON(http.StatusForbidden, pkg.ErrorResponse{Message: "you are not authorized"})
		return
	}
	err = h.svc.DeleteCommentsById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Message: "Your comment has been successfully deleted"})

}

func (h *commentsHandlerImpl) UpdateCommentsById(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("id"))
	if commentId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment, err := h.svc.GetCommentsById(ctx, uint64(commentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if comment.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Comment did not exist"})
		return
	}

	userId, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint64(comment.ID) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	commentEditData := model.UpdateComment{}
	err = ctx.ShouldBindJSON(&commentEditData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(commentEditData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment.Message = commentEditData.Message

	commentRes, err := h.svc.UpdateCommentsById(ctx, *comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, commentRes)

}

func (h *commentsHandlerImpl) GetCommentsById(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("id"))
	if commentId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment, err := h.svc.GetCommentsById(ctx, uint64(commentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if comment.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Comment did not exist"})
		return
	}

	userId, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint64(comment.User_id) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	commentEditData := model.UpdateComment{}
	err = ctx.ShouldBindJSON(&commentEditData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(commentEditData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	comment.Message = commentEditData.Message

	commentRes, err := h.svc.UpdateCommentsById(ctx, *comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, commentRes)
}
