package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Joker struct {
	ID              string        `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	CreatedAt       time.Time     `json:"created_at"`
	GolA            *int          `json:"gol_a" validate:"required" gorm:"type:integer"`
	GolB            *int          `json:"gol_b" validate:"required" gorm:"type:integer"`
	UserID          string        `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255);primary_key"`
	User            *User         `gorm:"foreignKey:UserID"`
	BracketID       string        `json:"bracket_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Bracket         *Bracket      `gorm:"foreignKey:BracketID"`
	NationalTeamAID string        `json:"national_team_a" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamBID string        `json:"national_team_b" validate:"required,uuid" gorm:"type:varchar(255)"`
	NationalTeamA   *NationalTeam `gorm:"foreignKey:NationalTeamAID"`
	NationalTeamB   *NationalTeam `gorm:"foreignKey:NationalTeamBID"`
	WinnerID        string        `json:"winner_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Winner          *NationalTeam `gorm:"foreignKey:WinnerID"`
}

func NewJoker(nationalTeamAID string,
	nationalTeamBID string,
	matchID string,
	userID string,
	golA *int,
	golB *int,
	winnerId string) (*Joker, error) {

	joker := &Joker{
		NationalTeamAID: nationalTeamAID,
		NationalTeamBID: nationalTeamBID,
		BracketID:       matchID,
		UserID:          userID,
		GolA:            golA,
		GolB:            golB,
		WinnerID:        winnerId,
	}
	joker.prepare()

	err := joker.Validate()

	if err != nil {
		return nil, err
	}

	return joker, nil
}

func (joker *Joker) prepare() {
	now := time.Now().UTC()
	joker.ID = uuid.NewV4().String()
	joker.CreatedAt = now
}

func (joker *Joker) Validate() error {
	err := validate.Struct(joker)

	if err != nil {
		return err
	}
	return nil
}
