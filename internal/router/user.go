package router

import (
	"github.com/salamalfis/projectfinal/internal/handler"
	"github.com/salamalfis/projectfinal/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

func (u *userRouterImpl) Mount() {
	// activity
	// /users/sign-up
	u.v.POST("/register", u.handler.UserSignUp)

	// /users/login
	u.v.POST("/login", u.handler.UserLogin)

	// users
	u.v.Use(middleware.CheckAuthBearer)
	// /users
	u.v.GET("", u.handler.GetUsers)
	// /users/:id
	u.v.GET("/:id", u.handler.GetUsersById)
	u.v.DELETE("/:id", u.handler.DeleteUsersById)
}
