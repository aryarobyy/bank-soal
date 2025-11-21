package update

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/utils/helper"
)

func ValidateAuthorization(oldUser *model.User, data model.UpdateUser, requesterRole model.Role) error {
	if requesterRole != model.RoleAdmin && requesterRole != model.RoleSuperAdmin {
		if data.Role != nil && *data.Role != oldUser.Role {
			return fmt.Errorf("you are not allowed to change role")
		}
	}

	if oldUser.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return fmt.Errorf("user not found")
	}

	if data.Role != nil && *data.Role == model.RoleSuperAdmin && requesterRole != model.RoleSuperAdmin {
		return fmt.Errorf("you cant access this role")
	}

	return nil
}

func NormalizeRoleTransition(oldUser *model.User, data *model.UpdateUser, newRole model.Role) {
	if oldUser.Role == newRole {
		return
	}

	switch {
	case oldUser.Role == model.RoleUser && newRole == model.RoleLecturer:
		emptyString := ""
		data.Nim = &emptyString

	case oldUser.Role == model.RoleLecturer && newRole == model.RoleUser:
		emptyString := ""
		data.Nip = &emptyString
		data.Nidn = &emptyString

	case oldUser.Role == model.RoleLecturer && (newRole == model.RoleAdmin || newRole == model.RoleSuperAdmin):
		emptyString := ""
		data.Nip = &emptyString
		data.Nidn = &emptyString

	case oldUser.Role == model.RoleUser && (newRole == model.RoleAdmin || newRole == model.RoleSuperAdmin):
		emptyString := ""
		data.Nim = &emptyString
	}
}

func ValidateRoleRequirements(data model.UpdateUser, role model.Role) error {
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
		if data.AcademicYear != nil && *data.AcademicYear != "" {
			return fmt.Errorf("lecturer role cannot have academic_year")
		}

	case model.RoleAdmin, model.RoleSuperAdmin: //Kalo ada exception
		// if data.Nim != nil && *data.Nim != "" {
		// 	return fmt.Errorf("admin role cannot have nim")
		// }
		// if data.Nip != nil && *data.Nip != "" {
		// 	return fmt.Errorf("admin role cannot have nip")
		// }
		// if data.Nidn != nil && *data.Nidn != "" {
		// 	return fmt.Errorf("admin role cannot have nidn")
		// }
		// if data.AcademicYear != nil && *data.AcademicYear != "" {
		// 	return fmt.Errorf("admin role cannot have academic_year")
		// }
	}

	return nil
}

func MergeDefaults(oldUser *model.User, data *model.UpdateUser, role model.Role) {
	if data.Username == nil {
		data.Username = oldUser.Username
	}

	data.Role = &role

	switch role {
	case model.RoleUser:
		if data.Nim == nil {
			data.Nim = oldUser.Nim
		}
		if data.AcademicYear == nil {
			data.AcademicYear = &oldUser.AcademicYear
		}
		emptyString := ""
		data.Nip = &emptyString
		data.Nidn = &emptyString

	case model.RoleLecturer:
		if data.Nip == nil {
			data.Nip = oldUser.Nip
		}
		if data.Nidn == nil {
			data.Nidn = oldUser.Nidn
		}
		emptyString := ""
		data.Nim = &emptyString
		emptyAcademicYear := ""
		data.AcademicYear = &emptyAcademicYear

	case model.RoleAdmin, model.RoleSuperAdmin:
		emptyString := ""
		data.Nim = &emptyString
		data.Nip = &emptyString
		data.Nidn = &emptyString
		emptyAcademicYear := ""
		data.AcademicYear = &emptyAcademicYear
	}
}

func HandleUserImageUpload(c *gin.Context, oldUser *model.User, data *model.UpdateUser, id int) error {
	if data.ImgDelete != nil && *data.ImgDelete {
		if oldUser.ImgUrl != "" {
			if err := helper.DeleteImage(oldUser.ImgUrl); err != nil {
				return fmt.Errorf("failed to delete image: %w", err)
			}
		}
		emptyUrl := ""
		data.ImgUrl = &emptyUrl
		return nil
	}

	file, _ := c.FormFile("image")
	if file == nil {
		data.ImgUrl = &oldUser.ImgUrl
		return nil
	}

	if oldUser.ImgUrl != "" {
		if err := helper.DeleteImage(oldUser.ImgUrl); err != nil {
			return fmt.Errorf("failed to delete old image: %w", err)
		}
	}

	imgDir := "./storages/images/user"
	newImageUrl, err := helper.UploadImage(c, id, imgDir)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	data.ImgUrl = &newImageUrl
	return nil
}

func FormatUpdateUserError(err error, data model.UpdateUser) error {
	if strings.Contains(err.Error(), "Unknown column") {
		var fieldName string
		parts := strings.Split(err.Error(), "'")
		if len(parts) >= 2 {
			fieldName = parts[1]
		}
		val := helper.GetFieldValue(data, fieldName)
		return fmt.Errorf("field '%s' with value '%v' is undefined", fieldName, val)
	}

	return fmt.Errorf("update gagal: %v", err)
}

func ValidateRoleTransitionRequirements(oldRole model.Role, data model.ChangeRoleCredential) error {
	if data.Role != "" {
		if oldRole == data.Role {
			return nil
		}

		switch data.Role {
		case model.RoleUser:
			if data.Nim == nil || *data.Nim == "" {
				return fmt.Errorf("nim is require for user")
			}
			if data.AcademicYear == nil || *data.AcademicYear == "" {
				return fmt.Errorf("academic year is require for user")
			}

		case model.RoleLecturer:
			if data.Nip == nil || *data.Nip == "" {
				return fmt.Errorf("nip is require for lecturer")
			}
			if data.Nidn == nil || *data.Nidn == "" {
				return fmt.Errorf("nidn is require for lecturer")
			}

		case model.RoleAdmin:
			if data.Username == nil || *data.Username == "" {
				return fmt.Errorf("username is require for admin")
			}

		case model.RoleSuperAdmin:
			return fmt.Errorf("you cant access this")
		}
	}
	return nil
}
