package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func OptionRoutes(r *gin.Engine, option *controller.OptionController) {
	options := r.Group("/option")
	{
		auth := options.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			options.POST("/", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{
				"question_id", "option_label",
				"option_text", "is_correct",
			}), option.Create)
			options.GET("/", option.GetAll)
			options.GET("/id", option.GetById)
			options.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), option.Update)
			options.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), option.Delete)
		}
	}
}
