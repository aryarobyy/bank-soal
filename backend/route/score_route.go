package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
	"latih.in-be/middleware"
)

func ScoreRoutes(r *gin.Engine, score *controller.ScoreController) {
	scores := r.Group("/score")
	{
		scores.POST("/", middleware.InputValidate([]string{"exam_id", "user_id", "question_id", "is_correct", "score"}), score.Create)
		scores.GET("/id", score.GetById)
		scores.GET("/all", score.GetAll)
		scores.PUT("/:id", score.Update)
		scores.DELETE("/:id", score.Delete)
	}
}
