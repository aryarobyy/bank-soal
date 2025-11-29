package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func ExamSessionRoutes(r *gin.Engine, examSession *controller.ExamSessionController) {
	routes := r.Group("/exam-session")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/", middleware.InputValidateJson([]string{"exam_id"}), examSession.Create)
			auth.GET("/", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), examSession.GetMany)
			auth.GET("/id", examSession.GetById)
			auth.PUT("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), examSession.Update)
			auth.PUT("/:id/no", middleware.InputValidateJson([]string{"current_no"}), examSession.UpdateCurrNo)
			auth.PUT("/finish", middleware.InputValidateJson([]string{"session_id", "user_id"}), examSession.FinishExam)
			auth.DELETE("/:id", middleware.RoleGuard(model.RoleAdmin, model.RoleLecturer), examSession.Delete)
			auth.GET("/score", examSession.GetScore)
		}
	}
}
