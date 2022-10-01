package domain

import "time"

type Match struct {
	ID            string    `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Date          time.Time `json:"date"`
	NationalTeamA string    `json:"national_team_a" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamB string    `json:"national_team_b" validate:"required,uuid" gorm:"type:varchar(255)"`
	GolA          int8      `json:"gol_a" validate:"required" gorm:"type:integer"`
	GolB          int8      `json:"gol_b" validate:"required" gorm:"type:integer"`
	KeyID         string    `json:"key_id" validate:"required,uuid" gorm:"type:varchar(255)"`
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
