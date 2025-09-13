package usecases_letter_soup

import (
	"context"

	contractsProviders "lettersoup/src/application/contracts/providers"
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/application/shared/guards"
	"lettersoup/src/application/shared/locales"
	"lettersoup/src/application/shared/locales/messages"
	"lettersoup/src/application/shared/services"
	"lettersoup/src/application/shared/status"
	usecase "lettersoup/src/application/shared/use_case"
	"lettersoup/src/domain/models"
)

type CreateLetterSoupUseCase struct {
	usecase.BaseUseCaseValidation[dtos.LetterSoupCreate, models.LetterSoup]
	log  contractsProviders.ILoggerProvider
	repo contracts_repositories.ILetterSoupRepository
}

var _ usecase.BaseUseCase[dtos.LetterSoupCreate, models.LetterSoup] = (*CreateLetterSoupUseCase)(nil)

func (uc *CreateLetterSoupUseCase) SetLocale(locale locales.LocaleTypeEnum) {
	if locale != "" {
		uc.Locale = locale
	}
}

func (uc *CreateLetterSoupUseCase) Execute(ctx context.Context,
	locale locales.LocaleTypeEnum,
	input dtos.LetterSoupCreate,
) *usecase.UseCaseResult[models.LetterSoup] {
	result := usecase.NewUseCaseResult[models.LetterSoup]()
	uc.SetLocale(locale)
	uc.Validate(ctx, input, result)
	if result.HasError() {
		return result
	}

	input.Validate()

	var entity dtos.LetterSoupCreateSolution
	entity.LetterSoupCreate = input

	letterSoupSolution, err := services.SolveLetterSoupService(input.LetterSoupBase)
	if err != nil {
		uc.log.Error("CreateLetterSoupUseCase: Execute: Error solving letter soup", err.ToError())
		result.SetError(
			err.Code,
			uc.AppMessages.Get(
				uc.Locale,
				err.Context,
			),
		)
		return result
	}

	entity.FoundedWords = letterSoupSolution.FoundedWords()
	entity.NotFoundedWords = letterSoupSolution.NotFoundedWords()

	res, err := uc.repo.Create(entity)
	if err != nil {
		uc.log.Error("CreateLetterSoupUseCase: Execute: Error creating letter soup", err.ToError())
		result.SetError(
			err.Code,
			uc.AppMessages.Get(
				uc.Locale,
				err.Context,
			),
		)
		return result
	}

	result.SetData(
		status.Success,
		*res,
		uc.AppMessages.Get(
			uc.Locale,
			messages.MessageKeysInstance.PASSWORD_CREATED,
		),
	)
	return result
}

func NewCreateLetterSoupUseCase(
	log contractsProviders.ILoggerProvider,
	repo contracts_repositories.ILetterSoupRepository,
) *CreateLetterSoupUseCase {
	return &CreateLetterSoupUseCase{
		BaseUseCaseValidation: usecase.BaseUseCaseValidation[dtos.LetterSoupCreate, models.LetterSoup]{
			AppMessages: locales.NewLocale(locales.EN_US),
			Guards: usecase.NewGuards(
				guards.RoleGuard("admin", "user"),
			),
		},
		log:  log,
		repo: repo,
	}
}
