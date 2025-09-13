package dtomocks

import (
	"encoding/hex"
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
)

var OneTimeTokenBase = models.OneTimeTokenBase{
	UserID:  1,
	Purpose: models.OneTimeTokenPurposeEmailVerify,
	Hash:    []byte(hex.EncodeToString([]byte("hashed_ott"))),
	IsUsed:  false,
}

var OneTimeTokenCreate = dtos.OneTimeTokenCreate{
	OneTimeTokenBase: OneTimeTokenBase,
}
