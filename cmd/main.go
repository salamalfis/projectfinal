package main

import (
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salamalfis/projectfinal/internal/handler"
	"github.com/salamalfis/projectfinal/internal/infrastructure"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
	"github.com/salamalfis/projectfinal/internal/router"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
	"github.com/salamalfis/projectfinal/pkg/helper"

	_ "github.com/salamalfis/projectfinal/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DUCUMENTATION
// @version		2.0
// @description	golong kominfo 006 api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	// requirement technical:
	// [x] middleware untuk recover ketika panic
	// [x] mengecheck basic auth
	server()
}

// Product:
// authorization menggunakan jwt
// authentication bisa dilakukan dengan login
// ketika user login, akan memunculkan JWT ketika success

func server() {
	g := gin.Default()
	g.Use(gin.Recovery())

	// /public => generate JWT public
	g.GET("/public", func(ctx *gin.Context) {
		now := time.Now()

		claim := model.StandardClaim{
			Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
			Iss: "projectfinal",
			Aud: "golang-006",
			Sub: "public-token",
			Exp: uint64(now.Add(time.Hour).Unix()),
			Iat: uint64(now.Unix()),
			Nbf: uint64(now.Unix()),
		}
		token, err := helper.GenerateToken(claim)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
				Message: "error generating public token",
				Errors:  []string{err.Error()},
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]any{"token": token})
	})

	usersGroup := g.Group("/users")
	// usersGroup.Use(middleware.CheckAuthBasic)
	// usersGroup.Use(middleware.CheckAuthBearer)

	// dependency injection
	// dig by uber
	// wire

	// https://s8sg.medium.com/solid-principle-in-go-e1a624290346
	gorm := infrastructure.NewGormPostgres()
	userRepo := repository.NewUserQuery(gorm)
	// userRepoMongo := repository.NewUserQueryMongo()
	userSvc := service.NewUserService(userRepo)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	// mount
	userRouter.Mount()
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// comments
	commentGroup := g.Group("/comments")
	commentRepo := repository.NewCommentQuery(gorm)
	commentSvc := service.NewCommentService(commentRepo)
	commentHdl := handler.NewCommentsHandler(commentSvc)
	commentRouter := router.NewCommentsRouter(commentGroup, commentHdl)
	commentRouter.Mount()

	// Photo
	photoGroup := g.Group("/photos")
	photoRepo := repository.NewPhotoQuery(gorm)
	photoSvc := service.NewPhotoService(photoRepo)
	photoHdl := handler.NewPhotoHandler(photoSvc)
	photoRouter := router.NewPhotoRouter(photoGroup, photoHdl)
	photoRouter.Mount()

	// Social Media
	socialMediaGroup := g.Group("/social-media")
	socialMediaRepo := repository.NewSocialMediaQuery(gorm)
	socialMediaSvc := service.NewSocialMediaService(socialMediaRepo)
	socialMediaHdl := handler.NewSocialMediaHandler(socialMediaSvc)
	socialMediaRouter := router.NewSocialMediaRouter(socialMediaGroup, socialMediaHdl)
	socialMediaRouter.Mount()
	





	g.Run(":3000")

}
