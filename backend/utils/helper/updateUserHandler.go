package helper

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
)

func ValidateAuthorization(oldUser *model.User, data model.User, requesterRole model.Role) error {
	if requesterRole != model.RoleAdmin && requesterRole != model.RoleSuperAdmin {
		if data.Role != "" && data.Role != oldUser.Role {
			return fmt.Errorf("you are not allowed to change role")
		}
	}

	if oldUser.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return fmt.Errorf("user not found")
	}

	if data.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return fmt.Errorf("you cant access this role")
	}

	return nil
}

func NormalizeRoleTransition(oldUser *model.User, data *model.User, newRole model.Role) {
	if oldUser.Role == newRole {
		return
	}

	switch oldUser.Role {
	case model.RoleUser:
		data.Nim = nil
		data.AcademicYear = ""

	case model.RoleLecturer:
		data.Nip = nil
		data.Nidn = nil

	case model.RoleAdmin, model.RoleSuperAdmin:
	}
}

func ValidateRoleRequirements(data model.User, role model.Role) error {
	switch role {
	case model.RoleUser:
		if data.Nip != nil && *data.Nip != "" {
			return fmt.Errorf("user role cannot have nip")
		}
		if data.Nidn != nil && *data.Nidn != "" {
			return fmt.Errorf("user role cannot have nidn")
		}

	case model.RoleLecturer:
		if data.Nim != nil && *data.Nim != "" {
			return fmt.Errorf("lecturer role cannot have nim")
		}
		if data.AcademicYear != "" {
			return fmt.Errorf("lecturer role cannot have academic_year")
		}

	case model.RoleAdmin, model.RoleSuperAdmin:
		if data.Nim != nil && *data.Nim != "" {
			return fmt.Errorf("admin role cannot have nim")
		}
		if data.Nip != nil && *data.Nip != "" {
			return fmt.Errorf("admin role cannot have nip")
		}
		if data.Nidn != nil && *data.Nidn != "" {
			return fmt.Errorf("admin role cannot have nidn")
		}
		if data.AcademicYear != "" {
			return fmt.Errorf("admin role cannot have academic_year")
		}
	}

	return nil
}

func MergeDefaults(oldUser *model.User, data *model.User, role model.Role) {
	if data.Username == nil {
		data.Username = oldUser.Username
	}

	data.Role = role

	switch role {
	case model.RoleUser:
		if data.Nim == nil {
			data.Nim = oldUser.Nim
		}
		if data.AcademicYear == "" {
			data.AcademicYear = oldUser.AcademicYear
		}
		data.Nip = nil
		data.Nidn = nil

	case model.RoleLecturer:
		if data.Nip == nil {
			data.Nip = oldUser.Nip
		}
		if data.Nidn == nil {
			data.Nidn = oldUser.Nidn
		}
		data.Nim = nil
		data.AcademicYear = ""

	case model.RoleAdmin, model.RoleSuperAdmin:
		data.Nim = nil
		data.Nip = nil
		data.Nidn = nil
		data.AcademicYear = ""
	}
}

func HandleImageUpload(c *gin.Context, oldUser *model.User, data *model.User, id int) error {
	file, _ := c.FormFile("image")
	if file == nil {
		data.ImgUrl = oldUser.ImgUrl
		return nil
	}

	if oldUser.ImgUrl != "" {
		if err := DeleteImage(oldUser.ImgUrl); err != nil {
			return fmt.Errorf("failed to delete old image: %w", err)
		}
	}

	imgDir := "./storages/images/user"
	newImageUrl, err := UploadImage(c, id, imgDir)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	data.ImgUrl = newImageUrl
	return nil
}

func FormatUpdateError(err error, data model.User) error {
	if strings.Contains(err.Error(), "Unknown column") {
		var fieldName string
		parts := strings.Split(err.Error(), "'")
		if len(parts) >= 2 {
			fieldName = parts[1]
		}
		val := GetFieldValue(data, fieldName)
		return fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
	}

	return fmt.Errorf("update gagal: %v", err)
}
