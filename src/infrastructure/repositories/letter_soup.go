package repositories

import (
	contractsProviders "gormgoskeleton/src/application/contracts/providers"
	contracts_repositories "gormgoskeleton/src/application/contracts/repositories"
	dtos "gormgoskeleton/src/application/shared/DTOs"
	"gormgoskeleton/src/domain/models"
	dbModels "gormgoskeleton/src/infrastructure/database/gormgoskeleton/models"
	dbmodels "gormgoskeleton/src/infrastructure/database/gormgoskeleton/models"

	"gorm.io/gorm"
)

type LetterSoupRepository struct {
	RepositoryBase[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup]
}

var _ contracts_repositories.ILetterSoupRepository = (*LetterSoupRepository)(nil)

type LetterSoupConverter struct{}

var _ ModelConverter[dtos.LetterSoupCreateSolution, dtos.LetterSoupUpdateSolution, models.LetterSoup, dbModels.LetterSoup] = (*LetterSoupConverter)(nil)

func (uc *LetterSoupConverter) ToGormCreate(model dtos.LetterSoupCreateSolution) *dbModels.LetterSoup {
	return &dbModels.LetterSoup{
		Rows:            model.Rows,
		Columns:         model.Columns,
		Grid:            model.Grid,
		Words:           model.Words,
		UserID:          model.UserID,
		FoundedWords:    model.FoundedWords,
		NotFoundedWords: model.NotFoundedWords,
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
			Grid:    ormModel.Grid,
			Words:   ormModel.Words,
			UserID:  ormModel.UserID,
		},
		FoundedWords:    ormModel.FoundedWords,
		NotFoundedWords: ormModel.NotFoundedWords,
	}
}

func (uc *LetterSoupConverter) ToGormUpdate(model dtos.LetterSoupUpdateSolution) *dbmodels.LetterSoup {
	var updated dbmodels.LetterSoup

	if model.Grid != nil {
		updated.Grid = *model.Grid
	}
	if model.Words != nil {
		updated.Words = *model.Words
	}

	if model.FoundedWords != nil {
		updated.FoundedWords = model.FoundedWords
	}
	if model.NotFoundedWords != nil {
		updated.NotFoundedWords = model.NotFoundedWords
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
