package models

import (
	"strconv"
	"unicode/utf8"
)

type LetterSoupBase struct {
	Rows    int
	Columns int
	UserID  uint     `json:"user_id"`
	Grid    []string `json:"grid"`
	Words   []string `json:"words"`
}

type LetterSoup struct {
	LetterSoupBase
	DBBaseModel
	FoundedWords    []string `json:"founded_words"`
	NotFoundedWords []string `json:"not_founded_words"`
}

func (ls *LetterSoupBase) Validate() []string {
	var errs []string

	ls.Rows = len(ls.Grid)
	ls.Columns = utf8.RuneCountInString(ls.Grid[0])

	if ls.Rows <= 0 {
		errs = append(errs, "rows must be greater than 0")
	}
	if ls.Columns <= 0 {
		errs = append(errs, "columns must be greater than 0")
	}
	if len(ls.Grid) != ls.Rows {
		errs = append(errs, "grid row count does not match rows")
	}
	for i, row := range ls.Grid {
		if len(row) != ls.Columns {
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
