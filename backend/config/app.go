package config

import (
	"os"
	"strings"
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
	User        *controller.UserController
	Exam        *controller.ExamController
	Question    *controller.QuestionController
	Option      *controller.OptionController
	ExamScore   *controller.ExamScoreController
	ExamSession *controller.ExamSessionController
	Subject     *controller.SubjectController
	XlsPath     *controller.XlsPathController
	UserAnswer  *controller.UserAnswerController
}

func NewApp(db *gorm.DB) *App {
	router := gin.Default()

	origins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	corsConfig := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	userRepo := repository.NewUserRepository(db)
	examRepo := repository.NewExamRepository(db)
	questionRepo := repository.NewQuestionRepository(db)
	optionRepo := repository.NewOptionRepository(db)
	// examScoreRepo := repository.NewExamScoreRepository(db)
	examSessionRepo := repository.NewExamSessionRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	xlsPathRepo := repository.NewXlsPathRepository(db)
	userAnswerRepo := repository.NewUserAnswerRepository(db)

	userService := service.NewUserService(userRepo)
	examService := service.NewExamService(examRepo, userRepo, questionRepo)
	questionService := service.NewQuestionService(questionRepo, userRepo, optionRepo)
	// optionService := service.NewOptionService(optionRepo)
	// examScoreService := service.NewExamScoreService(examScoreRepo)
	examSessionService := service.NewExamSessionService(examSessionRepo, examRepo, userAnswerRepo, questionRepo)
	subjectService := service.NewSubjectService(subjectRepo)
	xlsPathService := service.NewXlsPathService(xlsPathRepo)
	userAnswerService := service.NewUserAnswerService(userAnswerRepo, optionRepo, examRepo)

	controllers := &Controllers{
		User:     controller.NewUserController(userService, xlsPathService),
		Exam:     controller.NewExamController(examService),
		Question: controller.NewQuestionController(questionService),
		// Option:      controller.NewOptionController(optionService), //gk kepake
		// ExamScore:   controller.NewExamScoreController(examScoreService),
		ExamSession: controller.NewExamSessionController(examSessionService),
		Subject:     controller.NewSubjectController(subjectService),
		XlsPath:     controller.NewXlsPathController(xlsPathService),
		UserAnswer:  controller.NewUserAnswerController(userAnswerService, examSessionService),
	}

	store := middleware.InMemoryStore(&middleware.InMemoryOptions{
		Rate:  30 * time.Second, //deployment aja
		Limit: 15,
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
	r.Static("/storages/images/user", "./storages/images/user")
	protectedStorage := r.Group("/storages")
	protectedStorage.Use(middleware.AuthMiddleware())
	{
		protectedStorage.Static("/images/question", "./storages/images/question")
		protectedStorage.Static("/files", "./storages/files")
	}
	route.UserRoutes(r, ctrl.User)
	route.ExamRoutes(r, ctrl.Exam)
	route.QuestionRoutes(r, ctrl.Question)
	// route.OptionRoutes(r, ctrl.Option)
	// route.ExamScoreRoutes(r, ctrl.ExamScore)
	route.ExamSessionRoutes(r, ctrl.ExamSession)
	route.SubjectRoutes(r, ctrl.Subject)
	route.XlsPathRoutes(r, ctrl.XlsPath)
	route.UserAnswerRoutes(r, ctrl.UserAnswer)
}

func (a *App) Run(addr string) error {
	return a.Router.Run(addr)
}
