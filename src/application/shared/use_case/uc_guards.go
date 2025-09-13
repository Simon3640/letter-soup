package usecase

import (
	"lettersoup/src/application/shared/locales/messages"
	"lettersoup/src/domain/models"
)

type Guard func(user models.UserWithRole, input any) *messages.MessageKeysEnum

type Guards struct {
	list  []Guard
	actor models.UserWithRole
}

func (g Guards) Validate(input any) *messages.MessageKeysEnum {
	for _, guard := range g.list {
		if err := guard(g.actor, input); err != nil {
			return err
		}
	}
	return nil
}

func NewGuards(guards ...Guard) Guards {
	return Guards{
		list: guards,
	}
}

func (g *Guards) SetActor(actor models.UserWithRole) {
	g.actor = actor
}
