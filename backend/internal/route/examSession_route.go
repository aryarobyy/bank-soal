package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func ExamSessionRoutes(r *gin.Engine, examSession *controller.ExamSessionController) {
	routes := r.Group("/exam-session")
	{
		auth := routes.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/",
				middleware.InputValidate([]string{"exam_id"}),
				examSession.Create,
			)
			auth.GET("/",
				middleware.RoleGuard("admin", "lecturer", "super_admin"),
				examSession.GetMany,
			)
			auth.GET("/:id",
				middleware.RoleGuard("admin", "lecturer", "super_admin"),
				examSession.GetById,
			)
			auth.PUT("/:id",
				middleware.RoleGuard("admin", "lecturer"),
				examSession.Update,
			)
			auth.PUT("/:id/current-no",
				middleware.InputValidate([]string{"current_no"}),
				examSession.UpdateCurrNo,
			)
			auth.PUT("/:id/finish",
				examSession.FinishExam,
			)
			auth.DELETE("/:id",
				middleware.RoleGuard("admin", "lecturer"),
				examSession.Delete,
			)
		}
	}
}
