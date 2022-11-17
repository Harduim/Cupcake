package domain

import "time"

type Match struct {
	ID              string        `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Date            time.Time     `json:"date" validate:"required"`
	NationalTeamAID *string       `json:"nationalTeamA" validate:"uuid" gorm:"type:varchar(255)"`
	NationalTeamBID *string       `json:"nationalTeamB" validate:"uuid" gorm:"type:varchar(255)"`
	NationalTeamA   *NationalTeam `gorm:"foreignKey:NationalTeamAID"`
	NationalTeamB   *NationalTeam `gorm:"foreignKey:NationalTeamBID"`
	GolA            *int8         `json:"golA" gorm:"type:integer"`
	GolB            *int8         `json:"golB" gorm:"type:integer"`
	BracketID       string        `json:"bracketId" validate:"required,uuid" gorm:"type:varchar(255)"`
	Bracket         Bracket       `gorm:"foreignKey:BracketID"`
	WinnerID        *string       `json:"winnerId" validate:"uuid" gorm:"type:varchar(255)"`
	Winner          *NationalTeam `gorm:"foreignKey:WinnerID"`
}

func NewMatch() *Match {
	return &Match{}
}

func (match *Match) Validate() error {
	err := validate.Struct(match)

	if err != nil {
		return err
	}
	return nil
}
