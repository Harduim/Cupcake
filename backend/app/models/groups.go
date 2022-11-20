package models

type Group struct {
	UserID         string `json:"user_id" validate:"required,uuid" gorm:"type:varchar(255);primary_key"`
	NationalTeamID string `json:"national_team_id" gorm:"type:varchar(255);primary_key"`
}

type Groups struct {
	Groups []Group
}
