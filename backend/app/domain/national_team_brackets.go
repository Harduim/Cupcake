package domain

type NationalTeamBracket struct {
	NationalTeamID string       `json:"national_team_id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	NationalTeam   NationalTeam `gorm:"foreignKey:NationalTeamID"`
	BracketID      string       `json:"bracket_id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Bracket        Bracket      `gorm:"foreignKey:BracketID"`
}

func NewNationalTeamBracket() *NationalTeamBracket {
	return &NationalTeamBracket{}
}

func (nationalTeamBracket *NationalTeamBracket) Validate() error {
	err := validate.Struct(nationalTeamBracket)

	if err != nil {
		return err
	}
	return nil
}
