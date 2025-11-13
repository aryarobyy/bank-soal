package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
)

type UserController struct {
	service        service.UserService
	xlsPathService service.XlsPathService
}

func NewUserController(s service.UserService, x service.XlsPathService) *UserController {
	return &UserController{
		service:        s,
		xlsPathService: x,
	}
}

func (h *UserController) Register(c *gin.Context) {
	var user model.RegisterCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	rules := map[string]int{
		"name":    256,
		"email":   512,
		"faculty": 128,
		"major":   256,
	}

	if err := helper.ValidateFieldLengths(user, rules); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if !helper.IsValidEmail(user.Email) {
		helper.Error(c, http.StatusBadRequest, "wrong email format")
		return
	}

	if err := h.service.Register(c.Request.Context(), user); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, user, "user registered")
}

func (h *UserController) Login(c *gin.Context) {
	var cred model.LoginCredential
	if err := c.ShouldBindJSON(&cred); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if !helper.IsValidEmail(cred.Email) {
		helper.Error(c, http.StatusBadRequest, "wrong email format")
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(c.Request.Context(), cred)
	if err != nil {
		helper.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := helper.Write(c, refreshToken); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to set cookie")
		return
	}

	helper.Success(c, user, "login successful", accessToken)
}

func (h *UserController) GetById(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, user, "user found")
}

func (h *UserController) GetByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		helper.Error(c, http.StatusBadRequest, "invalid user email")
		return
	}
	if len(email) > 512 {
		helper.Error(c, http.StatusBadRequest, "invalid email")
		return
	}

	if !helper.IsValidEmail(email) {
		helper.Error(c, http.StatusBadRequest, "wrong email format")
		return
	}

	user, err := h.service.GetByEmail(c.Request.Context(), email)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, user, "user found")
}

func (h *UserController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	email := c.PostForm("email")

	if email != "" && !helper.IsValidEmail(email) {
		helper.Error(c, http.StatusBadRequest, "wrong email format")
		return
	}

	nim := c.PostForm("nim")
	nip := c.PostForm("nip")
	nidn := c.PostForm("nidn")

	var nimPtr, nipPtr, nidnPtr *string
	if nim != "" {
		nimPtr = &nim
	}
	if nip != "" {
		nipPtr = &nip
	}
	if nidn != "" {
		nidnPtr = &nidn
	}

	user := model.User{
		Name:         c.PostForm("name"),
		Email:        email,
		Nim:          nimPtr,
		Nip:          nipPtr,
		Nidn:         nidnPtr,
		Role:         model.Role(c.PostForm("role")),
		Major:        c.PostForm("major"),
		Faculty:      c.PostForm("faculty"),
		Status:       model.Status(c.PostForm("status")),
		AcademicYear: c.PostForm("academic_year"),
	}

	file, _ := c.FormFile("image")
	if file != nil {
		imageUrl, err := helper.UploadImage(c, id)
		if err != nil {
			helper.Error(c, http.StatusInternalServerError, "failed to upload image")
			return
		}
		user.ImgUrl = imageUrl
	}
	updatedUser, err := h.service.Update(c.Request.Context(), user, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Success(c, updatedUser, "user updated successfully")
}

func (h *UserController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}
	helper.Success(c, nil, "user deleted")

}

func (h *UserController) GetMany(c *gin.Context) {
	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}
	users, total, err := h.service.GetMany(c, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, gin.H{"data": users, "total": total}, "users found")
}

func (h *UserController) GetByNim(c *gin.Context) {
	nim := c.Query("nim")
	user, err := h.service.GetByNim(c, nim)
	if len(nim) >= 10 {
		helper.Error(c, http.StatusBadRequest, "invalid nim")
		return
	}
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, user, "user found")
}

func (h *UserController) GetByName(c *gin.Context) {
	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	name := c.Query("name")
	users, total, err := h.service.GetByName(c, name, limit, offset)
	if len(name) > 256 {
		helper.Error(c, http.StatusBadRequest, "invalid name")
		return
	}
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, gin.H{"data": users, "total": total}, "users found")
}

func (h *UserController) GetByRole(c *gin.Context) {
	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	role := c.Query("role")
	users, total, err := h.service.GetByRole(c, role, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, gin.H{"data": users, "total": total}, "users found")
}

func (h *UserController) ChangePassword(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	var req model.ChangePasswordCredential
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.ChangePassword(c.Request.Context(), id, req.NewPassword); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "password changed successfully")
}

func (h *UserController) ChangeRole(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	var input model.ChangeRoleCredential
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid input format")
		return
	}

	user, err := h.service.GetById(c, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, "user not found")
		return
	}

	user.Role = input.Role
	allowedRoles := map[string]bool{
		"admin":    true,
		"user":     true,
		"lecturer": true,
	}
	if !allowedRoles[string(user.Role)] {
		helper.Error(c, http.StatusBadRequest, "invalid role")
		return
	}

	userRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := userRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	roleValue := model.Role(roleStr)

	if err := h.service.ChangeRole(c, id, user.Role, roleValue); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to update user role")
		return
	}

	helper.Success(c, user, "user role updated successfully")
}

func (h *UserController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		helper.Error(c, http.StatusUnauthorized, "missing refresh token")
		return
	}

	newAccessToken, err := h.service.RefreshToken(c, refreshToken)
	if err != nil {
		helper.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.Success(c, newAccessToken, "token refreshed")
}

func (h *UserController) BulkInsert(c *gin.Context) {
	prefix := c.Query("prefix")
	startStr := c.Query("start")
	endStr := c.Query("end")

	if len(prefix) != 2 {
		helper.Error(c, http.StatusBadRequest, "prefix must be 2")
		return
	}

	startInt, err := strconv.Atoi(startStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid start")
		return
	}

	endInt, err := strconv.Atoi(endStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid end")
		return
	}

	prefixes := "G1A" + prefix

	users, err := h.service.BulkInsert(c, prefixes, startInt, endInt)
	if err != nil {
		helper.Error(c, 500, err.Error())
		return
	}

	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "NIM")
	f.SetCellValue("Sheet1", "B1", "Password")

	for i := range users {
		row := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", "A"+row, users[i].Nim)
		f.SetCellValue("Sheet1", "B"+row, users[i].Password)
	}

	storageDir := "./storages/files"
	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to create storage directory")
		return
	}

	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("bulk_users_%s.xlsx", timestamp)
	filepath := fmt.Sprintf("%s/%s", storageDir, filename)

	if err := f.SaveAs(filepath); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to save xls file")
		return
	}

	if err := h.xlsPathService.SaveXlsPath(c, filepath); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to save xls path")
		return
	}

	response := map[string]interface{}{
		"users":    users,
		"file":     filename,
		"filepath": filepath,
		"message":  "users created and xls file saved",
	}

	helper.Success(c, response, "users created and xls file saved")
}
