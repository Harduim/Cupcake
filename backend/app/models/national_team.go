package models

type NationalTeam struct {
	ID    string `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name  string `json:"name" validate:"required" gorm:"type:varchar(255)"`
	KeyID string `json:"key_id" validate:"required,uuid" gorm:"type:varchar(255)"`
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
