package mocks

import (
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
)

type MockLetterSoupRepository struct {
	MockRepositoryBase[
		dtos.LetterSoupCreateSolution,
		dtos.LetterSoupUpdateSolution,
		models.LetterSoup,
		models.LetterSoup,
	]
}

var _ contracts_repositories.ILetterSoupRepository = (*MockLetterSoupRepository)(nil)
