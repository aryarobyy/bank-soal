package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
	"latih.in-be/middleware"
)

func OptionRoutes(r *gin.Engine, option *controller.OptionController) {
	options := r.Group("/option")
	{
		options.POST("/", middleware.InputValidate([]string{
			"question_id", "option_label",
			"option_text", "is_correct",
		}), option.Create)
		options.GET("/", option.GetAll)
		options.GET("/id/:id", option.GetById)
		options.PUT("/:id", option.Update)
		options.DELETE("/:id", option.Delete)
	}
}
