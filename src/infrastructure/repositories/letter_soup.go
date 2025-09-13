package repositories

import (
	contractsProviders "gormgoskeleton/src/application/contracts/providers"
	contracts_repositories "gormgoskeleton/src/application/contracts/repositories"
	dtos "gormgoskeleton/src/application/shared/DTOs"
	"gormgoskeleton/src/domain/models"
	dbModels "gormgoskeleton/src/infrastructure/database/gormgoskeleton/models"
	dbmodels "gormgoskeleton/src/infrastructure/database/gormgoskeleton/models"
	"strings"

	"gorm.io/gorm"
)

type LetterSoupRepository struct {
	RepositoryBase[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup]
}

var _ contracts_repositories.ILetterSoupRepository = (*LetterSoupRepository)(nil)

type LetterSoupConverter struct{}

var _ ModelConverter[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup] = (*LetterSoupConverter)(nil)

const separator = "\u241F"

// Convierte []string a string para almacenar
func StringsToDBString(list []string) string {
	return strings.Join(list, separator)
}

// Convierte string de la DB de vuelta a []string
func DBStringToStrings(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, separator)
}

func (uc *LetterSoupConverter) ToGormCreate(model dtos.LetterSoupCreateSolution) *dbModels.LetterSoup {
	return &dbModels.LetterSoup{
		Rows:            model.Rows,
		Columns:         model.Columns,
		Grid:            StringsToDBString(model.Grid),
		Words:           StringsToDBString(model.Words),
		UserID:          model.UserID,
		FoundedWords:    StringsToDBString(model.FoundedWords),
		NotFoundedWords: StringsToDBString(model.NotFoundedWords),
	}
}

func (uc *LetterSoupConverter) ToDomain(ormModel *dbModels.LetterSoup) *models.LetterSoup {

	return &models.LetterSoup{
		DBBaseModel: models.DBBaseModel{
			ID:        ormModel.ID,
			CreatedAt: ormModel.CreatedAt,
			UpdatedAt: ormModel.UpdatedAt,
			DeletedAt: ormModel.DeletedAt.Time,
		},
		LetterSoupBase: models.LetterSoupBase{
			Rows:    ormModel.Rows,
			Columns: ormModel.Columns,
			Grid:    DBStringToStrings(ormModel.Grid),
			Words:   DBStringToStrings(ormModel.Words),
			UserID:  ormModel.UserID,
		},
		FoundedWords:    DBStringToStrings(ormModel.FoundedWords),
		NotFoundedWords: DBStringToStrings(ormModel.NotFoundedWords),
	}
}

func (uc *LetterSoupConverter) ToGormUpdate(model dtos.LetterSoupUpdateSolution) *dbmodels.LetterSoup {
	var updated dbmodels.LetterSoup

	if model.Grid != nil {
		updated.Grid = StringsToDBString(*model.Grid)
	}
	if model.Words != nil {
		updated.Words = StringsToDBString(*model.Words)
	}

	if model.FoundedWords != nil {
		updated.FoundedWords = StringsToDBString(model.FoundedWords)
	}
	if model.NotFoundedWords != nil {
		updated.NotFoundedWords = StringsToDBString(model.NotFoundedWords)
	}
	return &updated
}

func NewLetterSoupRepository(db *gorm.DB, logger contractsProviders.ILoggerProvider) *LetterSoupRepository {
	return &LetterSoupRepository{
		RepositoryBase: RepositoryBase[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup]{
			DB:             db,
			logger:         logger,
			modelConverter: &LetterSoupConverter{},
		},
	}
}
