package setups

import (
	dtos "gormgoskeleton/src/application/shared/DTOs"
	"gormgoskeleton/src/domain/models"
	dbModels "gormgoskeleton/src/infrastructure/database/gormgoskeleton/models"
	"gormgoskeleton/src/infrastructure/repositories"
)

type SetupLetterSoup struct {
	SetupBase[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup]
}

var _ SetupModel[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup] = (*SetupLetterSoup)(nil)

func NewSetupLetterSoup() *SetupLetterSoup {
	return &SetupLetterSoup{
		SetupBase: SetupBase[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup]{
			modelConverter: &repositories.LetterSoupConverter{},
		},
	}
}
