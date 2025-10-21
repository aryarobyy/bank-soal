package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/controller"
	"latih.in-be/middleware"
)

func QuestionRoutes(r *gin.Engine, question *controller.QuestionController) {
	questions := r.Group("/question")
	{
		questions.POST("/", middleware.InputValidate([]string{"creator_id", "question_text"}), question.Create)
		questions.GET("/", question.GetAll)
		questions.GET("/id", question.GetById)
		questions.PUT("/:id", question.Update)
		questions.DELETE("/:id", question.Delete)
		questions.POST("/options", middleware.InputValidate([]string{"creator_id", "question_text"}), question.CreateWithOptions)
		questions.POST("/json", question.CreateFromJson)
		questions.GET("/exam", question.GetByExam)
		questions.GET("/diff", question.GetByCreator)
		questions.GET("/creator", question.GetByCreator)
	}
}
