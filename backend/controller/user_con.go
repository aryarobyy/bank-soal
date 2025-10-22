package controller

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/helper"
	"latih.in-be/model"
	"latih.in-be/service"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{service: s}
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

	if !isValidEmail(user.Email) {
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

	if !isValidEmail(cred.Email) {
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

	if !isValidEmail(email) {
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

	if email != "" && !isValidEmail(email) {
		helper.Error(c, http.StatusBadRequest, "wrong email format")
		return
	}

	user := model.User{
		Name:    c.PostForm("name"),
		Email:   email,
		Nim:     c.PostForm("nim"),
		Major:   c.PostForm("major"),
		Faculty: c.PostForm("faculty"),
		Status:  model.Status(c.PostForm("status")),
	}

	academicYearStr := c.PostForm("academic_year")
	if academicYearStr != "" {
		academicYear, err := strconv.Atoi(academicYearStr)
		if err != nil {
			helper.Error(c, http.StatusBadRequest, "invalid academic_year")
			return
		}
		user.AcademicYear = academicYear
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
}

func (h *UserController) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, users, "users found")
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
	name := c.Query("name")
	users, err := h.service.GetByName(c, name)
	if len(name) > 256 {
		helper.Error(c, http.StatusBadRequest, "invalid name")
		return
	}
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, users, "users found")
}

func (h *UserController) GetByRole(c *gin.Context) {
	role := c.Query("role")
	users, err := h.service.GetByRole(c, role)
	if err != nil {
		helper.Error(c, http.StatusNotFound, err.Error())
		return
	}

	helper.Success(c, users, "users found")
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

	if err := h.service.ChangePassword(c.Request.Context(), id, req.OldPassword, req.NewPassword); err != nil {
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

	if input.AdminId == 0 || input.Role == "" {
		helper.Error(c, http.StatusBadRequest, "adminId and role are required")
		return
	}

	admin, err := h.service.GetById(c, input.AdminId)
	if err != nil {
		helper.Error(c, http.StatusUnauthorized, "admin not found")
		return
	}
	if admin.Role != "admin" {
		helper.Error(c, http.StatusForbidden, "you are not authorized to change roles")
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

	if err := h.service.ChangeRole(c, id, user.Role); err != nil {
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

func isValidEmail(e string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(e)
}
