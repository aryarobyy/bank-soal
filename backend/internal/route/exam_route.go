package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func ExamRoutes(r *gin.Engine, exam *controller.ExamController) {
	exams := r.Group("/exam")
	{
		auth := exams.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			exams.POST("/", middleware.RoleGuard("admin", "lecturer"),
				middleware.InputValidate([]string{"title", "creator_id", "long_time", "started_at", "finished_at"}),
				exam.Create)
			exams.GET("/", exam.GetAll)
			exams.GET("/id", exam.GetById)
			exams.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), exam.Update)
			exams.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), exam.Delete)
		}
	}
}
