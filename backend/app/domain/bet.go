package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bet struct {
	ID              string        `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	CreatedAt       time.Time     `json:"created_at"`
	GolA            *int          `json:"gol_a" validate:"required" gorm:"type:integer"`
	GolB            *int          `json:"gol_b" validate:"required" gorm:"type:integer"`
	UserID          string        `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	User            *User         `gorm:"foreignKey:UserID"`
	MatchID         string        `json:"match_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Match           *Match        `gorm:"foreignKey:MatchID"`
	NationalTeamAID string        `json:"national_team_a" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamBID string        `json:"national_team_b" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamA   *NationalTeam `gorm:"foreignKey:NationalTeamAID"`
	NationalTeamB   *NationalTeam `gorm:"foreignKey:NationalTeamBID"`
	WinnerID        string        `json:"winner_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Winner          *NationalTeam `gorm:"foreignKey:WinnerID"`
}

func NewBet(nationalTeamAID string,
	nationalTeamBID string,
	matchID string,
	userID string,
	golA *int,
	golB *int,
	winnerId string) (*Bet, error) {

	bet := &Bet{
		NationalTeamAID: nationalTeamAID,
		NationalTeamBID: nationalTeamBID,
		MatchID:         matchID,
		UserID:          userID,
		GolA:            golA,
		GolB:            golB,
		WinnerID:        winnerId,
	}
	bet.prepare()

	err := bet.Validate()

	if err != nil {
		return nil, err
	}

	return bet, nil
}

func (bet *Bet) prepare() {
	now := time.Now().UTC()
	bet.ID = uuid.NewV4().String()
	bet.CreatedAt = now
}

func (bet *Bet) Validate() error {
	err := validate.Struct(bet)

	if err != nil {
		return err
	}
	return nil
}
