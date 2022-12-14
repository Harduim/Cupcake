package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Bet struct {
	ID        string    `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	GolA      *int      `json:"golA" validate:"required" gorm:"type:integer"`
	GolB      *int      `json:"golB" validate:"required" gorm:"type:integer"`
	UserID    string    `json:"userId" validate:"required,uuid" gorm:"type:varchar(255)"`
	User      *User     `gorm:"foreignKey:UserID"`
	// TODO: add unique composto com matchId a userId
	MatchID  string        `json:"matchId" validate:"required,uuid" gorm:"type:varchar(255)"`
	Match    *Match        `gorm:"foreignKey:MatchID"`
	WinnerID string        `json:"winnerId" validate:"required,uuid" gorm:"type:varchar(255)"`
	Winner   *NationalTeam `gorm:"foreignKey:WinnerID"`
}

func (bet *Bet) Prepare() {
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
