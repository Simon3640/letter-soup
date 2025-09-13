package mocks

import (
	contracts_repositories "gormgoskeleton/src/application/contracts/repositories"
	dtos "gormgoskeleton/src/application/shared/DTOs"
	"gormgoskeleton/src/domain/models"
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
