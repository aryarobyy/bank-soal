package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func ExamRoutes(r *gin.Engine, exam *controller.ExamController) {
	routes := r.Group("/exam")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard("admin", "lecturer"),
				middleware.InputValidate([]string{"title", "creator_id", "long_time", "started_at", "finished_at"}),
				exam.Create)
			auth.GET("/", exam.GetMany)
			auth.GET("/id", exam.GetById)
			auth.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), exam.Update)
			auth.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), exam.Delete)
		}
	}
}
