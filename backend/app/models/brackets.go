package domain

import (
	"fmt"
	"time"
)

type DateError struct {
	OpenDate  string
	CloseDate string
}

func (e *DateError) Error() string {
	return fmt.Sprintf("Open date %s is higher then close date %s", e.OpenDate, e.CloseDate)
}

type Bracket struct {
	ID         string    `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name       string    `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Multiplier int8      `json:"multiplier" validate:"required"`
	OpenDate   time.Time `json:"openDate" validate:"required"`
	CloseDate  time.Time `json:"closeDate" validate:"required"`
}

func NewBracket() *Bracket {
	return &Bracket{}
}

func (bracket *Bracket) Validate() error {
	err := validate.Struct(bracket)

	if bracket.OpenDate.After(bracket.CloseDate) {
		return &DateError{
			OpenDate:  bracket.OpenDate.String(),
			CloseDate: bracket.CloseDate.String(),
		}
	}

	if err != nil {
		return err
	}

	return nil
}
