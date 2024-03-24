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

func (u *SocialMediaRouterImpl) Mount() {
	u.v.Use(middleware.CheckAuthBearer)
	u.v.GET("/socialMedia", u.handler.GetSocialMedia)
	u.v.POST("/socialMedia?user_id=id", u.handler.AddSocialMedia)
	u.v.DELETE("/socialMedia/:id", u.handler.DeleteSocialMediaById)
	u.v.PUT("/socialMedia/:id", u.handler.UpdateSocialMediaById)
	u.v.GET("/socialMedia/:id", u.handler.GetSocialMediaById)
}
