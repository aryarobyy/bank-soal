package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

func XlsPathRoutes(r *gin.Engine, xlspath *controller.XlsPathController, userRepo repository.UserRepository) {
	routes := r.Group("/xlspath")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware(userRepo))
		{
			auth.GET("/", middleware.RoleGuard(model.RoleAdmin), xlspath.GetMany)
			auth.GET("/id", middleware.RoleGuard(model.RoleAdmin), xlspath.GetById)
			auth.GET("/download", middleware.RoleGuard(model.RoleAdmin), xlspath.Download)
			auth.DELETE("/", middleware.RoleGuard(model.RoleAdmin), xlspath.Delete)
		}
	}
}
