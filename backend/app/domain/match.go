package domain

import "time"

type Match struct {
	ID              string        `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Date            time.Time     `json:"date"`
	NationalTeamAID string        `json:"national_team_a" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamBID string        `json:"national_team_b" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamA   *NationalTeam `gorm:"foreignKey:NationalTeamAID"`
	NationalTeamB   *NationalTeam `gorm:"foreignKey:NationalTeamBID"`
	GolA            int8          `json:"gol_a" validate:"required" gorm:"type:integer"`
	GolB            int8          `json:"gol_b" validate:"required" gorm:"type:integer"`
	BracketID       string        `json:"bracket_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Bracket         *Bracket      `gorm:"foreignKey:BracketID"`
	WinnerID        string        `json:"winner_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Winner          *User         `gorm:"foreignKey:WinnerID"`
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
