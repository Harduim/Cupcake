package models

import (
	uuid "github.com/satori/go.uuid"
)

type Joker struct {
	ID              string `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	GolA            *int   `json:"golA" validate:"required" gorm:"type:integer"`
	GolB            *int   `json:"golB" validate:"required" gorm:"type:integer"`
	UserID          string `json:"userId" validate:"required,uuid" gorm:"type:varchar(255);primary_key;uniqueIndex;"`
	NationalTeamAID string `json:"nationalTeamAId" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamBID string `json:"nationalTeamBId" validate:"required,uuid" gorm:"type:varchar(255)"`
	WinnerID        string `json:"winnerId" validate:"required,uuid" gorm:"type:varchar(255)"`
}

func (joker *Joker) Prepare() {
	joker.ID = uuid.NewV4().String()
}

func (joker *Joker) Validate() error {
	err := validate.Struct(joker)

	if err != nil {
		return err
	}
	return nil
}
