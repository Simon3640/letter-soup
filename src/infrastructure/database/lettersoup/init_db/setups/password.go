package setups

import (
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
	dbModels "lettersoup/src/infrastructure/database/lettersoup/models"
	"lettersoup/src/infrastructure/repositories"
)

type SetupPassword struct {
	SetupBase[dtos.PasswordCreate, dtos.PasswordUpdate, models.Password, dbModels.Password]
}

var _ SetupModel[dtos.PasswordCreate, dtos.PasswordUpdate, models.Password, dbModels.Password] = (*SetupPassword)(nil)

func NewSetupPassword() *SetupPassword {
	return &SetupPassword{
		SetupBase: SetupBase[dtos.PasswordCreate, dtos.PasswordUpdate, models.Password, dbModels.Password]{
			modelConverter: &repositories.PasswordConverter{},
		},
	}
}
