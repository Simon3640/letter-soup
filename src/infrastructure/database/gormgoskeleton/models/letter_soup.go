package dbmodels

import "gorm.io/gorm"

type LetterSoup struct {
	gorm.Model
	Rows            int      `gorm:"not null"`
	Columns         int      `gorm:"not null"`
	UserID          uint     `gorm:"not null;index"`
	Grid            []string `gorm:"type:text[];not null"`
	Words           []string `gorm:"type:text[];not null"`
	FoundedWords    []string `gorm:"type:text[];default:'{}'"`
	NotFoundedWords []string `gorm:"type:text[];default:'{}'"`
}

func (LetterSoup) TableName() string {
	return "letter_soup"
}

var _ DBModel = (*LetterSoup)(nil)
