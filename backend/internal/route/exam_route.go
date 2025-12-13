package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func ExamRoutes(r *gin.Engine, exam *controller.ExamController) {
	routes := r.Group("/exam")
	{
		routes.OPTIONS("/*path", func(c *gin.Context) {
			c.Status(204)
		})
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer),
				middleware.InputValidateJson([]string{"title", "creator_id", "long_time", "started_at", "finished_at"}),
				exam.Create)
			auth.GET("/", exam.GetMany)
			auth.GET("/id", exam.GetById)
			auth.PUT("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), exam.Update)
			auth.DELETE("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), exam.Delete)
			auth.PUT("/q/add/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidateJson([]string{"question_ids"}), exam.AddQuestions)
			auth.PUT("/q/replace/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidateJson([]string{"question_ids"}), exam.ReplaceQuestions)
			auth.DELETE("/q/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidateJson([]string{"question_ids"}), exam.RemoveQuestions)
			auth.GET("/creator", exam.GetByCreator)
		}
	}
}
