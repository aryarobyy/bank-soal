package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
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
			auth.POST("/", middleware.RoleGuard("admin", "lecturer"), examQ.AddQuestionsToExam)
			auth.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), examQ.UpdateQuestionsInExam)
			auth.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), examQ.RemoveQuestionsFromExam)
		}
	}
}
