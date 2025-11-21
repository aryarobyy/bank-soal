package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
	"latih.in-be/utils/response"
)

type XlsPathController struct {
	service service.XlsPathService
}

func NewXlsPathController(s service.XlsPathService) *XlsPathController {
	return &XlsPathController{service: s}
}

func (h *XlsPathController) GetById(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid data id")
		return
	}

	data, err := h.service.GetById(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	xlsPathRes := response.XlsPathResponse(*data)

	helper.Success(c, xlsPathRes, "data found")
}

func (h *XlsPathController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit and offset")
		return
	}

	data, err := h.service.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	xlsPathsRes := response.XlsPathsResponse(data)

	helper.Success(c, xlsPathsRes, "data found")
}

func (h *XlsPathController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.service.Delete(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, nil, "data deleted")
}

func (h *XlsPathController) Download(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	record, err := h.service.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "file metadata not found"})
		return
	}

	filePath := record.FilePath
	cleanPath := filepath.Clean(filePath)

	baseDir := filepath.Clean("./storages/files")
	absBase, _ := filepath.Abs(baseDir)
	absPath, _ := filepath.Abs(cleanPath)

	if !strings.HasPrefix(absPath, absBase) {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid file path"})
		return
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	filename := filepath.Base(absPath)

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	c.File(absPath)
}
