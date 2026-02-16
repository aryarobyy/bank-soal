package helper

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func Write(c *gin.Context, refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("token cannot be empty")
	}

	c.SetCookie(
		"refresh_token",
		refreshToken,
		3600*24*7,
		"/",
		"",
		false,
		true,
	)

	return nil
}

func Read(c *gin.Context) error {
	token, err := c.Cookie("refresh_token")
	if err != nil {
		return fmt.Errorf("cookie not found: %w", err)
	}

	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	return nil
}
