package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func ExamScoreRoutes(r *gin.Engine, examScore *controller.ExamScoreController) {
	routes := r.Group("/exam-score")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{"exam_id", "user_id", "status"}), examScore.Create)
			auth.GET("/", examScore.GetMany)
			auth.GET("/id", examScore.GetById)
			auth.PUT("/:id", examScore.Update)
			auth.DELETE("/:id", examScore.Delete)
		}
	}
}
