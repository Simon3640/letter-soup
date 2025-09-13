package usecases_user

import (
	"context"

	contractsProviders "lettersoup/src/application/contracts/providers"
	contracts_repositories "lettersoup/src/application/contracts/repositories"
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/application/shared/guards"
	"lettersoup/src/application/shared/locales"
	"lettersoup/src/application/shared/locales/messages"
	"lettersoup/src/application/shared/status"
	usecase "lettersoup/src/application/shared/use_case"
	"lettersoup/src/domain/models"
)

type UpdateUserUseCase struct {
	usecase.BaseUseCaseValidation[dtos.UserUpdate, models.User]
	log  contractsProviders.ILoggerProvider
	repo contracts_repositories.IUserRepository
}

var _ usecase.BaseUseCase[dtos.UserUpdate, models.User] = (*UpdateUserUseCase)(nil)

func (uc *UpdateUserUseCase) SetLocale(locale locales.LocaleTypeEnum) {
	if locale != "" {
		uc.Locale = locale
	}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context,
	locale locales.LocaleTypeEnum,
	input dtos.UserUpdate,
) *usecase.UseCaseResult[models.User] {
	result := usecase.NewUseCaseResult[models.User]()
	uc.SetLocale(locale)
	uc.Validate(ctx, input, result)
	if result.HasError() {
		return result
	}

	res, err := uc.repo.Update(input.ID, input)

	if err != nil {
		uc.log.Error("Error updating user", err.ToError())
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
		uc.AppMessages.Get(uc.Locale, messages.MessageKeysInstance.USER_WAS_CREATED))
	return result
}

func NewUpdateUserUseCase(
	log contractsProviders.ILoggerProvider,
	repo contracts_repositories.IUserRepository,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		BaseUseCaseValidation: usecase.BaseUseCaseValidation[dtos.UserUpdate, models.User]{
			AppMessages: locales.NewLocale(locales.EN_US),
			Guards:      usecase.NewGuards(guards.RoleGuard("admin", "user"), guards.UserResourceGuard[dtos.UserUpdate]()),
		},
		log:  log,
		repo: repo,
	}
}
