package dbmodels

import (
	"gorm.io/gorm"
)

type LetterSoup struct {
	gorm.Model
	Rows            int    `gorm:"not null"`
	Columns         int    `gorm:"not null"`
	UserID          uint   `gorm:"not null;index"`
	Grid            string `gorm:"type:varchar(255);not null"`
	Words           string `gorm:"type:varchar(255);not null"`
	FoundedWords    string `gorm:"type:varchar(255);default:'{}'"`
	NotFoundedWords string `gorm:"type:varchar(255);default:'{}'"`
}

func (LetterSoup) TableName() string {
	return "letter_soup"
}

var _ DBModel = (*LetterSoup)(nil)
