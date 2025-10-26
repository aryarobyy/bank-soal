package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func QuestionRoutes(r *gin.Engine, question *controller.QuestionController) {
	questions := r.Group("/question")
	{
		auth := questions.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			questions.POST("/", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{"creator_id", "question_text"}), question.Create)
			questions.GET("/", question.GetMany)
			questions.GET("/id", question.GetById)
			questions.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), question.Update)
			questions.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), question.Delete)
			questions.POST("/options", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{"creator_id", "question_text"}), question.CreateWithOptions)
			questions.POST("/json", middleware.RoleGuard("admin", "lecturer"), question.CreateFromJson)
			questions.GET("/exam", question.GetByExam)
			questions.GET("/diff", question.GetByDiff)
			questions.GET("/creator", question.GetByCreator)
		}
	}
}
