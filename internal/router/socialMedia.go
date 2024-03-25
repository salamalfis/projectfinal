package router

import (
	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/handler"
	"github.com/salamalfis/projectfinal/internal/middleware"
)

type SocialMediaRouter interface {
	Mount()
}

type SocialMediaRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.SocialMediaHandler
}

func NewSocialMediaRouter(v *gin.RouterGroup, handler handler.SocialMediaHandler) SocialMediaRouter {
	return &SocialMediaRouterImpl{
		v:       v,
		handler: handler,
	}
}

func (u *SocialMediaRouterImpl) Mount() {
	u.v.Use(middleware.CheckAuthBearer)
	u.v.POST("", u.handler.AddSocialMedia)
	u.v.GET("", u.handler.GetSocialMedia)
	u.v.GET("/:socialMediaId", u.handler.GetSocialMediaById)
	u.v.PUT("/:socialMediaId", u.handler.UpdateSocialMediaById)
	u.v.DELETE("/:socialMediaId", u.handler.DeleteSocialMediaById)

}
