package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func ExamScoreRoutes(r *gin.Engine, examScore *controller.ExamScoreController) {
	routes := r.Group("/exam-score")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), middleware.InputValidate([]string{"exam_id", "user_id", "status"}), examScore.Create)
			auth.GET("/", examScore.GetMany)
			auth.GET("/id", examScore.GetById)
			auth.PUT("/:id", examScore.Update)
			auth.DELETE("/:id", examScore.Delete)
		}
	}
}
