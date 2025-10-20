package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InputValidate(requiredFields []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]interface{}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		for _, field := range requiredFields {
			if _, ok := body[field]; !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": field + " is required",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
