package services

import (
	"fmt"
	application_errors "lettersoup/src/application/shared/errors"
	"lettersoup/src/application/shared/locales/messages"
	"lettersoup/src/application/shared/status"
	"lettersoup/src/domain/models"
	domain_utils "lettersoup/src/domain/utils"
	"strings"
)

func SolveLetterSoupService(ls models.LetterSoupBase) (*models.LetterSoupSolution, *application_errors.ApplicationError) {
	// Validate the letter soup
	if errs := ls.Validate(); len(errs) > 0 {
		return nil, &application_errors.ApplicationError{
			Code:    status.InternalError,
			Context: messages.MessageKeysInstance.SOMETHING_WENT_WRONG,
			ErrMsg:  fmt.Sprintf("Invalid letter soup: %v", errs),
		}
	}
	matrix := make([][]rune, ls.Rows)
	for i := 0; i < ls.Rows; i++ {
		matrix[i] = []rune(strings.ToUpper(string(ls.Grid[i])))
	}

	letterSoupSolution := models.LetterSoupSolution{
		LetterSoupBase: ls,
		FoundWords:     []models.FoundWord{},
	}

	transforms := domain_utils.MakeTransform("AllTransforms", matrix, false)
	// Add diagonal transforms
	diagSE := domain_utils.BuildDiagonalMatrixSE(matrix)
	transforms = append(transforms, domain_utils.MakeTransform("DiagonalSE", diagSE, true)...)

	diagSW := domain_utils.BuildDiagonalMatrixSW(matrix)
	transforms = append(transforms, domain_utils.MakeTransform("DiagonalSW", diagSW, true)...)

	var foundWords []models.FoundWord

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
						foundWords = append(foundWords, models.FoundWord{
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
