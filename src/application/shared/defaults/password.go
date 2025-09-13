package defaults

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
)

var AdminPassword = dtos.PasswordCreate{
	PasswordBase: models.PasswordBase{
		UserID:   1,
		Hash:     "hashed_password_for_admin",
		IsActive: false,
	},
}
var DefaultPasswords = []dtos.PasswordCreate{}
