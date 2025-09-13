package contracts_repositories

import (
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

type IOneTimePasswordRepository interface {
	IRepositoryBase[dtos.OneTimePasswordCreate, dtos.OneTimePasswordUpdate, models.OneTimePassword, models.OneTimePassword]
	GetByPasswordHash(tokenHash []byte) (*models.OneTimePassword, *application_errors.ApplicationError)
}
