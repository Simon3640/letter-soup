package dtos

import (
	"gormgoskeleton/src/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterSoupValidate(t *testing.T) {
	assert := assert.New(t)

	// Valid case
	validLS := LetterSoupCreate{
		LetterSoup: models.LetterSoup{
			Grid:  []string{"ABC", "DEF", "GHI"},
			Words: []string{"ADG", "BEH", "CFI"},
		},
	}
	errs := validLS.Validate()
	assert.Len(errs, 0, "Expected no validation errors for valid input")

	// Grid Conversion case
	convertLS := LetterSoupCreate{
		LetterSoup: models.LetterSoup{
			Grid:  []string{"A,B,C,D,E,F,G,H,I"},
			Words: []string{"ADG", "BEH", "CFI"},
		},
	}
	errs = convertLS.Validate()
	assert.Len(errs, 0, "Expected no validation errors after grid conversion")
	assert.Equal([]string{"ABC", "DEF", "GHI"}, convertLS.Grid, "Grid should be converted correctly")
}
