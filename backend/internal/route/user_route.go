package route

import (
	"github.com/gin-gonic/gin"
	"latih.in-be/internal/controller"
	"latih.in-be/internal/middleware"
)

func UserRoutes(r *gin.Engine, user *controller.UserController) {
	routes := r.Group("/user")
	{
		routes.POST("/register", middleware.InputValidate([]string{"name", "password", "major", "faculty"}), user.Register)
		routes.POST("/login", middleware.InputValidate([]string{"email", "password"}), user.Login)

		routes.POST("/refresh", user.RefreshToken)

		usersAuth := routes.Group("")
		usersAuth.Use(middleware.AuthMiddleware())
		{
			usersAuth.GET("/id", user.GetById)
			usersAuth.GET("/email", user.GetByEmail)
			usersAuth.GET("/nim", user.GetByNim)
			usersAuth.GET("/name", user.GetByName)
			usersAuth.GET("/role", user.GetByRole)
			usersAuth.GET("/", user.GetMany)
			usersAuth.PUT("/:id", user.Update)
			usersAuth.PUT("/password", middleware.RoleGuard("admin"), middleware.InputValidate([]string{"old_password"}), user.ChangePassword)
			usersAuth.DELETE("/:id", user.Delete)
			usersAuth.PUT("/role", middleware.RoleGuard("admin", "super_admin"), middleware.InputValidate([]string{"role"}), user.ChangeRole)
		}
	}
}
