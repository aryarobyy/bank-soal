package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
	"latih.in-be/utils/helper"
)

func AuthMiddleware(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := helper.ParseAndValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		user, err := userRepo.GetById(c.Request.Context(), claims.UserId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		if user.TokenVersion != claims.TokenVersion {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired, please login again"})
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RoleGuard(allowedRoles ...model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: role info missing"})
			return
		}

		userRole, ok := val.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "system error: invalid role type"})
			return
		}

		if strings.EqualFold(userRole, string(model.RoleSuperAdmin)) {
			c.Next()
			return
		}

		allowed := false
		for _, r := range allowedRoles {
			if strings.EqualFold(string(r), userRole) {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
			return
		}

		c.Next()
	}
}
