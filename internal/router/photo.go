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

func (u *PhotoRouterImpl) Mount() {
	u.v.Use(middleware.CheckAuthBearer)
	u.v.GET("/photo", u.handler.GetPhotos)
	u.v.POST("/photo?user_id=id", u.handler.AddPhotos)
	u.v.DELETE("/photo/:id", u.handler.DeletePhotosById)
	u.v.PUT("/photo/:id", u.handler.UpdatePhotosById)
	u.v.GET("/photo/:id", u.handler.GetPhotosById)
}

// Path: internal/router/router.go
