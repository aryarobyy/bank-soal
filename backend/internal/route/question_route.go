package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func QuestionRoutes(r *gin.Engine, question *controller.QuestionController) {
	routes := r.Group("/question")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{"creator_id", "question_text"}), question.Create)
			auth.GET("/", question.GetMany)
			auth.GET("/id", question.GetById)
			auth.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), question.Update)
			auth.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), question.Delete)
			auth.POST("/options", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{"creator_id", "question_text"}), question.CreateWithOptions)
			auth.POST("/json", middleware.RoleGuard("admin", "lecturer"), question.CreateFromJson)
			auth.GET("/exam", question.GetByExam)
			auth.GET("/diff", question.GetByDiff)
			auth.GET("/creator", question.GetByCreator)
			auth.GET("/subject", question.GetBySubject)
		}
	}
}
