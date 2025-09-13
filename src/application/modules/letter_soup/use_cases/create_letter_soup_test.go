package usecases_letter_soup

import (
	"context"
	dtos "gormgoskeleton/src/application/shared/DTOs"
	app_context "gormgoskeleton/src/application/shared/context"
	"gormgoskeleton/src/application/shared/locales"
	"gormgoskeleton/src/application/shared/mocks"
	dtomocks "gormgoskeleton/src/application/shared/mocks/dtos"
	"gormgoskeleton/src/domain/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLetterSoupUseCase(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()

	actor := dtomocks.UserWithRole

	letterSoupBase := models.LetterSoupBase{
		Rows:    4,
		Columns: 4,
		UserID:  actor.ID,
		Grid: []string{
			"ABCD",
			"EFGH",
			"IJKL",
			"MNOP",
		},
		Words: []string{"ABCD", "EFGH", "XYZ"},
	}

	testLogger := new(mocks.MockLoggerProvider)
	testLetterSoupRepository := new(mocks.MockLetterSoupRepository)
	testLetterSoup := dtos.LetterSoupCreate{
		LetterSoupBase: letterSoupBase,
	}

	contextWithUser := context.WithValue(ctx, app_context.UserKey, actor)

	testLetterSoupRepository.On(
		"Create",
		mock.AnythingOfType("dtos.LetterSoupCreateSolution"),
	).Return(&models.LetterSoup{
		LetterSoupBase: letterSoupBase,
		DBBaseModel: models.DBBaseModel{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FoundedWords:    []string{"ABCD", "EFGH"},
		NotFoundedWords: []string{"XYZ"},
	}, nil)

	uc := NewCreateLetterSoupUseCase(testLogger, testLetterSoupRepository)

	result := uc.Execute(contextWithUser, locales.ES_ES, testLetterSoup)

	assert.NotNil(result)
	assert.True(result.IsSuccess())
	assert.Equal(uint(1), result.Data.ID)
	assert.Equal(4, result.Data.Rows)
	assert.Equal(4, result.Data.Columns)
	assert.Equal([]string{"ABCD", "EFGH"}, result.Data.FoundedWords)
	assert.Equal([]string{"XYZ"}, result.Data.NotFoundedWords)
}
