package models

import (
	"errors"
	"time"
)

type Groups struct {
	CreatedAt     time.Time       `json:"created_at"`
	UserID        string          `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255);primary_key"`
	BracketID     string          `json:"bracket_id" validate:"required,uuid" gorm:"type:varchar(255)"`
	Bracket       Bracket         `gorm:"foreignKey:BracketID"`
	User          *User           `gorm:"foreignKey:UserID"`
	NationalTeams []*NationalTeam `gorm:"many2many:user_groups;"`
}

func NewGroups(
	NationalTeams []*NationalTeam,
	userID string,
	bracketID string,
) (*Groups, error) {

	groups := &Groups{
		UserID:        userID,
		NationalTeams: NationalTeams,
		BracketID:     bracketID,
	}
	groups.prepare()

	err := groups.Validate()

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (groups *Groups) prepare() {
	now := time.Now().UTC()
	groups.CreatedAt = now
}

func (groups *Groups) Validate() error {
	err := validate.Struct(groups)

	if err != nil {
		return err
	}

	if len(groups.NationalTeams) > 16 {
		return errors.New("only 16 teams can be chosen")
	}

	if err != nil {
		return err
	}
	return nil
}