package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
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
}

func NewCommentsHandler(svc service.CommentsService) CommentsHandler {
	return &commentsHandlerImpl{
		svc: svc,
	}
}

func (h *commentsHandlerImpl) GetComments(ctx *gin.Context) {
	comment, err := h.svc.GetComments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (h *commentsHandlerImpl) CreateComments(ctx *gin.Context) {
	comment, err := h.svc.CreateComments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (h *commentsHandlerImpl) DeleteCommentsById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid id"})
		return
	}
	comment, err := h.svc.DeleteCommentsById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comment)

}

func (h *commentsHandlerImpl) UpdateCommentsById(ctx *gin.Context) {
	
}

func (h *commentsHandlerImpl) GetCommentsById(ctx *gin.Context) {
	
}

