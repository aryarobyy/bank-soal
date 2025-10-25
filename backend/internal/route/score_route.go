package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func ScoreRoutes(r *gin.Engine, score *controller.ScoreController) {
	scores := r.Group("/score")
	{
		auth := scores.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			scores.POST("/", middleware.RoleGuard("admin", "lecturer"), middleware.InputValidate([]string{"exam_id", "user_id", "question_id", "is_correct", "score"}), score.Create)
			scores.GET("/id", score.GetById)
			scores.GET("/all", score.GetAll)
			scores.PUT("/:id", middleware.RoleGuard("admin", "lecturer"), score.Update)
			scores.DELETE("/:id", middleware.RoleGuard("admin", "lecturer"), score.Delete)
		}
	}
}
