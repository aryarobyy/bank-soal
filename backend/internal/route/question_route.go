package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func QuestionRoutes(r *gin.Engine, question *controller.QuestionController) {
	routes := r.Group("/question")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidate([]string{"creator_id", "question_text"}), question.Create)
			auth.GET("/", question.GetMany)
			auth.GET("/id", question.GetById)
			auth.PUT("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), question.Update)
			auth.DELETE("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), question.Delete)
			auth.POST("/options", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidate([]string{"creator_id", "question_text"}), question.CreateWithOptions)
			auth.POST("/json", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), question.CreateFromJson)
			auth.GET("/exam", question.GetByExam)
			auth.GET("/diff", question.GetByDiff)
			auth.GET("/creator", question.GetByCreator)
			auth.GET("/subject", question.GetBySubject)
		}
	}
}
