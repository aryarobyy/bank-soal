package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
	"latih.in-be/utils/response"
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
	ctx := c.Request.Context()

	var user model.RegisterCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// email is unnecessary
	// if !helper.IsValidEmail(user.Email) {
	// 	helper.Error(c, http.StatusBadRequest, "wrong email format")
	// 	return
	// }

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	role := model.Role(roleStr)

	if err := h.service.Register(ctx, user, role); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, user, "user registered")
}

func (h *UserController) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var cred model.LoginCredential
	if err := c.ShouldBindJSON(&cred); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(ctx, cred)
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
	ctx := c.Request.Context()

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.service.GetById(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	userRes := response.UserResponse(*user)

	helper.Success(c, userRes, "user found")
}

func (h *UserController) GetByEmail(c *gin.Context) {
	ctx := c.Request.Context()

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

	user, err := h.service.GetByEmail(ctx, email)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	userRes := response.UserResponse(*user)

	helper.Success(c, userRes, "user found")
}

func (h *UserController) Update(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	currentRole := model.Role(roleStr)

	name := c.PostForm("name")
	email := c.PostForm("email")
	nim := c.PostForm("nim")
	nip := c.PostForm("nip")
	nidn := c.PostForm("nidn")
	username := c.PostForm("username")
	roleForm := c.PostForm("role")
	major := c.PostForm("major")
	faculty := c.PostForm("faculty")
	academicYear := c.PostForm("academic_year")
	statusForm := c.PostForm("status")
	imgDelete := c.PostForm("img_delete")

	if email != "" && !helper.IsValidEmail(email) {
		helper.Error(c, http.StatusBadRequest, "wrong email format")
		return
	}

	updateData := model.UpdateUser{
		Name: helper.BindAndConvertToPtr(name),

		Email:    helper.BindAndConvertToPtr(email),
		Username: helper.BindAndConvertToPtr(username),
		Nim:      helper.BindAndConvertToPtr(nim),
		Nip:      helper.BindAndConvertToPtr(nip),
		Nidn:     helper.BindAndConvertToPtr(nidn),

		Major:        helper.BindAndConvertToPtr(major),
		Faculty:      helper.BindAndConvertToPtr(faculty),
		AcademicYear: helper.BindAndConvertToPtr(academicYear),

		Role:   (*model.Role)(helper.BindAndConvertToPtr(roleForm)),
		Status: (*model.Status)(helper.BindAndConvertToPtr(statusForm)),

		ImgDelete: helper.BindAndConvertToBoolPtr(imgDelete),
	}

	updatedUser, err := h.service.Update(ctx, c, updateData, id, currentRole)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	userRes := response.UserResponse(*updatedUser)

	helper.Success(c, userRes, "user updated successfully")
}

func (h *UserController) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	role := model.Role(roleStr)

	err = h.service.Delete(ctx, id, role)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}
	helper.Success(c, nil, "user deleted")
}

func (h *UserController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}
	users, total, err := h.service.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	usersRes := response.UsersResponse(users)

	helper.Success(c, gin.H{"data": usersRes, "total": total}, "users found")
}

func (h *UserController) GetByNim(c *gin.Context) {
	ctx := c.Request.Context()

	nim := c.Query("nim")
	if len(nim) >= 10 {
		helper.Error(c, http.StatusBadRequest, "invalid nim")
		return
	}

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	role := model.Role(roleStr)

	user, err := h.service.GetByNim(ctx, nim, role)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	userRes := response.UserResponse(*user)

	helper.Success(c, userRes, "user found")
}

func (h *UserController) GetByNidn(c *gin.Context) {
	ctx := c.Request.Context()

	nidn := c.Query("nidn")
	if len(nidn) >= 11 {
		helper.Error(c, http.StatusBadRequest, "invalid nidn")
		return
	}

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	role := model.Role(roleStr)

	user, err := h.service.GetByNidn(ctx, nidn, role)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	userRes := response.UserResponse(*user)

	helper.Success(c, userRes, "user found")
}

func (h *UserController) GetByUsn(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Query("username")
	if len(username) >= 10 {
		helper.Error(c, http.StatusBadRequest, "invalid username")
		return
	}

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	role := model.Role(roleStr)

	user, err := h.service.GetByUsn(ctx, username, role)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	userRes := response.UserResponse(*user)

	helper.Success(c, userRes, "user found")
}

func (h *UserController) GetByName(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	name := c.Query("name")
	users, total, err := h.service.GetByName(ctx, name, limit, offset)
	if len(name) > 256 {
		helper.Error(c, http.StatusBadRequest, "invalid name")
		return
	}
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	usersRes := response.UsersResponse(users)

	helper.Success(c, gin.H{"data": usersRes, "total": total}, "users found")
}

func (h *UserController) GetByRole(c *gin.Context) {
	ctx := c.Request.Context()

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	role := c.Query("role")

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}
	log.Printf("role: %s", currRole)

	userRole, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	users, total, err := h.service.GetByRole(ctx, role, limit, offset, userRole)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	usersRes := response.UsersResponse(users)

	helper.Success(c, gin.H{"data": usersRes, "total": total}, "users found")
}

func (h *UserController) ChangePassword(c *gin.Context) {
	ctx := c.Request.Context()

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

	currRole, exists := c.Get("role")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "role not found in context")
		return
	}

	roleStr, ok := currRole.(string)
	if !ok {
		helper.Error(c, http.StatusBadRequest, "invalid role type")
		return
	}

	role := model.Role(roleStr)

	if err := h.service.ChangePassword(ctx, id, req.NewPassword, role); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "password changed successfully")
}

func (h *UserController) ChangeRole(c *gin.Context) {
	ctx := c.Request.Context()

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

	user, err := h.service.GetById(ctx, id)
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

	if err := h.service.ChangeRole(ctx, id, user.Role, roleValue); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to update user role")
		return
	}

	userRes := response.UserResponse(*user)

	helper.Success(c, userRes, "user role updated successfully")
}

func (h *UserController) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		helper.Error(c, http.StatusUnauthorized, "missing refresh token")
		return
	}

	newAccessToken, err := h.service.RefreshToken(ctx, refreshToken)
	if err != nil {
		helper.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.Success(c, newAccessToken, "token refreshed")
}

func (h *UserController) BulkInsert(c *gin.Context) {
	ctx := c.Request.Context()

	var batchUser model.BulkUserCredential

	if err := c.ShouldBindJSON(&batchUser); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid input body")
		return
	}

	year := strings.TrimSpace(batchUser.AcademicYear)

	if len(year) < 2 {
		helper.Error(c, http.StatusBadRequest, "academic year must be at least 2 characters")
		return
	}

	if len(year) != 4 {
		helper.Error(c, http.StatusBadRequest, "academic year must be 4 digits, e.g. 2025")
		return
	}

	prefix := year[len(year)-2:]

	startStr := c.Query("start")
	endStr := c.Query("end")

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

	users, err := h.service.BulkInsert(ctx, batchUser, prefix, startInt, endInt)
	if err != nil {
		helper.Error(c, 500, err.Error())
		return
	}

	storageDir := "./storages/files"
	filename, filepath, err := h.xlsPathService.ExportUsersToExcel(users, storageDir)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.xlsPathService.SaveXlsPath(c, filepath); err != nil {
		helper.Error(c, http.StatusInternalServerError, "failed to save xls path")
		return
	}

	response := map[string]interface{}{
		"file":     filename,
		"filepath": filepath,
		"message":  "users created and xls file saved",
		"users":    users,
	}

	helper.Success(c, response, "users created and xls file saved")
}
