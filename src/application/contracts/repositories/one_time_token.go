package contracts_repositories

import (
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

type IOneTimeTokenRepository interface {
	IRepositoryBase[dtos.OneTimeTokenCreate, dtos.OneTimeTokenUpdate, models.OneTimeToken, models.OneTimeToken]
	GetByTokenHash(tokenHash []byte) (*models.OneTimeToken, *application_errors.ApplicationError)
}
