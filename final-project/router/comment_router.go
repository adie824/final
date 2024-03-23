package router

import (
	"github.com/gin-gonic/gin"
	"github.com/adie824/final/final-project/app"
	"github.com/adie824/final/final-project/controller"
	"github.com/adie824/final/final-project/middleware"
	"github.com/adie824/final/final-project/repository"
	"github.com/adie824/final/final-project/service"
)

func CommentRouter(router *gin.Engine) {
	db := app.NewDB()

	repoPhoto := repository.NewPhotoRepository(db)
	srvPhoto := service.PhotoService(repoPhoto)

	repoComment := repository.NewCommentRepository(db)
	srvComment := service.NewCommentService(repoComment)

	ctrl := controller.NewCommentController(srvComment, srvPhoto)

	commentRouter := router.Group("/comments", middleware.AuthMiddleware())

	{
		commentRouter.POST("/", ctrl.Create)
		commentRouter.GET("/", ctrl.GetAll)
		{
			commentRouter.PUT("/:commentId", middleware.CommentMiddleware(srvComment), ctrl.Update)
			commentRouter.DELETE("/:commentId", middleware.CommentMiddleware(srvComment), ctrl.Delete)
		}
	}
}
