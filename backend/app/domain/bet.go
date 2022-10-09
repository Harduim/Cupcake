package domain

import "time"

type Bet struct {
	ID              string        `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	CreatedAt       time.Time     `json:"created_at"`
	GolA            int8          `json:"gol_a" validate:"required" gorm:"type:integer"`
	GolB            int8          `json:"gol_b" validate:"required" gorm:"type:integer"`
	UserID          string        `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	User            User          `gorm:"foreignKey:UserID"`
	MatchID         string        `json:"match_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Match           *Match        `gorm:"foreignKey:MatchID"`
	NationalTeamAID string        `json:"national_team_a" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamBID string        `json:"national_team_b" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamA   *NationalTeam `gorm:"foreignKey:NationalTeamAID"`
	NationalTeamB   *NationalTeam `gorm:"foreignKey:NationalTeamBID"`
}

func NewBet() *Bet {
	return &Bet{}
}

func (bet *Bet) Validate() error {
	err := validate.Struct(bet)

	if err != nil {
		return err
	}
	return nil
}
