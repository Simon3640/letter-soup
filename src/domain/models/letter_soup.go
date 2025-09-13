package models

import (
	"fmt"
	domain_utils "gormgoskeleton/src/domain/utils"
	"strconv"
	"strings"
	"unicode/utf8"
)

type LetterSoupBase struct {
	rows    int
	columns int
	Grid    []string `json:"grid"`
	Words   []string `json:"words"`
}

type LetterSoup struct {
	LetterSoupBase
	DBBaseModel
}

func (ls *LetterSoupBase) Validate() []string {
	var errs []string

	ls.rows = len(ls.Grid)
	ls.columns = utf8.RuneCountInString(ls.Grid[0])

	if ls.rows <= 0 {
		errs = append(errs, "rows must be greater than 0")
	}
	if ls.columns <= 0 {
		errs = append(errs, "columns must be greater than 0")
	}
	if len(ls.Grid) != ls.rows {
		errs = append(errs, "grid row count does not match rows")
	}
	for i, row := range ls.Grid {
		if len(row) != ls.columns {
			errs = append(errs, "grid column count does not match columns at row "+strconv.Itoa(i))
		}
	}
	if len(ls.Words) == 0 {
		errs = append(errs, "at least one word is required")
	}

	return errs
}

type FoundWord struct {
	Word      string `json:"word"`
	StartRow  int    `json:"start_row"`
	StartCol  int    `json:"start_col"`
	EndRow    int    `json:"end_row"`
	EndCol    int    `json:"end_col"`
	Direction string `json:"direction"`
}

type LetterSoupSolution struct {
	LetterSoupBase
	FoundWords []FoundWord `json:"found_words"`
}

func (lss LetterSoupSolution) FoundedWords() []string {
	words := make([]string, len(lss.FoundWords))
	for i, fw := range lss.FoundWords {
		words[i] = fw.Word
	}
	return words
}

func (lss LetterSoupSolution) NotFoundedWords() []string {
	foundMap := make(map[string]bool)
	for _, fw := range lss.FoundWords {
		foundMap[fw.Word] = true
	}

	var notFound []string
	for _, w := range lss.Words {
		if !foundMap[w] {
			notFound = append(notFound, w)
		}
	}
	return notFound
}

func SolveLetterSoup(ls LetterSoupBase) (*LetterSoupSolution, error) {
	// Validate the letter soup
	if errs := ls.Validate(); len(errs) > 0 {
		return nil, fmt.Errorf("validation errors: %v", errs)
	}
	matrix := make([][]rune, ls.rows)
	for i := 0; i < ls.rows; i++ {
		matrix[i] = []rune(strings.ToUpper(string(ls.Grid[i])))
	}

	letterSoupSolution := LetterSoupSolution{
		LetterSoupBase: ls,
		FoundWords:     []FoundWord{},
	}

	transforms := domain_utils.MakeTransform("AllTransforms", matrix, false)
	// Add diagonal transforms
	diagSE := domain_utils.BuildDiagonalMatrixSE(matrix)
	transforms = append(transforms, domain_utils.MakeTransform("DiagonalSE", diagSE, true)...)

	diagSW := domain_utils.BuildDiagonalMatrixSW(matrix)
	transforms = append(transforms, domain_utils.MakeTransform("DiagonalSW", diagSW, true)...)

	var foundWords []FoundWord

	for _, tr := range transforms {
		Rt := len(tr.Grid)
		Ct := len(tr.Grid[0])
		for _, word := range ls.Words {
			wordRunes := []rune(strings.ToUpper(word))
			wordLen := len(wordRunes)
			if wordLen > Rt && wordLen > Ct {
				continue
			}
			for r := 0; r < Rt; r++ {
				for c := 0; c <= Ct-wordLen; c++ {
					match := true
					for k := 0; k < wordLen; k++ {
						if tr.Grid[r][c+k] != wordRunes[k] {
							match = false
							break
						}
					}
					if match {
						startPos := tr.PosMap[r][c]
						endPos := tr.PosMap[r][c+wordLen-1]
						foundWords = append(foundWords, FoundWord{
							Word:      word,
							StartRow:  startPos.R,
							StartCol:  startPos.C,
							EndRow:    endPos.R,
							EndCol:    endPos.C,
							Direction: tr.Name,
						})
					}
				}
			}
		}
	}

	letterSoupSolution.FoundWords = foundWords

	return &letterSoupSolution, nil

}
