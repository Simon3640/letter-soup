package services

import (
	contractsProviders "lettersoup/src/application/contracts/providers"
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/domain/models"
)

func CreateOneTimeTokenService(
	userID uint,
	purpose models.OneTimeTokenPurpose,
	hashProvider contractsProviders.IHashProvider,
	tokenRepository contracts_repositories.IOneTimeTokenRepository,
) (string, *application_errors.ApplicationError) {
	token, hash, err := hashProvider.OneTimeToken()
	if err != nil {
		return "", err
	}

	tokenCreate := dtos.NewOneTimeTokenCreate(userID, purpose, hash)
	_, err = tokenRepository.Create(*tokenCreate)
	if err != nil {
		return "", err
	}
	return token, nil
}
