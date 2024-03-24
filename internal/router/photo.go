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
	u.v.GET("/photos", u.handler.GetPhotos)
	u.v.POST("/photos", u.handler.AddPhotos)
	u.v.POST("/photos?user_id=id", u.handler.AddPhotos)
	u.v.DELETE("/photos/:id", u.handler.DeletePhotosById)
	u.v.PUT("/photos/:photoId", u.handler.UpdatePhotosById)
	u.v.GET("/photos/:photoId", u.handler.GetPhotosById)
}
