package services

import (
	contractsProviders "lettersoup/src/application/contracts/providers"
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

func CreatePasswordService(
	passwordCreateNoHash dtos.PasswordCreateNoHash,
	hashProvider contractsProviders.IHashProvider,
	passwordRepository contracts_repositories.IPasswordRepository,
) (*models.Password, *application_errors.ApplicationError) {
	hashedPassword, err := hashProvider.HashPassword(passwordCreateNoHash.NoHashedPassword)
	if err != nil {
		return nil, err
	}
	passwordCreate := dtos.NewPasswordCreate(
		passwordCreateNoHash.UserID,
		hashedPassword,
		passwordCreateNoHash.ExpiresAt,
		passwordCreateNoHash.IsActive,
	)
	res, err := passwordRepository.Create(passwordCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}
