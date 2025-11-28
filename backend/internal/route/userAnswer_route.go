package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func UserAnswerRoutes(r *gin.Engine, userAnswer *controller.UserAnswerController) {
	routes := r.Group("/user-answer")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.InputValidateJson([]string{
				"exam_session_id", "user_id",
				"question_id", "answer", "exam_id",
			}), userAnswer.Create)
			auth.GET("/", userAnswer.GetMany)
			auth.GET("/id", userAnswer.GetById)
			auth.PUT("/:id", userAnswer.Update)
			auth.DELETE("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), userAnswer.Delete)
			auth.GET("/session", userAnswer.GetByExamSessionId)
			auth.GET("/question", userAnswer.GetByQuestionId)
			auth.GET("/user", userAnswer.GetUserAnswer)
		}
	}
}
