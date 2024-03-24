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
	u.v.GET("/socialmedias", u.handler.GetSocialMedia)
	u.v.POST("/socialmedias", u.handler.AddSocialMedia)
	u.v.POST("/socialmedias?user_id=id", u.handler.AddSocialMedia)
	u.v.DELETE("/socialmedias/:socialMediaId", u.handler.DeleteSocialMediaById)
	u.v.PUT("/socialmedias/:socialMediaId", u.handler.UpdateSocialMediaById)

}
