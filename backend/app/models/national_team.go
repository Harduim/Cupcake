package models

type NationalTeam struct {
	ID       string    `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name     string    `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Brackets []Bracket `gorm:"many2many:national_team_brackets;"`
}

func NewNationalTeam() *NationalTeam {
	return &NationalTeam{}
}

func (nationalTeam *NationalTeam) Validate() error {
	err := validate.Struct(nationalTeam)

	if err != nil {
		return err
	}
	return nil
}
