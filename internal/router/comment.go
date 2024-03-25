package router

import (
	"github.com/salamalfis/projectfinal/internal/handler"
	"github.com/salamalfis/projectfinal/internal/middleware"

	"github.com/gin-gonic/gin"
)

type CommentsRouter interface {
	Mount()
}

type CommentsRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.CommentsHandler
}

func NewCommentsRouter(v *gin.RouterGroup, handler handler.CommentsHandler) CommentsRouter {
	return &CommentsRouterImpl{
		v:       v,
		handler: handler,
		
	}
}

func (u *CommentsRouterImpl) Mount() {
	u.v.Use(middleware.CheckAuthBearer)
	u.v.POST("", u.handler.CreateComments)
	u.v.GET("", u.handler.GetComments)
	u.v.POST("/?photo_id=id", u.handler.CreateComments)
	u.v.DELETE("/:commentId", u.handler.DeleteCommentsById)
	u.v.PUT("/:commentId", u.handler.UpdateCommentsById)
	u.v.GET("/:commentId", u.handler.GetCommentsById)
}



