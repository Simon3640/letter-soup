package defaults

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
)

var AdminUser = dtos.UserCreate{
	UserBase: models.UserBase{
		Name:   "Admin",
		Email:  "admin@lettersoup.com",
		Phone:  "1234567890",
		Status: "active",
		RoleID: 1,
	},
}

var DefaultUsers = []dtos.UserCreate{
	AdminUser,
}
