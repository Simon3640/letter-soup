package services

import (
	"testing"

	"lettersoup/src/domain/models"

	"github.com/stretchr/testify/assert"
)

func TestSolveLetterSoup_BasicDirections(t *testing.T) {
	assert := assert.New(t)
	ls := models.LetterSoupBase{
		Grid: []string{
			"ABCD",
			"EFGH",
			"IJKL",
			"MNOP",
		},
		Words: []string{"ABCD", "DCBA", "AEIM", "MIEA", "MJGD"},
	}

	sol, err := SolveLetterSoupService(ls)
	assert.Nil(err)

	found := map[string]bool{}
	for _, fw := range sol.FoundWords {
		found[fw.Word] = true
	}

	tests := []string{"ABCD", "DCBA", "AEIM", "MIEA"}
	for _, w := range tests {
		assert.True(found[w], "Expected to find word "+w+" but not found")
	}
}

func TestSolveLetterSoup_Diagonals(t *testing.T) {
	assert := assert.New(t)
	ls := models.LetterSoupBase{
		Grid: []string{
			"AXXX",
			"XBXZ",
			"XXCX",
			"XXXD",
		},
		Words: []string{"ABCD", "DCBA"},
	}

	sol, err := SolveLetterSoupService(ls)
	assert.Nil(err)

	found := map[string]bool{}
	for _, fw := range sol.FoundWords {
		found[fw.Word] = true
	}

	assert.True(found["ABCD"], "Expected to find diagonal word ABCD ↘ but not found")
	assert.True(found["DCBA"], "Expected to find diagonal word DCBA ↖ but not found")
}

func TestSolveLetterSoup_DiagonalReverse(t *testing.T) {
	assert := assert.New(t)
	ls := models.LetterSoupBase{
		Grid: []string{
			"XXXA",
			"XXBX",
			"XCXX",
			"DXXX",
		},
		Words: []string{"ABCD", "DCBA"},
	}

	sol, err := SolveLetterSoupService(ls)
	assert.Nil(err)

	found := map[string]bool{}
	for _, fw := range sol.FoundWords {
		found[fw.Word] = true
	}

	assert.Equal(2, len(sol.FoundWords), "Expected to find 2 words but found %d", len(sol.FoundWords))
	assert.True(found["ABCD"], "Expected to find diagonal word ABCD ↙ but not found")
	assert.True(found["DCBA"], "Expected to find diagonal word DCBA ↗ but not found")
}

func TestSolveLetterSoup_MultipleOccurrences(t *testing.T) {
	assert := assert.New(t)
	ls := models.LetterSoupBase{
		Grid: []string{
			"TEST",
			"ESTE",
			"STES",
			"TEST",
		},
		Words: []string{"TEST"},
	}

	sol, err := SolveLetterSoupService(ls)
	assert.Nil(err)

	assert.Equal(5, len(sol.FoundWords), "Expected multiple occurrences of word TEST, got %d", len(sol.FoundWords))
}

func TestSolveLetterSoup_NoWordsFound(t *testing.T) {
	assert := assert.New(t)
	ls := models.LetterSoupBase{
		Grid: []string{
			"ABCD",
			"EFGH",
			"IJKL",
			"MNOP",
		},
		Words: []string{"XYZ", "123"},
	}

	sol, err := SolveLetterSoupService(ls)
	assert.Nil(err)
	assert.Equal(0, len(sol.FoundWords), "Expected to find 0 words but found %d", len(sol.FoundWords))
	assert.Equal([]string{"XYZ", "123"}, sol.NotFoundedWords(), "Expected not found words to be [XYZ, 123]")
	assert.Equal([]string{}, sol.FoundedWords(), "Expected found words to be []")
}
