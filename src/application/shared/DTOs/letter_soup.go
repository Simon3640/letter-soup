package dtos

import (
	"math"
	"strings"

	"lettersoup/src/domain/models"
)

type LetterSoupCreate struct {
	models.LetterSoupBase
}

type LetterSoupCreateSolution struct {
	LetterSoupCreate
	FoundedWords    []string
	NotFoundedWords []string
}

func (lsc *LetterSoupCreate) Validate() []string {
	errs := []string{}
	// If grid is a list of comma separated letters, convert it to full grid
	if len(lsc.Grid) == 1 && strings.Contains(lsc.Grid[0], ",") {
		letters := strings.Split(lsc.Grid[0], ",")
		square := math.Sqrt(float64(len(letters)))
		if square != math.Trunc(square) {
			errs = append(errs, "grid is not a perfect square")
			return errs
		}
		lsc.Grid = make([]string, int(square))
		for i := 0; i < int(square); i++ {
			lsc.Grid[i] = strings.Join(letters[i*int(square):(i+1)*int(square)], "")
		}
	}

	errs = append(errs, lsc.LetterSoupBase.Validate()...)
	return errs
}

type LetterSoupUpdate struct {
	Grid  *[]string `json:"grid,omitempty"`
	Words *[]string `json:"words,omitempty"`
	ID    uint      `json:"id"`
}

func (lsu LetterSoupUpdate) Validate() []string {
	var errs []string
	if lsu.Grid != nil {
		tempLS := models.LetterSoupBase{
			Grid:  *lsu.Grid,
			Words: []string{},
		}
		errs = append(errs, tempLS.Validate()...)
	}
	if lsu.Words != nil && len(*lsu.Words) == 0 {
		errs = append(errs, "at least one word is required")
	}
	return errs
}

type LetterSoupUpdateSolution struct {
	FoundedWords    []string
	NotFoundedWords []string
	Grid            *[]string
	Words           *[]string
}
