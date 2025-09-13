package setups

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
	dbModels "lettersoup/src/infrastructure/database/lettersoup/models"
	"lettersoup/src/infrastructure/repositories"
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
