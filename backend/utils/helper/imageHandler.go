package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, id int, imgDir string) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", nil
	}

	if err := os.MkdirAll(imgDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create storage directory: %w", err)
	}

	filename := strconv.Itoa(id)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	imageName := fmt.Sprintf("%s%s", filename, ext)

	savePath := filepath.Join(imgDir, imageName)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", fmt.Errorf("failed to save image: %w", err)
	}

	publicPath := strings.TrimPrefix(imgDir, "./")

	imageURL := fmt.Sprintf("http://localhost:8080/%s/%s", publicPath, imageName)

	return imageURL, nil
}

func DeleteImage(imageUrl string) error {
	if imageUrl == "" {
		return nil
	}

	parts := strings.Split(imageUrl, "/")
	imageName := parts[len(parts)-1]

	filePath := filepath.Join("./storages/images", imageName)

	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to delete image file: %w", err)
		}
	}
	return nil
}
