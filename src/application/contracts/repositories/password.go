package contracts_repositories

import (
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

type IPasswordRepository interface {
	IRepositoryBase[dtos.PasswordCreate, dtos.PasswordUpdate, models.Password, models.PasswordInDB]
	GetActivePassword(userEmail string) (*models.Password, *application_errors.ApplicationError)
}
