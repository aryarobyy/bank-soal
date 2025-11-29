package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func SubjectRoutes(r *gin.Engine, subject *controller.SubjectController) {
	routes := r.Group("/subject")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidateJson([]string{"title", "code"}), subject.Create)
			auth.GET("/", subject.GetMany)
			auth.GET("/id", subject.GetById)
			auth.GET("/code", subject.GetByCode)
			auth.PUT("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), subject.Update)
			auth.DELETE("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), subject.Delete)
		}
	}
}
