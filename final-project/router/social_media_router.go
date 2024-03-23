package router

import (
	"github.com/gin-gonic/gin"
	"github.com/adie824/final/final-project/app"
	"github.com/adie824/final/final-project/controller"
	"github.com/adie824/final/final-project/middleware"
	"github.com/adie824/final/final-project/repository"
	"github.com/adie824/final/final-project/service"
)

func SocialMediaRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewSocialMediaRepository(db)
	srv := service.NewSocialMediaService(repo)
	ctrl := controller.NewSocialMediaController(srv)

	socialMedia := router.Group("/socialmedias", middleware.AuthMiddleware())

	{
		socialMedia.GET("/", ctrl.GetAll)
		socialMedia.POST("/", ctrl.Create)
		{
			socialMedia.PUT("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Update)
			socialMedia.DELETE("/:id", middleware.SocialMediaMiddleware(srv), ctrl.Delete)
		}
	}
}
