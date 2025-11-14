package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func OptionRoutes(r *gin.Engine, option *controller.OptionController) {
	routes := r.Group("/option")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidate([]string{
				"question_id", "option_label",
				"option_text", "is_correct",
			}), option.Create)
			auth.GET("/", option.GetMany)
			auth.GET("/id", option.GetById)
			auth.PUT("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), option.Update)
			auth.DELETE("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), option.Delete)
		}
	}
}
