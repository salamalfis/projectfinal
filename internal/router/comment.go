package router

import (
	"github.com/salamalfis/projectfinal/internal/handler"

	"github.com/gin-gonic/gin"
)

type CommentsRouter interface {
	Mount()
}

type CommentsRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.CommentsHandler
}

func (u *CommentsRouterImpl) Mount() {
	u.v.GET("/comments", u.handler.GetComments)
	u.v.POST("/comments?user_id=id", u.handler.AddComments)
	u.v.DELETE("/comments/:id", u.handler.DeleteCommentsById)
	u.v.PUT("/comments/:id", u.handler.UpdateCommentsById)
	u.v.GET("/comments/:id", u.handler.GetCommentsById)
}

// Path: internal/router/router.go
