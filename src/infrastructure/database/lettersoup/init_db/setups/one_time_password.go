package setups

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
	dbModels "lettersoup/src/infrastructure/database/lettersoup/models"
	"lettersoup/src/infrastructure/repositories"
)

type SetupOneTimePassword struct {
	SetupBase[dtos.OneTimePasswordCreate, dtos.OneTimePasswordUpdate, models.OneTimePassword, dbModels.OneTimePassword]
}

var _ SetupModel[dtos.OneTimePasswordCreate, dtos.OneTimePasswordUpdate, models.OneTimePassword, dbModels.OneTimePassword] = (*SetupOneTimePassword)(nil)

func NewSetupOneTimePassword() *SetupOneTimePassword {
	return &SetupOneTimePassword{
		SetupBase: SetupBase[dtos.OneTimePasswordCreate, dtos.OneTimePasswordUpdate, models.OneTimePassword, dbModels.OneTimePassword]{
			modelConverter: &repositories.OneTimePasswordConverter{},
		},
	}
}
