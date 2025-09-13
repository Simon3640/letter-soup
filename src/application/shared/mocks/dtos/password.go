package dtomocks

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
)

var PasswordBase = models.PasswordBase{
	UserID:    1,
	Hash:      "$trongPassword123",
	ExpiresAt: nil,
	IsActive:  true,
}

var PasswordCreate = dtos.PasswordCreate{
	PasswordBase: PasswordBase,
}
