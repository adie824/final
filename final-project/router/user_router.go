package router

import (
	"github.com/gin-gonic/gin"
	"github.com/adie824/final/final-project/app"
	"github.com/adie824/final/final-project/controller"
	"github.com/adie824/final/final-project/middleware"
	"github.com/adie824/final/final-project/repository"
	"github.com/adie824/final/final-project/service"
)

func UserRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	ctrl := controller.NewUserController(srv)

	userRouter := router.Group("/users")

	{
		userRouter.POST("/register", ctrl.Register)
		userRouter.POST("/login", ctrl.Login)

		{
			userRouter.PUT("/", middleware.AuthMiddleware(), ctrl.Update)
			userRouter.DELETE("/", middleware.AuthMiddleware(), ctrl.Delete)
		}
	}
}
