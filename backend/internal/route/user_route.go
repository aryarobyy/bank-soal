package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
	"latih.in-be/internal/model"
)

func UserRoutes(r *gin.Engine, user *controller.UserController) {
	routes := r.Group("/user")
	{
		routes.POST("/register", middleware.InputValidate([]string{"name", "password", "major", "faculty"}), user.Register)
		routes.POST("/login", middleware.InputValidate([]string{"email", "password"}), user.Login)

		routes.POST("/refresh", user.RefreshToken)
		routes.GET("/id", user.GetById)

		usersAuth := routes.Group("")
		usersAuth.Use(middleware.AuthMiddleware())
		{
			usersAuth.GET("/email", user.GetByEmail)
			usersAuth.GET("/nim", user.GetByNim)
			usersAuth.GET("/name", user.GetByName)
			usersAuth.GET("/role", user.GetByRole)
			usersAuth.GET("/", user.GetMany)
			usersAuth.PUT("/:id", user.Update)
			usersAuth.PUT("/password", middleware.RoleGuard("admin"), middleware.InputValidate([]string{"new_password"}), user.ChangePassword)
			usersAuth.DELETE("/:id", user.Delete)
			usersAuth.PUT("/role", middleware.RoleGuard(model.RoleAdmin, model.RoleSuperAdmin), middleware.InputValidate([]string{"role"}), user.ChangeRole)
			usersAuth.POST("/generate", middleware.RoleGuard("admin"), user.BulkInsert)
		}
	}
}
