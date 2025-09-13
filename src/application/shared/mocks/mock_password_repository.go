package mocks

import (
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

type MockPasswordRepository struct {
	MockRepositoryBase[dtos.PasswordCreate, dtos.PasswordUpdate, models.Password, models.PasswordInDB]
}

var _ contracts_repositories.IPasswordRepository = (*MockPasswordRepository)(nil)

func (m *MockPasswordRepository) GetActivePassword(userEmail string) (*models.Password, *application_errors.ApplicationError) {
	args := m.Called(userEmail)
	errorArg := args.Get(1)
	if errorArg != nil {
		return args.Get(0).(*models.Password), errorArg.(*application_errors.ApplicationError)
	}
	return args.Get(0).(*models.Password), nil
}
