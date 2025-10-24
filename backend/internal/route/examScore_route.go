package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func ExamScoreRoutes(r *gin.Engine, examScore *controller.ExamScoreController) {
	examScores := r.Group("/exam-score")
	{
		examScores.POST("/", middleware.InputValidate([]string{"exam_id", "user_id", "status"}), examScore.Create)
		examScores.GET("/", examScore.GetAll)
		examScores.GET("/id", examScore.GetById)
		examScores.PUT("/:id", examScore.Update)
		examScores.DELETE("/:id", examScore.Delete)
	}
}
