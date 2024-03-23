package service

import (
	"github.com/gin-gonic/gin"
)

// CommentsService ...
type CommentsService interface {
	// comments
	GetComments(c *gin.Context)
	AddComments(c *gin.Context)
	DeleteCommentsById(c *gin.Context)
	UpdateCommentsById(c *gin.Context)
	GetCommentsById(c *gin.Context)
}

type commentsServiceImpl struct {
	
}

func NewCommentsService() CommentsService {
	return &commentsServiceImpl{}
}
