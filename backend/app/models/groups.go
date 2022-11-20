package models

import (
	"errors"
)

type Groups struct {
	UserID        string          `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255);primary_key"`
	NationalTeams []*NationalTeam `gorm:"many2many:user_groups;"`
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
