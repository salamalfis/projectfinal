package handler

import (

	//"github.com/salamalfis/projectfinal/internal/middleware"
	//"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	//"github.com/salamalfis/projectfinal/pkg"
	"github.com/gin-gonic/gin"
)

type CommentsHandler interface {
	GetComments(c *gin.Context)
	AddComments(c *gin.Context)
	DeleteCommentsById(c *gin.Context)
	UpdateCommentsById(c *gin.Context)
	GetCommentsById(c *gin.Context)
}

type commentsHandlerImpl struct {
	svc service.CommentsService
}

func NewCommentsHandler(svc service.CommentsService) CommentsHandler {
	return &commentsHandlerImpl{
		svc: svc,
	}
}
