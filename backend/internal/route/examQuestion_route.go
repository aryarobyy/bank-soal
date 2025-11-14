package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func ExamQuestionRoutes(r *gin.Engine, examQ *controller.ExamQuestionController) {
	routes := r.Group("/exam-question")
	{
		routes.OPTIONS("/*path", func(c *gin.Context) {
			c.Status(204)
		})
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), examQ.AddQuestionsToExam)
			auth.PUT("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), examQ.UpdateQuestionsInExam)
			auth.DELETE("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), examQ.RemoveQuestionsFromExam)
		}
	}
}
