package mocks

import (
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

type MockOneTimeTokenRepository struct {
	MockRepositoryBase[dtos.OneTimeTokenCreate, dtos.OneTimeTokenUpdate, models.OneTimeToken, models.OneTimeToken]
}

func (m *MockOneTimeTokenRepository) GetByTokenHash(tokenHash []byte) (*models.OneTimeToken, *application_errors.ApplicationError) {
	args := m.Called(tokenHash)
	errorArg := args.Get(1)
	if errorArg != nil {
		return args.Get(0).(*models.OneTimeToken), errorArg.(*application_errors.ApplicationError)
	}
	return args.Get(0).(*models.OneTimeToken), nil
}

var _ contracts_repositories.IOneTimeTokenRepository = (*MockOneTimeTokenRepository)(nil)
