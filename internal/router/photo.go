package router

import (
	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/handler"
	"github.com/salamalfis/projectfinal/internal/middleware"
)

type PhotoRouter interface {
	Mount()
}

type PhotoRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.PhotoHandler
}

func NewPhotoRouter(v *gin.RouterGroup, handler handler.PhotoHandler) PhotoRouter {
	return &PhotoRouterImpl{
		v:       v,
		handler: handler,
	}
}

func (u *PhotoRouterImpl) Mount() {
	u.v.Use(middleware.CheckAuthBearer)
	u.v.GET("", u.handler.GetPhotos)
	u.v.POST("", u.handler.AddPhotos)
	u.v.POST("/?user_id=id", u.handler.AddPhotos)
	u.v.DELETE("/:id", u.handler.DeletePhotosById)
	u.v.PUT("/:photoId", u.handler.UpdatePhotosById)
	u.v.GET("/:photoId", u.handler.GetPhotosById)
}
