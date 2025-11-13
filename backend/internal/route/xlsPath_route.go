package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func XlsPathRoutes(r *gin.Engine, xlspath *controller.XlsPathController) {
	routes := r.Group("/xlspath")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.GET("/", middleware.RoleGuard("admin"), xlspath.GetMany)
			auth.GET("/id", middleware.RoleGuard("admin"), xlspath.GetById)
			auth.DELETE("/", middleware.RoleGuard("admin"), xlspath.Delete)
		}
	}
}
