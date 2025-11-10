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
	User         *controller.UserController
	Exam         *controller.ExamController
	Question     *controller.QuestionController
	Option       *controller.OptionController
	ExamScore    *controller.ExamScoreController
	ExamSession  *controller.ExamSessionController
	ExamQuestion *controller.ExamQuestionController
	Subject      *controller.SubjectController
}

func NewApp(db *gorm.DB) *App {
	router := gin.Default()

	// ✅ CORS untuk mengizinkan FE akses BE
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // FE URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	// ✅ Rate limiter tetap digunakan
	store := middleware.InMemoryStore(&middleware.InMemoryOptions{
		Rate:  10 * time.Second,
		Limit: 5,
	})
	router.Use(middleware.RateLimiter(store, nil))

	// Repository
	userRepo := repository.NewUserRepository(db)
	examRepo := repository.NewExamRepository(db)
	questionRepo := repository.NewQuestionRepository(db)
	optionRepo := repository.NewOptionRepository(db)
	examScoreRepo := repository.NewExamScoreRepository(db)
	examSessionRepo := repository.NewExamSessionRepository(db)
	examQuestionRepo := repository.NewExamQuestionRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)

	// Service
	userService := service.NewUserService(userRepo)
	examService := service.NewExamService(examRepo, userRepo)
	questionService := service.NewQuestionService(questionRepo, userRepo, optionRepo)
	optionService := service.NewOptionService(optionRepo)
	examScoreService := service.NewExamScoreService(examScoreRepo)
	examSessionService := service.NewExamSessionService(examSessionRepo, examRepo)
	examQuestionService := service.NewExamQuestionService(examQuestionRepo, questionRepo, examRepo)
	subjectService := service.NewSubjectService(subjectRepo)

	// Controller
	controllers := &Controllers{
		User:         controller.NewUserController(userService),
		Exam:         controller.NewExamController(examService),
		Question:     controller.NewQuestionController(questionService),
		Option:       controller.NewOptionController(optionService),
		ExamScore:    controller.NewExamScoreController(examScoreService),
		ExamSession:  controller.NewExamSessionController(examSessionService),
		ExamQuestion: controller.NewExamQuestionController(examQuestionService),
		Subject:      controller.NewSubjectController(subjectService),
	}

	// Routes
	setupRoutes(router, controllers)

	return &App{
		Router: router,
		DB:     db,
	}
}

func setupRoutes(r *gin.Engine, ctrl *Controllers) {
	r.Static("/storages/images", "./storages/images")
	route.UserRoutes(r, ctrl.User)
	route.ExamRoutes(r, ctrl.Exam)
	route.QuestionRoutes(r, ctrl.Question)
	route.OptionRoutes(r, ctrl.Option)
	route.ExamScoreRoutes(r, ctrl.ExamScore)
	route.ExamSessionRoutes(r, ctrl.ExamSession)
	route.ExamQuestionRoutes(r, ctrl.ExamQuestion)
	route.SubjectRoutes(r, ctrl.Subject)
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}
