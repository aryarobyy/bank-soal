package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/utils/helper"
)

func AuthMiddleware() gin.HandlerFunc {
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

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}
		tokenString := parts[1]

		claims, err := helper.ParseAndValidateToken(tokenString)
		if err == nil {
			c.Set("user", claims)
			c.Set("user_id", claims.UserId)
			c.Set("role", claims.Role)
			c.Next()
			return
		}

		expiredClaims, parseErr := helper.ParseTokenAllowExpired(tokenString)
		if parseErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		refreshToken, cookieErr := c.Cookie("refresh_token")
		if cookieErr != nil || refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired, please login again"})
			return
		}

		_, refreshErr := helper.ValidateRefreshToken(refreshToken)
		if refreshErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired, please login again"})
			return
		}

		newAccessToken, resignErr := helper.ReSignAccessToken(expiredClaims)
		if resignErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to refresh token"})
			return
		}

		c.Header("X-New-Access-Token", newAccessToken)

		c.Set("user", expiredClaims)
		c.Set("user_id", expiredClaims.UserId)
		c.Set("role", expiredClaims.Role)
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
