package update

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/utils/helper"
)

func ValidateAuthorization(
	effectiveRole model.Role,
	oldUser *model.User,
	data model.UpdateUser,
	requesterRole model.Role,
	currentId int,
) error {

	targetRole := oldUser.Role

	if requesterRole == model.RoleUser || requesterRole == model.RoleLecturer {
		if effectiveRole != "" && effectiveRole != targetRole {
			return fmt.Errorf("you are not allowed to change role")
		}
		return nil
	}

	if requesterRole == model.RoleSuperAdmin {
		if targetRole == model.RoleSuperAdmin {
			return fmt.Errorf("super admin cannot modify another super admin")
		}

		if effectiveRole != "" {
			if effectiveRole != model.RoleAdmin &&
				effectiveRole != model.RoleLecturer &&
				effectiveRole != model.RoleUser {
				return fmt.Errorf("invalid role assignment")
			}
		}

		return nil
	}

	if requesterRole == model.RoleAdmin {

		if targetRole == model.RoleAdmin || targetRole == model.RoleSuperAdmin {
			return fmt.Errorf("admin cannot modify %s", targetRole)
		}

		if effectiveRole != "" {
			if effectiveRole == model.RoleAdmin || effectiveRole == model.RoleSuperAdmin {
				return fmt.Errorf("admin cannot assign elevated roles")
			}
		}

		return nil
	}

	return fmt.Errorf("invalid requester role")
}

func NormalizeRoleTransition(oldUser *model.User, data *model.UpdateUser, effectiveRole model.Role) {
	if oldUser.Role == effectiveRole {
		return
	}

	switch {
	case oldUser.Role == model.RoleUser && effectiveRole == model.RoleLecturer:
		emptyString := ""
		data.Nim = &emptyString

	case oldUser.Role == model.RoleLecturer && effectiveRole == model.RoleUser:
		emptyString := ""
		data.Nip = &emptyString

	case oldUser.Role == model.RoleLecturer && (effectiveRole == model.RoleAdmin || effectiveRole == model.RoleSuperAdmin):
		emptyString := ""
		data.Nip = &emptyString

	case oldUser.Role == model.RoleUser && (effectiveRole == model.RoleAdmin || effectiveRole == model.RoleSuperAdmin):
		emptyString := ""
		data.Nim = &emptyString
	}
}

func ValidateRoleRequirements(data model.UpdateUser, effectiveRole model.Role) error {
	switch effectiveRole {
	case model.RoleUser:
		if data.Nip != nil && *data.Nip != "" {
			return fmt.Errorf("user role cannot have nip")
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

func MergeDefaults(oldUser *model.User, data *model.UpdateUser, effectiveRole model.Role) {
	if data.Username == nil {
		data.Username = oldUser.Username
	}

	data.Role = &effectiveRole

	switch effectiveRole {
	case model.RoleUser:
		if data.Nim == nil {
			data.Nim = oldUser.Nim
		}
		if data.AcademicYear == nil {
			data.AcademicYear = &oldUser.AcademicYear
		}
		emptyString := ""
		data.Nip = &emptyString

	case model.RoleLecturer:
		if data.Nip == nil {
			data.Nip = oldUser.Nip
		}
		emptyString := ""
		data.Nim = &emptyString
		emptyAcademicYear := ""
		data.AcademicYear = &emptyAcademicYear

	case model.RoleAdmin, model.RoleSuperAdmin:
		emptyString := ""
		data.Nim = &emptyString
		data.Nip = &emptyString
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

func ValidateRoleTransitionRequirements(oldUser *model.User, data model.UpdateUser, newRole model.Role) error {
	if data.Role != nil {
		if oldUser.Role == *data.Role {
			return nil
		}

		switch newRole {
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
