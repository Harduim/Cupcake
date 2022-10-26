package domain

type Bracket struct {
	ID         string `json:"id" validate:"required,uuid" gorm:"type:uuid;primary_key"`
	Name       string `json:"name" validate:"required" gorm:"type:varchar(255)"`
	Multiplier int8   `json:"multiplier" validate:"required"`
}

func NewBracket() *Bracket {
	return &Bracket{}
}

func (bracket *Bracket) Validate() error {
	err := validate.Struct(bracket)

	if err != nil {
		return err
	}
	return nil
}
