package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
)

func ExamScoreRoutes(r *gin.Engine, examScore *controller.ExamScoreController) {
	examScores := r.Group("/exam-score")
	{
		examScores.POST("/", examScore.Create)
		examScores.GET("/", examScore.GetAll)
		examScores.GET("/id/:id", examScore.GetById)
		examScores.PUT("/:id", examScore.Update)
		examScores.DELETE("/:id", examScore.Delete)
	}
}
