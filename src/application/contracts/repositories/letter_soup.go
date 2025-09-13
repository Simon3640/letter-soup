package contracts_repositories

import (
	dtos "gormgoskeleton/src/application/shared/DTOs"
	"gormgoskeleton/src/domain/models"
)

type ILetterSoupRepository interface {
	IRepositoryBase[
		dtos.LetterSoupCreateSolution,
		dtos.LetterSoupUpdateSolution,
		models.LetterSoup,
		models.LetterSoup,
	]
}
