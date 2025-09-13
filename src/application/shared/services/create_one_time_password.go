package services

import (
	contractsProviders "lettersoup/src/application/contracts/providers"
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

func CreateOneTimePasswordService(
	userID uint,
	purpose models.OneTimePasswordPurpose,
	hashProvider contractsProviders.IHashProvider,
	passwordRepository contracts_repositories.IOneTimePasswordRepository,
) (string, *application_errors.ApplicationError) {
	password, hash, err := hashProvider.GenerateOTP()
	if err != nil {
		return "", err
	}

	passwordCreate := dtos.NewOneTimePasswordCreate(userID, purpose, hash)
	_, err = passwordRepository.Create(*passwordCreate)
	if err != nil {
		return "", err
	}
	return password, nil
}
