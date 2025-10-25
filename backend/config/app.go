package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/repository"
	"latih.in-be/internal/route"
	"latih.in-be/internal/service"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

type Controllers struct {
	User      *controller.UserController
	Exam      *controller.ExamController
	Question  *controller.QuestionController
	Option    *controller.OptionController
	Score     *controller.ScoreController
	ExamScore *controller.ExamScoreController
}

func NewApp(db *gorm.DB) *App {
	userRepo := repository.NewUserRepository(db)
	examRepo := repository.NewExamRepository(db)
	questionRepo := repository.NewQuestionRepository(db)
	optionRepo := repository.NewOptionRepository(db)
	scoreRepo := repository.NewScoreRepository(db)
	examScoreRepo := repository.NewExamScoreRepository(db)

	userService := service.NewUserService(userRepo)
	examService := service.NewExamService(examRepo)
	questionService := service.NewQuestionService(questionRepo, userRepo, optionRepo)
	optionService := service.NewOptionService(optionRepo)
	scoreService := service.NewScoreService(scoreRepo)
	examScoreService := service.NewExamScoreService(examScoreRepo)

	controllers := &Controllers{
		User:      controller.NewUserController(userService),
		Exam:      controller.NewExamController(examService),
		Question:  controller.NewQuestionController(questionService),
		Option:    controller.NewOptionController(optionService),
		Score:     controller.NewScoreController(scoreService),
		ExamScore: controller.NewExamScoreController(examScoreService),
	}

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	store := middleware.InMemoryStore(&middleware.InMemoryOptions{
		Rate:  10 * time.Second, // 10s
		Limit: 5,                // maks 5 request
		// Skip: func(c *gin.Context) bool { //skip rate limit
		// 	return c.FullPath() == "/"
		// },
	})

	router.Use(middleware.RateLimiter(store, nil))

	setupRoutes(router, controllers)

	return &App{
		Router: router,
		DB:     db,
	}
}

func setupRoutes(r *gin.Engine, ctrl *Controllers) {
	route.UserRoutes(r, ctrl.User)
	route.ExamRoutes(r, ctrl.Exam)
	route.QuestionRoutes(r, ctrl.Question)
	route.OptionRoutes(r, ctrl.Option)
	route.ScoreRoutes(r, ctrl.Score)
	route.ExamScoreRoutes(r, ctrl.ExamScore)
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}
