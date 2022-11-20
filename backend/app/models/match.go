package models

import "time"

type Match struct {
	ID              string       `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name            string       `json:"name"`
	Date            time.Time    `json:"date" validate:"required"`
	NationalTeamAID *string      `json:"nationalTeamAId" validate:"uuid" gorm:"type:varchar(255)"`
	NationalTeamBID *string      `json:"nationalTeamBId" validate:"uuid" gorm:"type:varchar(255)"`
	NationalTeamA   NationalTeam `gorm:"foreignKey:NationalTeamAID"`
	NationalTeamB   NationalTeam `gorm:"foreignKey:NationalTeamBID"`
	GolA            *int8        `json:"golA" gorm:"type:integer"`
	GolB            *int8        `json:"golB" gorm:"type:integer"`
	BracketID       string       `json:"bracketId" validate:"required,uuid" gorm:"type:varchar(255)"`
	WinnerID        *string      `json:"winnerId" validate:"uuid" gorm:"type:varchar(255)"`
	Winner          NationalTeam `gorm:"foreignKey:WinnerID"`
}

func (match *Match) Validate() error {
	err := validate.Struct(match)

	if err != nil {
		return err
	}
	return nil
}
