package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InputValidate(requiredFields []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "error",
				"message": "Failed to read request body: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if len(bodyBytes) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "error",
				"message": "Request body is empty or missing JSON data",
			})
			c.Abort()
			return
		}

		var body map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "error",
				"message": "Invalid JSON format: " + err.Error(),
			})
			c.Abort()
			return
		}

		if len(body) == 0 && len(requiredFields) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"status":  "error",
				"message": "Request body cannot be empty",
			})
			c.Abort()
			return
		}

		for _, field := range requiredFields {
			value, ok := body[field]
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"status":  "error",
					"message": field + " is required",
				})
				c.Abort()
				return
			}

			if value == nil || value == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"status":  "error",
					"message": field + " cannot be empty",
				})
				c.Abort()
				return
			}
		}

		c.Set("body", body)
		c.Next()
	}
}
