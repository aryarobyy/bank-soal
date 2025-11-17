// helper/image.go

package helper

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}

func DeleteImage(imgUrl string) error {
	if imgUrl == "" {
		return nil
	}

	imgPath := imgUrl
	if strings.Contains(imgUrl, "://") {
		parts := strings.Split(imgUrl, "/")
		for i, part := range parts {
			if part == "storages" {
				imgPath = strings.Join(parts[i:], "/")
				break
			}
		}
	}

	imgPath = filepath.FromSlash(imgPath)

	ext := filepath.Ext(imgPath)
	basePathWithoutExt := strings.TrimSuffix(imgPath, ext)

	extensions := []string{".png", ".jpg", ".jpeg", ".PNG", ".JPG", ".JPEG"}

	deleted := false
	for _, ext := range extensions {
		filePath := basePathWithoutExt + ext
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("failed to delete %s: %w", filePath, err)
			}
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("no image file found to delete")
	}

	return nil
}

func UploadImage(c *gin.Context, id int, imgDir string) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		if err.Error() == "http: no such file" {
			return "", nil
		}
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", fmt.Errorf("invalid file format: %s", ext)
	}

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	var img image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(src)
	} else {
		img, err = png.Decode(src)
	}
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	if err := os.MkdirAll(imgDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	filename := fmt.Sprintf("%d.png", id)
	filePath := filepath.Join(imgDir, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if err := png.Encode(dst, img); err != nil {
		return "", fmt.Errorf("failed to encode PNG: %w", err)
	}

	urlPath := filepath.ToSlash(filePath)

	urlPath = strings.TrimPrefix(urlPath, "./")

	baseURL := GetBaseURL(c)
	fullURL := fmt.Sprintf("%s/%s", baseURL, urlPath)

	return fullURL, nil
}
