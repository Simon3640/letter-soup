package setups

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
	dbModels "lettersoup/src/infrastructure/database/lettersoup/models"
	"lettersoup/src/infrastructure/repositories"
)

type SetupOneTimeToken struct {
	SetupBase[dtos.OneTimeTokenCreate, dtos.OneTimeTokenUpdate, models.OneTimeToken, dbModels.OneTimeToken]
}

var _ SetupModel[dtos.OneTimeTokenCreate, dtos.OneTimeTokenUpdate, models.OneTimeToken, dbModels.OneTimeToken] = (*SetupOneTimeToken)(nil)

func NewSetupOneTimeToken() *SetupOneTimeToken {
	return &SetupOneTimeToken{
		SetupBase: SetupBase[dtos.OneTimeTokenCreate, dtos.OneTimeTokenUpdate, models.OneTimeToken, dbModels.OneTimeToken]{
			modelConverter: &repositories.OneTimeTokenConverter{},
		},
	}
}
