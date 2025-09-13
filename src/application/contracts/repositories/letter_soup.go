package contracts_repositories

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
)

type ILetterSoupRepository interface {
	IRepositoryBase[
		dtos.LetterSoupCreateSolution,
		dtos.LetterSoupUpdateSolution,
		models.LetterSoup,
		models.LetterSoup,
	]
}
